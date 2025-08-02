package llm

import (
	"context"
	"github.com/cohesion-org/deepseek-go"
	"github.com/yumosx/poc/internal/domain"
)

type Handler struct {
	client *deepseek.Client
}

func NewHandler(token string, baseUrl string) *Handler {
	client := deepseek.NewClient(token, baseUrl)
	return &Handler{client: client}
}

func (h *Handler) Handle(ctx context.Context, req domain.LLMRequest) (string, error) {
	resp, err := h.client.CreateChatCompletion(ctx, &deepseek.ChatCompletionRequest{
		Messages: []deepseek.ChatCompletionMessage{
			{Role: deepseek.ChatMessageRoleSystem, Content: h.getSystemContent(req.Type)},
			{Role: deepseek.ChatMessageRoleUser, Content: req.Content},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func (h *Handler) Stream(ctx context.Context) {

}

// getSystemContent 一个简单的提示词
func (h *Handler) getSystemContent(ty string) string {
	switch ty {
	case domain.ZH2EN:
		return "将输入的中文翻译成为英文"
	case domain.EN2ZH:
		return "将输入的英文翻译成为中文"
	case domain.SUMMARIZE:
		return "对输入的文字进行总结"
	default:
	}
	return ""
}
