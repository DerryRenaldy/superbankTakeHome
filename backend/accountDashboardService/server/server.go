package server

import (
	"accountDashboardService/api/handler"
	service "accountDashboardService/api/services"
	config "accountDashboardService/configs"
	"accountDashboardService/pkgs/database/mysql"
	"accountDashboardService/server/middleware"
	authagent "accountDashboardService/stores/agents/auth_agent"
	store "accountDashboardService/stores/mysql"
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
	serviceUser service.IService
	handlerUser handler.IHandler
	authMiddleware middleware.AuthMiddleware
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
	// Initiate SQL Connection
	dbConn := mysql.NewConnection(s.log)
	if dbConn == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	if db = dbConn.Connect(); db == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	authClient := authagent.NewClientAuth(s.cfg, s.log)

	repo := store.NewRepoImpl(db, s.log)
	s.serviceUser = service.NewServiceImpl(repo, s.log)
	s.handlerUser = handler.NewHandlerImpl(s.log, s.serviceUser)
	s.authMiddleware = middleware.AuthMiddleware{
		AuthClient: authClient,
		Logger:     s.log,
	}
}

func (s *Server) StartServer() {
	addr := ":8090"

	// Initialize CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8090"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Cookie", "retry"},
		Debug:           true,
	})

	r := mux.NewRouter()
	r.Use(s.authMiddleware.VerifyToken)

	public := r.PathPrefix("/v1/dashboard").Subrouter()
	public.Handle("/accounts", middleware.ErrHandler(s.handlerUser.GetListAccount)).Methods(http.MethodGet)

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
