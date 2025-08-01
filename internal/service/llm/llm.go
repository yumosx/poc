package llm

import (
	"context"
	"github.com/cohesion-org/deepseek-go"
)

type Handler struct {
	client *deepseek.Client
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context) (string, error) {
	return "", nil
}

func (h *Handler) Stream(ctx context.Context) {

}
