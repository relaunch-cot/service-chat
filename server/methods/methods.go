package methods

import (
	"github.com/relaunch-cot/service-chat/resource"
	"google.golang.org/grpc"

	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
)

func RegisterGrpcServices(s *grpc.Server) {
	pb.RegisterChatServiceServer(s, resource.Server.Chat)
}
