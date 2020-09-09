package helpers

const (
	CreatePermission = "INSERT INTO permissions(name, description) VALUES (?,?)"
	CreateRoles = "INSERT INTO roles(name, description) VALUES(?,?)"
	CreateRolesPermissions = "INSERT INTO roles_permissions(roles_id, permission_id) VALUES(?,?)"
	CreateUser  = "INSERT INTO users(first_name , last_name, email, phone_number, dob, password) VALUES( ?, ?, ?, ?,?,? )"
	CreateUsersRoles = "INSERT INTO users_roles(users_id, roles_id) VALUES (?, ?)"
	CreateTags = "INSERT INTO tags(name, description) VALUES (?,?)"
	CreateUsersTags = "INSERT INTO users_tags(user_id, tag_id) VALUES(?,?)"
	)
