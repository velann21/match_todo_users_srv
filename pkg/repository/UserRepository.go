package repository

import (
	"context"
	"database/sql"
	"github.com/velann21/match_todo_user_srv/pkg/helpers"
	pf "github.com/velann21/todo-commonlib/proto_files"
	"strconv"
)

type Repository interface {
	CreateUser(ctx context.Context, request pf.UserRegistrationRequest)(string,error)
	GetUserByEmail(ctx context.Context,emailID string)
	CreateRole(ctx context.Context, req *pf.CreateRoleRequest)(string,error)
	CreatePermission(ctx context.Context, req *pf.CreatePermissionRequest)(string,error)
	CreateTags(ctx context.Context, req *pf.CreateTagRequest)(string,error)
	AttachRolesPermissions(ctx context.Context, req *pf.AttachRolesPermissionsRequest)(string,error)
	AttachUsersRoles(ctx context.Context, req *pf.AttachUsersRolesRequest)(string,error)
}

type UserRepository struct {
	SqlDriver SqlInterface
}

type USERID string

const USERIDCONT  = ""
func (rep *UserRepository) CreateUser(ctx context.Context, request pf.UserRegistrationRequest)(string,error){
	  tx, err := rep.SqlDriver.Begin()
	  if err != nil{
         return USERIDCONT, err
	  }
	  stmt, err := rep.SqlDriver.Prepare(tx, helpers.CreateUser)
	  if err != nil{
		  _ = rep.SqlDriver.RollBack(tx)
		  return USERIDCONT, err
	  }
	  res, err := rep.SqlDriver.Exec(stmt, request.FirstName, request.LastName, request.EmailId, request.EmailId, request.EmailId, request.Password)
	  if err != nil{
		  _ = rep.SqlDriver.RollBack(tx)
		  return USERIDCONT, err
	  }
	  id, err := rep.SqlDriver.LastInsertedID(res)
	  if err != nil{
		  _ = rep.SqlDriver.RollBack(tx)
		  return USERIDCONT, nil
	  }
	  _ = rep.SqlDriver.Commit(tx)
	  return strconv.Itoa(int(id)) , nil
}

func (rep *UserRepository) CreateTags(ctx context.Context, req *pf.CreateTagRequest)(string,error){
	tx, err :=  rep.SqlDriver.Begin()
	if err != nil{
		_ = rep.SqlDriver.RollBack(tx)
		return "", err
	}
	stmt, err := rep.SqlDriver.Prepare(tx, helpers.CreateTags)
	if err != nil{
		_ = rep.SqlDriver.RollBack(tx)
		return "", err
	}
	res, err := rep.SqlDriver.Exec(stmt, req.Name, req.Description)
	if err != nil{
		_ = rep.SqlDriver.RollBack(tx)
		return "", err
	}
	id, err := rep.SqlDriver.LastInsertedID(res)
	if err != nil{
		_ = rep.SqlDriver.RollBack(tx)
		return "", nil
	}
	_ = rep.SqlDriver.Commit(tx)
	return strconv.Itoa(int(id)) , nil
}

func (rep *UserRepository) GetUserByEmail(ctx context.Context,emailID string){
	//tx, err :=  rep.SqlDriver.Begin()


}

func (rep *UserRepository) CreateRole(ctx context.Context, req *pf.CreateRoleRequest)(string,error){
	tx, err :=  rep.SqlDriver.Begin()
	stmt, err := rep.SqlDriver.Prepare(tx, helpers.CreateRoles)
	if err != nil{
		_ = rep.SqlDriver.RollBack(tx)
		return "", err
	}
	res, err := rep.SqlDriver.Exec(stmt, req.Name, req.Description)
	if err != nil{
		return "", err
	}
	id, err := rep.SqlDriver.LastInsertedID(res)
	if err != nil{
		_ = rep.SqlDriver.RollBack(tx)
		return "", nil
	}
	_ = rep.SqlDriver.Commit(tx)
	return strconv.Itoa(int(id)) , nil
}

