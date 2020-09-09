package main

import (
	"github.com/sirupsen/logrus"
	"github.com/velann21/match_todo_user_srv/pkg/routes"
	"google.golang.org/grpc"
	//"github.com/velann21/todo-commonlib/commons/helpers"
	db "github.com/velann21/todo-commonlib/databases/mysql"
	pf "github.com/velann21/todo-commonlib/proto_files"
	"log"
	"net"
	"os"
)

func main() {
	//helpers.SetEnv()
	sqlconn := db.SQLConnection{}
	err := sqlconn.OpenSqlConnection()
	if err != nil {
		logrus.WithField("EventType", "DbConnection").WithError(err).Error("Db Connection Error")
		os.Exit(100)
	}

	grpcController := routes.Intialize(db.GetSqlConnection())
	server, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logrus.WithError(err).Error("failed to listen")
		os.Exit(100)
	}
	s := grpc.NewServer()
	pf.RegisterUserManagementServiceServer(s, &routes.ServerRoutes{GrpcController:grpcController})
	logrus.Info("Started user_srv grpc server")
	err = s.Serve(server)
	if err != nil {
		log.Fatal("Something wrong while booting up grpc")
	}
}
