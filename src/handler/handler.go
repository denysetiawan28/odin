package handler

import (
	"odin/src/server/container"
)

type Handler struct {
}

func InitializeHandler(container *container.DefaultContainer) *Handler {
	return &Handler{}
}
