package usersservice

import (
	usersreqdto "authenticationService/dto/request/auth"
	usersrespdto "authenticationService/dto/response/auth"
	utils "authenticationService/pkgs"
	cError "authenticationService/pkgs/errors"
	"context"
	"errors"
	"time"
)

func (u *UserServiceImpl) Login(ctx context.Context, payload *usersreqdto.LoginRequest) (*usersrespdto.RegisterLoginResponse, error) {
	functionName := "UserServiceImpl.Login"

	userData := new(usersrespdto.UserResponse)

	if utils.IsInputEmail(payload.Email) {
		user, err := u.userRepo.GetUserByEmail(ctx, payload.Email)
		if err != nil {
			return nil, err
		}

		userData = user
	}

	if !u.util.MatchPassword(payload.Password, userData.PasswordHash) {
		u.l.Errorf("[%s] = Password not match! : %s", functionName, errors.New("incorrect password"))
		return nil, cError.GetError(cError.InvalidRequestError, errors.New("password not match"))
	}

	refreshToken, refreshClaims, err := u.tokenMaker.GenerateTokenJWT(userData, time.Minute * time.Duration(u.cfg.TokenCache.RefreshTokenTimeout))
	if err != nil {
		u.l.Errorf("[%s] = Fail to generate JWT token! : %s", functionName, errors.New("jwt token not created"))
		return nil, err
	}

	accessToken, accessClaims, err := u.tokenMaker.GenerateTokenJWT(userData, time.Minute * time.Duration(u.cfg.TokenCache.AccessTokenTimeout))
	if err != nil {
		u.l.Errorf("[%s] = Fail to generate JWT token! : %s", functionName, errors.New("jwt token not created"))
		return nil, err
	}

	err = u.StoreTokenInCache(ctx, userData.Email, accessToken)
	if err != nil {
		u.l.Errorf("[%s] = Fail to store token in cache! : %s", functionName, err.Error())
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
		User:                  userData,
	}

	return result, nil
}
