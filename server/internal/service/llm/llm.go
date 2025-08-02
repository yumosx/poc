package llm

import (
	"context"
	"errors"
	"github.com/cohesion-org/deepseek-go"
	"github.com/yumosx/poc/internal/domain"
	"io"
)

type Handler struct {
	client *deepseek.Client
}

func NewHandler(token string) *Handler {
	client := deepseek.NewClient(token)
	return &Handler{client: client}
}

func (h *Handler) Handle(ctx context.Context, req domain.LLMRequest) (string, error) {
	resp, err := h.client.CreateChatCompletion(ctx, &deepseek.ChatCompletionRequest{
		Model: deepseek.DeepSeekChat,
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

// Stream 实现大模型的流式调用
func (h *Handler) Stream(ctx context.Context, req domain.LLMRequest) (chan domain.StreamResponse, error) {
	request := deepseek.StreamChatCompletionRequest{
		Model: deepseek.DeepSeekChat,
		Messages: []deepseek.ChatCompletionMessage{
			{Role: deepseek.ChatMessageRoleSystem, Content: h.getSystemContent(req.Type)},
			{Role: deepseek.ChatMessageRoleUser, Content: req.Content},
		},
		Stream: true,
	}
	ch := make(chan domain.StreamResponse, 10)
	stream, err := h.client.CreateChatCompletionStream(ctx, &request)
	if err != nil {
		return nil, err
	}
	events := make(chan domain.StreamResponse, 10)
	go func() {
		defer close(events)
		h.recv(events, stream)
	}()
	return ch, nil
}

func (h *Handler) recv(eventCh chan domain.StreamResponse, stream deepseek.ChatCompletionStream) {
	for {
		chunk, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				eventCh <- domain.StreamResponse{Done: true}
				break
			}
			eventCh <- domain.StreamResponse{Err: err}
		}
		eventCh <- domain.StreamResponse{
			Content: chunk.Choices[0].Delta.Content,
			Err:     nil,
		}
	}
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
