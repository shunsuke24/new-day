package app

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	go_openai "github.com/sashabaranov/go-openai"
)

func setupOpenAIClient() *go_openai.Client {
	client := go_openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	return client
}

func setupGinServer() *gin.Engine {
	server := gin.New()
	server.Use(gin.LoggerWithConfig(customLogger()))
	server.Use(gin.Recovery())

	return server
}

func customLogger() gin.LoggerConfig {
	conf := gin.LoggerConfig{
		SkipPaths: []string{"/health"},
	}
	conf.Formatter = customLogFormatter
	return conf
}

var customLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor,
		param.StatusCode,
		resetColor,
		param.Latency,
		param.ClientIP,
		methodColor,
		param.Method,
		resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
