package usersstore

const (
	QueryCreateUser           = "INSERT INTO auth.users(users.email, users.password_hash) VALUES (?,?);"
	QueryCreateSessionUser    = "INSERT INTO auth.sessions(sessions.user_id, sessions.refresh_token, sessions.is_revoked, sessions.expires_at) VALUES (?,?,?,?);"
	QueryGetOneUserByEmail    = "SELECT u.user_id, r.role_name, u.email, u.password_hash FROM auth.users AS u INNER JOIN auth.user_roles ur ON u.user_id = ur.user_id INNER JOIN auth.roles r ON ur.role_id = r.role_id WHERE u.email =?;"
	QueryDeleteSession        = "DELETE FROM auth.sessions WHERE refresh_token = ?;"
	QueryRevokeUserSession    = "UPDATE auth.sessions SET is_revoked = true WHERE refresh_token = ?;"
	QueryAssignRoleToUser     = "INSERT INTO auth.user_roles (user_id, role_id) VALUES (?, ?);"
	QueryGetRoleName          = "SELECT r.role_name FROM auth.roles r WHERE r.role_id = ?;"
	QueryGetSessionDetail     = "SELECT s.user_id, s.refresh_token, s.is_revoked, s.expires_at FROM auth.sessions s WHERE s.refresh_token = ?;"
	QueryGetOneUserById       = "SELECT r.role_name, u.email, u.password_hash FROM auth.users AS u INNER JOIN auth.user_roles ur ON u.user_id = ur.user_id INNER JOIN auth.roles r ON ur.role_id = r.role_id WHERE u.user_id =?;"
)