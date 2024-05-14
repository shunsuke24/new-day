package router

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/new-day/controller"
)

type Router struct {
	Engine     *gin.Engine
	controller controller.Controller
}

type Params struct {
	Engine     *gin.Engine
	Controller controller.Controller
}

func NewRouter(p *Params) *Router {
	return &Router{
		Engine:     p.Engine,
		controller: p.Controller,
	}
}

func (r *Router) Route() {
	r.Engine.Use(errorMiddleware())
	r.Engine.GET("/health", r.controller.HealthCheck)
	r.Engine.GET("/recommend", r.controller.Recommend)
	r.Engine.POST("/send", r.controller.Send)
}

func (r *Router) CORS() {
	// allowOriginList := strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	r.Engine.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{"*"},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS", // preflightに必要
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Accept-Encoding",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Origin",
			"Access-Control-Expose-Headers",
			"Access-Control-Max-Age",
			"Access-Control-Request-Headers",
			"Access-Control-Request-Method",
			"Authorization",
			"Content-Length",
			"Content-Type",
			"X-GPT4",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
}

func errorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			log.Print(err.Err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Error": err.Error(),
			})
		}
	}
}
