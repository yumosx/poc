package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yumosx/poc/cmd/ioc"
)

func main() {
	handler := ioc.InitApp()
	engine := gin.Default()
	handler.Route(engine)
	engine.Run("localhost:8080")
}
