package usersstore

const (
	QueryCreateUser        = "INSERT INTO auth.users(email, password_hash) VALUES ($1, $2) RETURNING user_id;"

	QueryCreateSessionUser = "INSERT INTO auth.sessions(user_id, refresh_token, is_revoked, expires_at) VALUES ($1, $2, $3, $4);"
	
	QueryGetOneUserByEmail = `
			SELECT u.user_id, r.role_name, u.email, u.password_hash 
			FROM auth.users AS u 
			INNER JOIN auth.user_roles ur ON u.user_id = ur.user_id 
			INNER JOIN auth.roles r ON ur.role_id = r.role_id 
			WHERE u.email = $1;`
	
	QueryDeleteSession = "DELETE FROM auth.sessions WHERE refresh_token = $1;"
	
	QueryRevokeUserSession = "UPDATE auth.sessions SET is_revoked = true WHERE refresh_token = $1;"
	
	QueryAssignRoleToUser = "INSERT INTO auth.user_roles (user_id, role_id) VALUES ($1, $2);"
	
	QueryGetRoleName = "SELECT role_name FROM auth.roles WHERE role_id = $1;"
	
	QueryGetSessionDetail = "SELECT user_id, refresh_token, is_revoked, expires_at FROM auth.sessions WHERE refresh_token = $1;"
	
	QueryGetOneUserById = `
			SELECT u.user_id, r.role_name, u.email, u.password_hash 
			FROM auth.users AS u 
			INNER JOIN auth.user_roles ur ON u.user_id = ur.user_id 
			INNER JOIN auth.roles r ON ur.role_id = r.role_id 
			WHERE u.user_id = $1;`
)
