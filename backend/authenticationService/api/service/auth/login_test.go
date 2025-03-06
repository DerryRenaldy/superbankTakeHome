package usersservice

// import (
// 	usersrepo "authenticationService/api/repositories/auth"
// 	usersreqdto "authenticationService/dto/request/auth"
// 	usersrespdto "authenticationService/dto/response/auth"
// 	utils "authenticationService/pkgs"
// 	"context"
// 	"errors"
// 	"reflect"
// 	"testing"

// 	"github.com/DerryRenaldy/logger/logger"
// 	"github.com/golang/mock/gomock"
// )

// func TestUserServiceImpl_Login(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	userRepo := usersrepo.NewMockIRepository(ctrl)
// 	utility := utils.NewMockIUtils(ctrl)
// 	log := logger.New("service_test", "dev", "testing")

// 	u := &UserServiceImpl{
// 		userRepo: userRepo,
// 		util:     utility,
// 		l:        log,
// 	}

// 	hashedPassword, _ := utils.GeneratePasswordHash("testingPassword")

// 	ctx := context.TODO()

// 	type args struct {
// 		ctx     context.Context
// 		payload *usersreqdto.LoginRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		setMock func()
// 		want    string
// 		want1   *usersrespdto.UserResponse
// 		wantErr bool
// 	}{
// 		{
// 			name: "Success - login with username",
// 			args: args{
// 				ctx: ctx,
// 				payload: &usersreqdto.LoginRequest{
// 					UserIdentity: "testing",
// 					Password:     "testing123",
// 				},
// 			},
// 			setMock: func() {
// 				userRepo.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(&usersrespdto.UserResponse{
// 					Username:     "testing",
// 					Email:        "testing@gmail.com",
// 					PasswordHash: hashedPassword,
// 				}, nil)
// 				utility.EXPECT().MatchPassword(gomock.Any(), gomock.Any()).Return(true)
// 				utility.EXPECT().GenerateJWT(gomock.Any()).Return("tokenTest", nil)
// 			},
// 			want: "tokenTest",
// 			want1: &usersrespdto.UserResponse{
// 				Username:     "testing",
// 				Email:        "testing@gmail.com",
// 				PasswordHash: hashedPassword,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Success - login with email",
// 			args: args{
// 				ctx: ctx,
// 				payload: &usersreqdto.LoginRequest{
// 					UserIdentity: "testing@gmail.com",
// 					Password:     "testing123",
// 				},
// 			},
// 			setMock: func() {
// 				userRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(&usersrespdto.UserResponse{
// 					Username:     "testing",
// 					Email:        "testing@gmail.com",
// 					PasswordHash: hashedPassword,
// 				}, nil)
// 				utility.EXPECT().MatchPassword(gomock.Any(), gomock.Any()).Return(true)
// 				utility.EXPECT().GenerateJWT(gomock.Any()).Return("tokenTest", nil)
// 			},
// 			want: "tokenTest",
// 			want1: &usersrespdto.UserResponse{
// 				Username:     "testing",
// 				Email:        "testing@gmail.com",
// 				PasswordHash: hashedPassword,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Failed - password not match",
// 			args: args{
// 				ctx: ctx,
// 				payload: &usersreqdto.LoginRequest{
// 					UserIdentity: "testing",
// 					Password:     "testing123",
// 				},
// 			},
// 			setMock: func() {
// 				userRepo.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(&usersrespdto.UserResponse{
// 					Username:     "testing",
// 					Email:        "testing@gmail.com",
// 					PasswordHash: hashedPassword,
// 				}, nil)
// 				utility.EXPECT().MatchPassword(gomock.Any(), gomock.Any()).Return(false)
// 			},
// 			want:    "",
// 			want1:   nil,
// 			wantErr: true,
// 		},
// 		{
// 			name: "Failed - JWT token not created",
// 			args: args{
// 				ctx: ctx,
// 				payload: &usersreqdto.LoginRequest{
// 					UserIdentity: "testing",
// 					Password:     "testing123",
// 				},
// 			},
// 			setMock: func() {
// 				userRepo.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).Return(&usersrespdto.UserResponse{
// 					Username:     "testing",
// 					Email:        "testing@gmail.com",
// 					PasswordHash: hashedPassword,
// 				}, nil)
// 				utility.EXPECT().MatchPassword(gomock.Any(), gomock.Any()).Return(true)
// 				utility.EXPECT().GenerateJWT(gomock.Any()).Return("", errors.New("error token not created"))
// 			},
// 			want:    "",
// 			want1:   nil,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setMock()
// 			got, got1, err := u.Login(tt.args.ctx, tt.args.payload)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("Login() got = %v, want %v", got, tt.want)
// 			}
// 			if !reflect.DeepEqual(got1, tt.want1) {
// 				t.Errorf("Login() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }
