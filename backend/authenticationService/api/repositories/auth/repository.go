package usersrepo

import (
	usersreqdto "authenticationService/dto/request/auth"
	usersrespdto "authenticationService/dto/response/auth"
	usersstore "authenticationService/stores/postgres/auth"
	"context"
)

//go:generate mockgen -source repository.go -destination repository_mock.go -package usersrepo
type IRepository interface {
	CreateUser(ctx context.Context, payload *usersreqdto.CreateUserRequest) (*usersrespdto.UserResponse, error)
	CreateUserSession(ctx context.Context, payload *usersrespdto.Session) error
	DeleteUserSession(ctx context.Context, refreshToken string) error
	RevokeUserSession(ctx context.Context, refreshToken string) error
	GetUserByEmail(ctx context.Context, userEmail string) (*usersrespdto.UserResponse, error)
	AssignRoleToUser(ctx context.Context, userID int, roleID int) (string, error)
	GetSessionDetail(ctx context.Context, refreshToken string)(*usersrespdto.Session, error)
	GetUserById(ctx context.Context, userId int)(*usersrespdto.UserResponse, error)
}

var _ IRepository = (*usersstore.UserRepoImpl)(nil)