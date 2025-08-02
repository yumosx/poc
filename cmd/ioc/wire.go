//go:build wireinject
// +build wireinject

package ioc

import (
	"github.com/google/wire"
	"github.com/yumosx/poc/internal/handler"
	"github.com/yumosx/poc/internal/repo"
	"github.com/yumosx/poc/internal/repo/dao"
	"github.com/yumosx/poc/internal/service"
)

func InitApp() *handler.Handler {
	wire.Build(
		initDB,
		dao.NewAIDao,
		dao.NewTaskDao,
		repo.NewAIRepo,
		initLLMHandler,
		service.NewAIService,
		handler.NewHandler)
	return nil
}
