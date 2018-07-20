package messaging

import (
	"log"
	"net"

	pb "github.com/mahesh-ufo/test_project/messaging/schema"
	taskManager "github.com/mahesh-ufo/test_project/task_manager"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TestRPC implements TestMessagingServer
func (server *GRPCServer) TestRPC(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	//server.taskManager.RunTask()
	return &pb.Reply{Message: "Server responding for the [" + in.Name + "]."}, nil
}

//SetTaskManager : SetTaskManager
func (server *GRPCServer) SetTaskManager(myTaskManager taskManager.TMInterface) error {
	server.taskManager = myTaskManager
	return nil
}

//Start : Start the server
func (server *GRPCServer) Start(ip string, port string) {
	var lis, err = net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: % v", err)
	}

	var s = grpc.NewServer()
	pb.RegisterAdvtEngineMessagingServer(s, server)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: % v", err)
	}

}
