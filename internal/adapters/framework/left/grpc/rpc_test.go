package rpc

import (
	"HexArithmatics/internal/adapters/framework/left/grpc/pb"
	"HexArithmatics/internal/adapters/framework/right/db"
	"HexArithmatics/internal/application/api"
	"HexArithmatics/internal/application/core/arithmetic"
	"context"
	"net"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const buffsize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error
	lis = bufconn.Listen(buffsize)
	grpcServer := grpc.NewServer()

	dBaseDrive := os.Getenv("DB_DRIVER")
	dSourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err := db.NewAdapter(dBaseDrive, dSourceName)
	if err != nil {
		logrus.Fatalf("Failed to intitate database connections: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	// core
	core := arithmetic.New()

	applicationAPI := api.NewApplication(dbaseAdapter, core)

	gRPCAdpter := NewAdapter(applicationAPI)

	grpcServer = grpc.NewServer()

	pb.RegisterArtihmeticServiceServer(grpcServer, gRPCAdpter)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Fatalf("Test server start error: %v", err)
		}
	}()
}

func bufDailer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDailer), grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("Failed to Dial bufnet: %v", err)
	}
	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArtihmeticServiceClient(conn)
	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetAddition(ctx, params)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
	}
	require.Equal(t, answer.Value, int32(2))
}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArtihmeticServiceClient(conn)
	params := &pb.OperationParameters{
		A: 2,
		B: 1,
	}

	answer, err := client.GetSubstraction(ctx, params)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
	}
	require.Equal(t, answer.Value, int32(1))
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArtihmeticServiceClient(conn)
	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetMultiplication(ctx, params)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
	}
	require.Equal(t, answer.Value, int32(1))
}

func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArtihmeticServiceClient(conn)
	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetDivision(ctx, params)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
	}
	require.Equal(t, answer.Value, int32(1))
}
