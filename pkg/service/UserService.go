package service

import (
	"context"
	"github.com/velann21/match_todo_user_srv/pkg/repository"
	pf "github.com/velann21/todo-commonlib/proto_files"
	migration "github.com/velann21/match_todo_user_srv/pkg/migration_scripts"
)

type USERID string

const EMTPYUSERID  = ""
type UserServiceInterface interface {
	UserRegistrationService(ctx context.Context, req *pf.UserRegistrationRequest)(string, error)
	UserLogin(ctx context.Context, req *pf.UserLoginRequest) (*pf.UserLoginResponse, error)
	CreateRole(ctx context.Context, req *pf.CreateRoleRequest) (string, error)
	CreatePermission(ctx context.Context, req *pf.CreatePermissionRequest) (string, error)
	CreateTags(ctx context.Context, req *pf.CreateTagRequest) (string, error)
	SqlMigration(ctx context.Context, req *pf.SqlMigrationRequest)error
}
type UserService struct {
	UserRepository repository.Repository
}

func (srv *UserService) UserRegistrationService(ctx context.Context, req *pf.UserRegistrationRequest)(string, error){


	return EMTPYUSERID, nil
}

func (srv *UserService) UserLogin(ctx context.Context, req *pf.UserLoginRequest) (*pf.UserLoginResponse, error){

	return nil, nil
}

func (srv *UserService) CreateRole(ctx context.Context, req *pf.CreateRoleRequest) (string, error) {
	ID, err := srv.UserRepository.CreateRole(ctx, req)
	if err != nil{
		return "", err
	}
	return ID, nil
}

func (srv *UserService) CreatePermission(ctx context.Context, req *pf.CreatePermissionRequest) (string, error) {
    ID, err := srv.UserRepository.CreatePermission(ctx, req)
    if err != nil{
		return "", err
	}
    return ID, nil
}

func (srv *UserService) CreateTags(ctx context.Context, req *pf.CreateTagRequest) (string, error) {
	return "", nil
}

func (srv *UserService) SqlMigration(ctx context.Context, req *pf.SqlMigrationRequest)error{
	if req.Upcount > 0{
		err := migration.MigrateDb(uint(req.Upcount))
		if err != nil{
			return err
		}
	}
	return nil
}


