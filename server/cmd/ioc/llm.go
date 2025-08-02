package ioc

import (
	"github.com/spf13/viper"
	"github.com/yumosx/poc/server/internal/service/llm"
	"github.com/yumosx/poc/server/internal/utils/logger"
	"os"
)

func initLLMHandler() *llm.Handler {
	type Config struct {
		Token string `yaml:"token"`
	}
	config := Config{
		Token: viper.GetString("AI_TOKEN"),
	}
	config.Token = os.Getenv("AI_TOKEN")
	logger.Debug(config.Token)
	return llm.NewHandler(config.Token)
}
