package rpc

import (
	"HexArithmatics/internal/adapters/framework/left/grpc/pb"
	"HexArithmatics/internal/ports"
	"context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	api ports.APIPort
}

// GetAddition implements ports.GRPCPort
func (*Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	panic("unimplemented")
}

// GetAddtion implements pb.ArtihmeticServiceServer
func (grpca Adapter) GetAddtion(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	//panic("unimplemented")
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, MISSING_VALUE_ERROR)

	}
	answer, err := grpca.api.GetAddtion(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, UNEXPECTED_ERROR)

	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		logrus.Fatalf("Failed to listen on Port 9000: %v", err)
	}
	arithmeticServiceServer := grpca

	grpcServer := grpc.NewServer()
	//pb.RegisterArtihmeticServiceServer(grpcServer, arithmeticServiceServer)
	pb.RegisterArtihmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		logrus.Fatalf("Failed to Serve gRPC server over port 9000: %v", err)
	}

}
