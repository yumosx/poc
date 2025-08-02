//go:build wireinject
// +build wireinject

package ioc

import (
	"github.com/google/wire"
	"github.com/yumosx/poc/server/internal/handler"
	"github.com/yumosx/poc/server/internal/repo"
	dao2 "github.com/yumosx/poc/server/internal/repo/dao"
	"github.com/yumosx/poc/server/internal/service"
)

func InitApp() *handler.Handler {
	wire.Build(
		initDB,
		dao2.NewAIDao,
		dao2.NewTaskDao,
		repo.NewAIRepo,
		initLLMHandler,
		service.NewAIService,
		handler.NewHandler)
	return nil
}
