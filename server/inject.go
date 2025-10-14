package server

import (
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
	"github.com/relaunch-cot/service-chat/handler"
)

type Servers struct {
	Chat pb.ChatServiceServer
}

func (s *Servers) Inject(handler *handler.Handlers) {
	s.Chat = NewChatServer(handler)
}
