package rpc

import (
	"HexArithmatics/internal/adapters/framework/left/grpc/pb"
	"HexArithmatics/internal/ports"
	context "context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

// GetAddtion implements pb.ArtihmeticServiceServer
func (Adapter) GetAddtion(context.Context, *pb.OperationParameters) (*pb.Answer, error) {
	panic("unimplemented")
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
