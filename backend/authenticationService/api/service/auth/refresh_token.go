package usersservice

import (
	usersrespdto "authenticationService/dto/response/auth"
	cError "authenticationService/pkgs/errors"
	"context"
	"errors"
	"fmt"
	"time"
)

func (u *UserServiceImpl) RefreshToken(ctx context.Context, refreshToken string) (*usersrespdto.RenewAccessTokenResponse, error) {
	functionName := "UserServiceImpl.Login"

	session, err := u.userRepo.GetSessionDetail(ctx, refreshToken)
	if err != nil {
		u.l.Errorf("[%s] = Fail to get session detail! : %s", functionName, err.Error())
		return nil, err
	}

	fmt.Println("session", session)

	user, err := u.userRepo.GetUserById(ctx, session.UserID)
	if err != nil {
		u.l.Errorf("[%s] = Fail to get user detail! : %s", functionName, err.Error())
		return nil, err
	}

	if session.IsRevoked {
		u.l.Errorf("[%s] = Session is revoked! : %s", functionName, errors.New("session is revoked"))
		return nil, cError.GetError(cError.UnauthorizedError, errors.New("session is revoked"))
	}

	if session.RefreshTokenExpiresAt.Before(time.Now()) {
		u.l.Errorf("[%s] = Refresh token is expired! : %s", functionName, errors.New("refresh token is expired"))
		return nil, cError.GetError(cError.UnauthorizedError, errors.New("refresh token is expired"))
	}

	accessToken, accessClaims, err := u.tokenMaker.GenerateTokenJWT(user, time.Minute*15)
	if err != nil {
		u.l.Errorf("[%s] = Fail to generate JWT token! : %s", functionName, errors.New("jwt token not created"))
		return nil, err
	}

	err = u.StoreTokenInCache(ctx, user.Email, accessToken)
	if err != nil {
		u.l.Errorf("[%s] = Fail to store token in cache! : %s", functionName, err.Error())
		return nil, err
	}

	result := &usersrespdto.RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.RegisteredClaims.ExpiresAt.Time,
	}

	return result, nil
}