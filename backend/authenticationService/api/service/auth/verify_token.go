package usersservice

import (
	usersrespdto "authenticationService/dto/response/auth"
	cError "authenticationService/pkgs/errors"
	"context"
	"errors"
	"fmt"
)

func (u *UserServiceImpl) VerifyToken(ctx context.Context, accessToken string) (*usersrespdto.VerifyTokenResponse, error) {
	functionName := "UserServiceImpl.VerifyToken"

	claims, err := u.tokenMaker.ValidateJWT(accessToken)
	if err != nil {
		u.l.Errorf("[%s] = Fail to validate JWT token! : %s", functionName, err.Error())
		return nil, cError.GetError(cError.UnauthorizedError, err)
	}

	accessTokenFromCache, err := u.tokenCache.GetToken(ctx, fmt.Sprintf("access_token:%s", claims.Email))
	if err != nil {
		u.l.Errorf("[%s] = Fail to get token from cache! : %s", functionName, err.Error())
		return nil, cError.GetError(cError.UnauthorizedError, err)
	}

	u.l.Infof("Access token from cache: %s", accessTokenFromCache)
	u.l.Infof("Access token from request: %s", accessToken)

	if accessTokenFromCache != accessToken {
		u.l.Errorf("[%s] = Access token from cache is not equal to access token! : %s", functionName, errors.New("access token from cache is not equal to access token"))
		return nil, cError.GetError(cError.UnauthorizedError, errors.New("access token from cache is not equal to access token"))
	}

	response := &usersrespdto.VerifyTokenResponse{
		Status: "success",
		Message: "Token is valid",
		Data: usersrespdto.VerifyTokenUserDetail{
			Email: claims.Email,
			Role: claims.Role,
			ExpiredAt: claims.ExpiresAt.Time,
		},
	}

	return response, nil
}
