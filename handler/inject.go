package handler

import "github.com/relaunch-cot/service-chat/repositories"

type Handlers struct {
	Chat IChatHandler
}

func (h *Handlers) Inject(repositories *repositories.Repositories) {
	h.Chat = NewChatHandler(repositories)
}
