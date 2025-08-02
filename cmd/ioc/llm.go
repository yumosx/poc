package ioc

import (
	"github.com/spf13/viper"
	"github.com/yumosx/poc/internal/service/llm"
)

func initLLMHandler() *llm.Handler {
	type Config struct {
		Token string `yaml:"token"`
	}
	config := Config{
		Token: viper.GetString("AI_TOKEN"),
	}
	return llm.NewHandler(config.Token)
}
