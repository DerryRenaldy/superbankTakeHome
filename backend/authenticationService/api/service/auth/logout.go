package usersservice

import "context"

func (u *UserServiceImpl) Logout(ctx context.Context, refreshToken string) error {
	functionName := "UserServiceImpl.Logout"
	err := u.userRepo.RevokeUserSession(ctx, refreshToken)
	if err != nil {
		u.l.Errorf("[%s] = Fail to revoke user session! : %s", functionName, err)
		return err
	}

	return nil
}
