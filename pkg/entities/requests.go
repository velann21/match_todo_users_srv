package entities

import (
	"errors"
	pf "github.com/velann21/todo-commonlib/proto_files"
)

func ValidateUserRegistration(req *pf.UserRegistrationRequest) error {
	if req.FirstName == "" {
		return errors.New("Invalid request")
	}
	if req.EmailId == "" {
		return errors.New("Invalid request")
	}
	if req.LastName == "" {
		return errors.New("Invalid request")
	}
	if req.Password == "" {
		return errors.New("Invalid request")
	}
	if req.ConfirmPassword == "" {
		return errors.New("Invalid request")
	}
	if req.Tags == "" {
		return errors.New("Invalid request")
	}
	if req.Location.Lng == 0 {
		return errors.New("Invalid request")
	}
	if req.Location.Lat == 0 {
		return errors.New("Invalid request")
	}
	return nil
}

func ValidateCreatePermissionRequest(req *pf.CreatePermissionRequest) error {
	if req.Name == ""{
		return errors.New("Invalid request")
	}
	if req.Description == ""{
		return errors.New("Invalid request")
	}
	return nil
}

func ValidateCreateTagsRequest(req *pf.CreateTagRequest)error{
	if req.Name == ""{
		return errors.New("Invalid request")
	}
	if req.Description == ""{
		return errors.New("Invalid request")
	}
	return nil
}

func ValidateCreateRolesRequest(req *pf.CreateRoleRequest)error{
	if req.Name == ""{
		return errors.New("Invalid request")
	}
	if req.Description == ""{
		return errors.New("Invalid request")
	}
	return nil
}

func ValidateAttachRolesPermission(req *pf.AttachRolesPermissionsRequest)error{
	if len(req.Roles) < 0{
		return errors.New("Invalid request")
	}
	if len(req.Permissions) < 0{
		return errors.New("Invalid request")
	}
	return nil
}

func ValidateAttachUsersRoles(req *pf.AttachUsersRolesRequest)error{
	if len(req.Roles) < 0{
		return errors.New("Invalid request")
	}
	if len(req.Users) < 0{
		return errors.New("Invalid request")
	}
	return nil
}

func ValidateSqlMigrationRequest(req *pf.SqlMigrationRequest)error{
	if req.Upcount < 0 && req.Downcount < 0{
		return errors.New("Invalid request")
	}

	if req.Upcount > 0 && req.Downcount > 0{
		return errors.New("Invalid request")
	}

	return nil
}