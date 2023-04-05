package main

import (
	"HexArithmatics/internal/adapters/app/api"
	"HexArithmatics/internal/adapters/core/arithmetic"
	gRPC "HexArithmatics/internal/adapters/framework/left/grpc"
	"HexArithmatics/internal/adapters/framework/right/db"
	"HexArithmatics/internal/ports"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	var err error

	//Ports
	var dbaseAdapter ports.DBPort
	var core ports.ArthimeticsPort
	var appAdapter ports.APIPort
	var gRPCAdpter ports.GRPCPort

	dBaseDrive := os.Getenv("DB_DRIVER")
	dSourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dBaseDrive, dSourceName)
	if err != nil {
		logrus.Fatalf("Failed to intitate database connections: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdpter = gRPC.NewAdapter(appAdapter)

	gRPCAdpter.Run()

}
