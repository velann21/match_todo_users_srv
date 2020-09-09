package routes

import (
	"context"
	"database/sql"
	//"database/sql"
	"github.com/velann21/match_todo_user_srv/pkg/repository"
	"github.com/velann21/match_todo_user_srv/pkg/service"

	//"github.com/gorilla/mux"
	"github.com/velann21/match_todo_user_srv/pkg/controllers"
	pf "github.com/velann21/todo-commonlib/proto_files"
)
type ServerRoutes struct {
     GrpcController *controllers.GrpcController
}

func (server *ServerRoutes) AttachUsersTags(context.Context, *pf.AttachUsersTagsRequest) (*pf.AttachTagsResponse, error) {
	return nil, nil
}

//func (i *ServerRoutes) ApiRoutes(route *mux.Router){
//	gwc := controllers.UserController{
//		EtcdClient:nil,
//	}
//	route.PathPrefix("/users/healthy").HandlerFunc(gwc.HealthyMe).Methods("GET")
//	route.PathPrefix("/users/registration").HandlerFunc(gwc.UsersRegistrationController).Methods("POST")
//	route.PathPrefix("/users/listuser").HandlerFunc(gwc.ListUserController).Methods("GET")
//	route.PathPrefix("/users/listUsers").HandlerFunc(gwc.ListUsersController).Methods("GET")
//}

func (server *ServerRoutes) UserRegistration(ctx context.Context, req *pf.UserRegistrationRequest) (*pf.UserRegistrationResponse, error){
	return server.GrpcController.UserRegistrationController(ctx, req)
}

func (server *ServerRoutes) UserLogin(ctx context.Context, req *pf.UserLoginRequest) (*pf.UserLoginResponse, error){

	return server.GrpcController.UserLogin(ctx, req)
}

func (server *ServerRoutes) CreateRole(ctx context.Context, req *pf.CreateRoleRequest) (*pf.CreateRoleResponse, error) {
	return server.GrpcController.CreateRole(ctx, req)
}

func (server *ServerRoutes)  CreatePermission(ctx context.Context, req *pf.CreatePermissionRequest) (*pf.CreatePermissionResponse, error) {
	return server.GrpcController.CreatePermission(ctx, req)
}

func (server *ServerRoutes)  CreateTags(ctx context.Context, req *pf.CreateTagRequest) (*pf.CreateTagResponse, error) {
	return server.GrpcController.CreateTags(ctx, req)
}

func (server *ServerRoutes)  AttachRolesPermissions(ctx context.Context, req *pf.AttachRolesPermissionsRequest) (*pf.AttachRolesPermissionsResponse, error) {
	return server.GrpcController.AttachRolesPermissions(ctx, req)
}

func (server *ServerRoutes)  AttachUsersRoles(ctx context.Context, req *pf.AttachUsersRolesRequest) (*pf.AttachUsersRolesResponse, error) {
	return server.GrpcController.AttachUsersRoles(ctx, req)
}

func (server *ServerRoutes) AttachTagsUsers(ctx context.Context, req *pf.AttachUsersTagsRequest)(*pf.AttachTagsResponse, error){
	return server.GrpcController.AttachTagsUsers(ctx, req)
}

func (server *ServerRoutes) SqlMigration(ctx context.Context, req *pf.SqlMigrationRequest)(*pf.SqlMigrationResponse, error){
	return server.GrpcController.SqlMigration(ctx, req)
}

func Intialize(sql *sql.DB)*controllers.GrpcController{
	sqlImpl := repository.New(sql)
	repo := repository.UserRepository{SqlDriver:sqlImpl}
	userSrv := service.UserService{UserRepository:&repo}
	grpcController := controllers.GrpcController{Service: &userSrv}
	return &grpcController
}