func (rep *UserRepository) CreatePermission(ctx context.Context, req *pf.CreatePermissionRequest)(string,error){

	tx, err :=  rep.SqlDriver.Begin()
	stmt, err := rep.SqlDriver.Prepare(tx, helpers.CreatePermission)
	if err != nil{
		_ = rep.SqlDriver.RollBack(tx)
		return "", err
	}
	res, err := rep.SqlDriver.Exec(stmt, req.Name, req.Description)
	if err != nil{
		return "", err
	}
	id, err := rep.SqlDriver.LastInsertedID(res)
	if err != nil{
		_ = rep.SqlDriver.RollBack(tx)
		return "", nil
	}
	_ = rep.SqlDriver.Commit(tx)
	return strconv.Itoa(int(id)) , nil
}

func (rep *UserRepository)  AttachRolesPermissions(ctx context.Context, req *pf.AttachRolesPermissionsRequest)(map[int64]int64){
	errors := make(map[int64]int64, 0)
	for index, permission := range req.Permissions{
		tx, err :=  rep.SqlDriver.Begin()
		stmt, err := rep.SqlDriver.Prepare(tx, helpers.CreateRolesPermissions)
		if err != nil{
			errors[req.Roles[index]] = permission
			_ = rep.SqlDriver.RollBack(tx)
			continue
		}
		result, err := rep.SqlDriver.Exec(stmt,req.Roles[index], permission)
		if err != nil{
			errors[req.Roles[index]] = permission
			_ = rep.SqlDriver.RollBack(tx)
			continue
		}
		_, err = rep.SqlDriver.LastInsertedID(result)
		if err != nil{
			errors[req.Roles[index]] = permission
			_ = rep.SqlDriver.RollBack(tx)
			continue
		}
		err = rep.SqlDriver.Commit(tx)
		if err != nil{
			errors[req.Roles[index]] = permission
			_ = rep.SqlDriver.RollBack(tx)
			continue
		}
	}

	return errors
}

func (rep *UserRepository)  AttachUsersRoles(ctx context.Context, req *pf.AttachUsersRolesRequest)map[string]string{
	errors := make(map[string]string, 0)
	for key, users := range req.Role {
		for _, value := range users.Users{
			tx, err :=  rep.SqlDriver.Begin()
			stmt, err := rep.SqlDriver.Prepare(tx, helpers.CreateUsersRoles)
			if err != nil{
				errors[value] = key
				_ = rep.SqlDriver.RollBack(tx)
				continue
			}
			result, err := rep.SqlDriver.Exec(stmt,value, key)
			if err != nil{
				errors[value] = key
				_ = rep.SqlDriver.RollBack(tx)
				continue
			}
			_, err = rep.SqlDriver.LastInsertedID(result)
			if err != nil{
				errors[value] = key
				_ = rep.SqlDriver.RollBack(tx)
				continue
			}
			err = rep.SqlDriver.Commit(tx)
			if err != nil{
				errors[value] = key
				_ = rep.SqlDriver.RollBack(tx)
				continue
			}
		}
	}

	return errors
}


type SqlInterface interface {
	Begin()(*sql.Tx, error)
	Prepare(tx *sql.Tx, query string)(*sql.Stmt, error)
	Exec(stmt *sql.Stmt,args ...interface{})(sql.Result, error)
	LastInsertedID(result sql.Result)(int64, error)
	RowEffected(result sql.Result)(int64, error)
	Commit(tx *sql.Tx)error
	RollBack(tx *sql.Tx)error
}

type SQLImplementation struct {
     sqlClient *sql.DB
}

func New(sql *sql.DB)SqlInterface{
	return &SQLImplementation{sqlClient:sql}
}

func (sql *SQLImplementation) Begin()(*sql.Tx, error){
	return sql.sqlClient.Begin()
}

func (sql *SQLImplementation) Prepare(tx *sql.Tx, query string)(*sql.Stmt, error){

	return tx.Prepare(query)
}

func (sql *SQLImplementation) Exec(stmt *sql.Stmt,args ...interface{})(sql.Result, error){
    return stmt.Exec(args...)
}

func (sql *SQLImplementation) LastInsertedID(result sql.Result)(int64, error){
	return result.LastInsertId()
}

func (sql *SQLImplementation) RowEffected(result sql.Result)(int64, error){
	return result.RowsAffected()
}

func (sql *SQLImplementation) Commit(tx *sql.Tx)error{
	return tx.Commit()
}

func (sql *SQLImplementation) RollBack(tx *sql.Tx)error{
	return tx.Rollback()
}



