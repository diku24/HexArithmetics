package main

import (
	gRPC "HexArithmatics/internal/adapters/framework/left/grpc"
	"HexArithmatics/internal/adapters/framework/right/db"
	"HexArithmatics/internal/application/api"
	"HexArithmatics/internal/application/core/arithmetic"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	var err error

	dBaseDrive := os.Getenv("DB_DRIVER")
	dSourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err := db.NewAdapter(dBaseDrive, dSourceName)
	if err != nil {
		logrus.Fatalf("Failed to intitate database connections: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	//core
	core := arithmetic.New()

	applicationAPI := api.NewApplication(dbaseAdapter, core)

	gRPCAdpter := gRPC.NewAdapter(applicationAPI)

	gRPCAdpter.Run()

}
