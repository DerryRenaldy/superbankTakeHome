package usershandler

// import (
// 	usersservice "authenticationService/api/service/auth"
// 	usersrespdto "authenticationService/dto/response/auth"
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/DerryRenaldy/logger/logger"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// )

// func TestUserHandlerImpl_Login(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	userServiceMock := usersservice.NewMockIService(ctrl)
// 	log := logger.New("service_test", "dev", "testing")

// 	u := &UserHandlerImpl{
// 		userService: userServiceMock,
// 		l:           log,
// 	}

// 	recorder := httptest.NewRecorder()

// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		setMock func()
// 		wantErr bool
// 	}{
// 		{
// 			name: "Success",
// 			args: args{
// 				w: recorder,
// 				r: func() *http.Request {
// 					req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(
// 						fmt.Sprintf(
// 							`{
// 						"user_identity":"%s",
// 						"password":"%s"
// 					}`, "test@test.com", "test_password"),
// 					))

// 					return req.WithContext(context.TODO())
// 				}(),
// 			},
// 			setMock: func() {
// 				userServiceMock.EXPECT().Login(gomock.Any(), gomock.Any()).Return("token", &usersrespdto.UserResponse{
// 					FullName:     "Test Full Name",
// 					Username:     "test_username",
// 					Email:        "test@test.com",
// 					PasswordHash: "passwordHash",
// 				}, nil)
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "Failure - invalid body reader",
// 			args: args{
// 				w: recorder,
// 				r: func() *http.Request {
// 					req := httptest.NewRequest(http.MethodPost, "/login", nil)

// 					return req.WithContext(context.TODO())
// 				}(),
// 			},
// 			setMock: func() {
// 				// no need mock call
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Failure - error unmarshall",
// 			args: args{
// 				w: recorder,
// 				r: func() *http.Request {
// 					req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(
// 						fmt.Sprintf(
// 							`{
// 						"user_identity":%d,
// 						"password":%d
// 					}`, 1, 2),
// 					))

// 					return req.WithContext(context.TODO())
// 				}(),
// 			},
// 			setMock: func() {
// 				// no need mock call
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "Failure - error validate request body",
// 			args: args{
// 				w: recorder,
// 				r: func() *http.Request {
// 					req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(
// 						fmt.Sprintf(
// 							`{
// 						"user_identity":"%s",
// 						"password":"%s"
// 					}`, "", ""),
// 					))

// 					return req.WithContext(context.TODO())
// 				}(),
// 			},
// 			setMock: func() {
// 				// no need mock call
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setMock()
// 			if err := u.Login(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
// 				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})

// 		// Verify the cookie
// 		if !tt.wantErr {
// 			recorder := tt.args.w.(*httptest.ResponseRecorder)
// 			cookies := recorder.Result().Cookies()
// 			assert.NotEmpty(t, cookies, "Expected cookie to be set, but it was not")

// 			found := false
// 			for _, cookie := range cookies {
// 				if cookie.Name == "AuthToken" && cookie.Value == "token" {
// 					found = true
// 					break
// 				}
// 			}
// 			assert.True(t, found, "Expected cookie 'AuthToken' with value '%s' but it was not found", "token")
// 		}
// 	}
// }
