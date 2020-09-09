package controllers

import (
	"context"
	"github.com/velann21/match_todo_user_srv/pkg/entities"
	"github.com/velann21/match_todo_user_srv/pkg/service"
	pf "github.com/velann21/todo-commonlib/proto_files"
)

type GrpcController struct {
	Service service.UserServiceInterface
}

func (grpcCtrl *GrpcController) UserRegistrationController(ctx context.Context, req *pf.UserRegistrationRequest) (*pf.UserRegistrationResponse, error) {
	err := entities.ValidateUserRegistration(req)
	if err != nil{
		return nil, err
	}
	ID, err := grpcCtrl.Service.UserRegistrationService(ctx, req)
	if err != nil{
		return nil, err
	}

	finalResp := &pf.UserRegistrationResponse{
		Success:true,
		UserId: string(ID),
	}
	return finalResp, nil
}

func (grpcCtrl *GrpcController) UserLogin(ctx context.Context, req *pf.UserLoginRequest) (*pf.UserLoginResponse, error){

	return nil, nil
}

func (grpcCtrl *GrpcController) CreateRole(ctx context.Context, req *pf.CreateRoleRequest) (*pf.CreateRoleResponse, error) {
	err := entities.ValidateCreateRolesRequest(req)
	if err != nil{
		return nil, err
	}
	ID, err := grpcCtrl.Service.CreateRole(ctx, req)
	if err != nil{
		return nil, err
	}
	resp := &pf.CreateRoleResponse{
		Success:true,
		RoleId:ID,
	}
	return resp, nil
}

func (grpcCtrl *GrpcController) CreatePermission(ctx context.Context, req *pf.CreatePermissionRequest) (*pf.CreatePermissionResponse, error) {
	err := entities.ValidateCreatePermissionRequest(req)
	if err != nil{
		return nil, err
	}
	ID, err :=grpcCtrl.Service.CreatePermission(ctx, req)
	if err != nil{
		return nil, err
	}
	resp := &pf.CreatePermissionResponse{
		Success:true,
		PermissionId:ID,
	}
	return resp, nil
}

func (grpcCtrl *GrpcController) CreateTags(ctx context.Context, req *pf.CreateTagRequest) (*pf.CreateTagResponse, error) {
	err := entities.ValidateCreateTagsRequest(req)
	if err != nil{
		return nil, err
	}
	ID, err :=grpcCtrl.Service.CreateTags(ctx, req)
	if err != nil{
		return nil, err
	}
	resp := &pf.CreateTagResponse{
		Success:true,
		TagId:ID,
	}
	return resp, nil
}

func (grpcCtrl *GrpcController) AttachRolesPermissions(ctx context.Context, req *pf.AttachRolesPermissionsRequest) (*pf.AttachRolesPermissionsResponse, error) {
	//return grpcCtrl.GrpcController.CreatePermission(ctx, req)
	return nil, nil
}

func (grpcCtrl *GrpcController) AttachUsersRoles(ctx context.Context, req *pf.AttachUsersRolesRequest) (*pf.AttachUsersRolesResponse, error) {
	//return server.GrpcController.CreateTags(ctx, req)
	return nil, nil
}

func (grpcCtrl *GrpcController) AttachTagsUsers(ctx context.Context, req *pf.AttachUsersTagsRequest)(*pf.AttachTagsResponse, error){
	return nil, nil
}

func (grpcCtrl *GrpcController) SqlMigration(ctx context.Context, req *pf.SqlMigrationRequest)(*pf.SqlMigrationResponse, error){
	err := entities.ValidateSqlMigrationRequest(req)
	if err != nil{
		return nil, err
	}
	err = grpcCtrl.Service.SqlMigration(ctx, req)
	if err != nil{
		return nil, err
	}
	resp := &pf.SqlMigrationResponse{
		Success:true,
	}
	return resp, nil
}