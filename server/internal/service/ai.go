package service

import (
	"context"
	"github.com/yumosx/poc/server/internal/domain"
	"github.com/yumosx/poc/server/internal/repo"
	"github.com/yumosx/poc/server/internal/service/llm"
	"github.com/yumosx/poc/server/internal/utils/logger"
	"runtime/debug"
	"time"
)

type AIService struct {
	handler *llm.Handler
	repo    *repo.AIRepo
}

func NewAIService(repo *repo.AIRepo, handler *llm.Handler) *AIService {
	return &AIService{repo: repo, handler: handler}
}

func (svc *AIService) RunTask(ctx context.Context, task domain.Task) (string, error) {
	task.State = domain.Padding
	id, err := svc.repo.SaveTask(ctx, task)
	if err != nil {
		return "", err
	}
	go func(uuid string) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("PANIC in AIService.RunTask goroutine: %v", err)
				logger.Errorf("Stack trace: %s", string(debug.Stack()))
			}
		}()
		innerCtx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancelFunc()

		content, err := svc.handler.Handle(innerCtx, domain.LLMRequest{Type: task.Type, Content: task.Content})
		state := domain.Running
		if err != nil {
			logger.Errorf("AIService|RunTask | %s: %v", uuid, err)
			state = domain.Failed
		} else {
			state = domain.Success
		}
		logger.Debugf("AIService|RunTask|%s|%s", content, state)

		dbCtx, cancelDb := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelDb()
		// 更新对应的结果
		_, err = svc.repo.SaveTask(dbCtx, domain.Task{UUID: uuid, Result: content, State: state})
		if err != nil {
			logger.Errorf("AIService|RunTask | %s: %v", uuid, err)
			return
		}
	}(id)
	return id, nil
}

// GetTask 获取 task 的 detail
func (svc *AIService) GetTask(ctx context.Context, id string) (domain.Task, error) {
	return svc.repo.GetTask(ctx, id)
}

// Stream 大模型的流式返回
func (svc *AIService) Stream(ctx context.Context, request domain.LLMRequest) (chan domain.StreamResponse, error) {
	return svc.handler.Stream(ctx, request)
}
