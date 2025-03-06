package usersservice

import (
	"authenticationService/constants"
	usersreqdto "authenticationService/dto/request/auth"
	usersrespdto "authenticationService/dto/response/auth"
	cError "authenticationService/pkgs/errors"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

func (u *UserServiceImpl) Register(ctx context.Context, payload *usersreqdto.CreateUserRequest) (*usersrespdto.RegisterLoginResponse, error) {
	functionName := "UserServiceImpl.Register"

	hashedPassword, err := u.util.GeneratePasswordHash(payload.Password)
	if err != nil {
		u.l.Errorf("[%s] = While Generating Hashed Password : %s", functionName, err.Error())
		return nil, cError.GetError(cError.InternalServerError, err)
	}

	u.l.Infof("Hashed Password : [%s]-[%s]", payload.Password, hashedPassword)

	payload.Password = hashedPassword

	user, err := u.userRepo.CreateUser(ctx, payload)
	if err != nil {
		// Check for MySQL specific duplicate key error (error number 1062)
		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) {
			u.l.Debugf("[%s] = While Create User : %s", functionName, err.Error())
			if mysqlErr.Number == 1062 {
				return nil, cError.GetError(cError.ConflictError, errors.New("user already exists"))
			}
		}
		return nil, err
	}

	fmt.Println("user", user)

	roleName, err := u.userRepo.AssignRoleToUser(ctx, user.UserID, constants.RoleUser)
	if err != nil {
		u.l.Errorf("[%s] = Fail to assign role to user! : %s", functionName, err.Error())
		return nil, err
	}

	user.Role = roleName

	refreshToken, refreshClaims, err := u.tokenMaker.GenerateTokenJWT(user, time.Hour*24)
	if err != nil {
		u.l.Errorf("[%s] = Fail to generate JWT token! : %s", functionName, errors.New("jwt token not created"))
		return nil, err
	}

	accessToken, accessClaims, err := u.tokenMaker.GenerateTokenJWT(user, time.Minute*15)
	if err != nil {
		u.l.Errorf("[%s] = Fail to generate JWT token! : %s", functionName, errors.New("jwt token not created"))
		return nil, err
	}

	err = u.userRepo.CreateUserSession(ctx, &usersrespdto.Session{
		UserID:                refreshClaims.UserID,
		RefreshToken:          refreshToken,
		IsRevoked:             false,
		RefreshTokenExpiresAt: refreshClaims.RegisteredClaims.ExpiresAt.Time,
	})
	if err != nil {
		u.l.Errorf("[%s] = Fail to create Session! : %s", functionName, errors.New("session not created"))
		return nil, err
	}

	result := &usersrespdto.RegisterLoginResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessClaims.RegisteredClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshClaims.RegisteredClaims.ExpiresAt.Time,
		User:                  user,
	}

	return result, nil
}
