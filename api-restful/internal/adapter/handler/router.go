package handler

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(heroHandler HeroHandler) (*Router, error) {
	// Disable debug mode and write logs to file in production environment
	if env := os.Getenv("APP_ENV"); env == "production" {
		gin.SetMode(gin.ReleaseMode)
		logFile, _ := os.Create("gin.log")
		gin.DefaultWriter = io.Writer(logFile)
	}

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery())

	// Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		hero := v1.Group("/heroes")
		{
			hero.GET("", heroHandler.GetHeroes)
			hero.GET("/:id", heroHandler.GetHeroById)
		}
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(httpUrl string) error {
	return r.Run(httpUrl)
}

// customLogger is a custom Gin logger
func customLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s [%s]\"\n",
		param.TimeStamp.Format(time.RFC1123),
		param.ClientIP,
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency.Round(time.Millisecond),
		param.Request.UserAgent(),
	)
}
