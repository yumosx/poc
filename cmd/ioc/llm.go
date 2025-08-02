package ioc

import "github.com/yumosx/poc/internal/service/llm"

func initLLMHandler() *llm.Handler {
	type Config struct {
		Token   string `yaml:"token"`
		BaseURL string `yaml:"baseURL"`
	}
	var cfg Config
	return llm.NewHandler(cfg.Token, cfg.BaseURL)
}
