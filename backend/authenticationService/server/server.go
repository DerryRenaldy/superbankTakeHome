package server

import (
	usershandler "authenticationService/api/handler"
	usersservice "authenticationService/api/service/auth"
	config "authenticationService/configs"
	"authenticationService/constants"
	utils "authenticationService/pkgs"
	"authenticationService/pkgs/database/mysql"
	"authenticationService/pkgs/database/redis"
	"authenticationService/pkgs/token"
	"authenticationService/server/middleware"
	authcachestore "authenticationService/stores/auth_cache"
	usersrepo "authenticationService/stores/mysql/auth"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/DerryRenaldy/logger/logger"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	cfg         *config.Config
	log         logger.ILogger
	tokenMaker  *token.JWTImpl
	serviceUser usersservice.IService
	handlerUser usershandler.IHandler
}

var signalChan chan os.Signal = make(chan os.Signal, 1)
var db *sql.DB
var SVR *Server

func NewServer(cfg *config.Config, log logger.ILogger) *Server {
	SVR = &Server{
		cfg: cfg,
		log: log,
	}

	SVR.RegisterServer()
	return SVR
}

func (s *Server) RegisterServer() {
	redisClient, err := redis.ConnectRedis(
		&redis.Config{
			Addr:     s.cfg.Redis.Addr,
			Timeout:  s.cfg.Redis.Timeout,
			PoolSize: s.cfg.Redis.PoolSize,
		},
	)
	if err != nil {
		s.log.Fatalf("Error while connecting to redis : ", err)
	}

	authRedisClient := authcachestore.New(redisClient, constants.RedisNamespace, s.cfg.Redis.Timeout)
	
	// Initiate SQL Connection
	dbConn := mysql.NewConnection(s.log)
	if dbConn == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	if db = dbConn.Connect(); db == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	utility := utils.NewUtilsImpl(s.log)
	userRepo := usersrepo.NewUserRepoImpl(db, s.log)
	s.tokenMaker = token.NewJWTImpl(s.cfg.JWTSecret)
	s.serviceUser = usersservice.NewUserServiceImpl(userRepo, utility, s.tokenMaker, s.log, authRedisClient, s.cfg)
	s.handlerUser = usershandler.NewUserHandlerImpl(s.serviceUser, s.log)
}

func (s *Server) StartServer() {
	addr := ":8091"

	r := mux.NewRouter()
	
	// Public endpoints
	public := r.PathPrefix("/v1/auth").Subrouter()
	public.Handle("/register", middleware.ErrHandler(s.handlerUser.Register)).Methods(http.MethodPost)
	public.Handle("/login", middleware.ErrHandler(s.handlerUser.Login)).Methods(http.MethodPost)
	public.Handle("/logout", middleware.ErrHandler(s.handlerUser.Logout)).Methods(http.MethodDelete)
	public.Handle("/refresh-token", middleware.ErrHandler(s.handlerUser.RefreshToken)).Methods(http.MethodGet)
	public.Handle("/verify-token", middleware.ErrHandler(s.handlerUser.VerifyToken)).Methods(http.MethodGet)
	// Initialize CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8091", "http://localhost:8090"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Cookie"},
		Debug:           true,
	})

	// Wrap the router with the CORS handler
	handler := c.Handler(r)

	go func() {
		err := http.ListenAndServe(addr, handler)
		if err != nil {
			fmt.Printf("error listening to address %v, err=%v", addr, err)
		}
	}()

	fmt.Printf("HTTP server started %v \n", addr)
	fmt.Println("Holding off the server")

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChan

	fmt.Printf("Interrupt Signal received : %s \n", sig)
	fmt.Println("Shutting down server...")

	// Clean up resources
	s.Close()
	fmt.Println("Server Terminated")
}

// Clean Up
func (s *Server) Close() {
	if db != nil {
		err := db.Close()
		if err != nil {
			s.log.Errorf("Error closing database connection: %v", err)
		} else {
			s.log.Info("Database connection closed")
		}
	}
}
