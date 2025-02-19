package api

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/IaC/go-kcloutie/docs"
)

var config *ServerConfiguration

func PathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false

}

func (c *ServerConfiguration) Start(ctx context.Context) error {

	DB, err := c.configureDB(ctx)
	if err != nil {
		return err
	}

	controller := NewController(DB)

	r := gin.Default()

	memoryStore := persist.NewMemoryStore(time.Duration(*c.CacheInSeconds) * time.Second)

	r.GET("", cache.CacheByRequestURI(memoryStore, time.Duration(*c.CacheInSeconds)*time.Second), func(c *gin.Context) {
		Home(ctx, c)
	})

	r.GET("/healthz", Health)
	r.GET("/livez", Liveness)
	r.GET("/readyz", Readyz)

	r.Use(PrometheusMiddleware())
	r.GET("/metrics", gin.WrapH(MetricsHandler()))

	apiV1 := r.Group("/api/v1")
	apiV1.Use()
	{
		apiV1.GET("/time", func(c *gin.Context) {
			controller.GetTime(ctx, c)
		})
		apiV1.POST("/time", func(c *gin.Context) {
			controller.PostTime(ctx, c)
		})
	}
	if c.DBEnabled() {
		apiV1.Use()
		{
			apiV1.GET("/widgets", func(c *gin.Context) {
				controller.GetAllWidgets(ctx, c)
			})

			apiV1.GET("/widgets/:id", func(c *gin.Context) {
				controller.GetWidget(ctx, c)
			})

			apiV1.POST("/widgets", func(c *gin.Context) {
				controller.CreateWidget(ctx, c)
			})

			apiV1.PUT("/widgets/:id", func(c *gin.Context) {
				controller.UpdateWidget(ctx, c)
			})

			apiV1.DELETE("/widgets/:id", func(c *gin.Context) {
				controller.DeleteWidget(ctx, c)
			})
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:              *c.ListeningAddr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           r,
	}

	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func NewServerConfiguration() *ServerConfiguration {
	return &ServerConfiguration{}
}

func FromCtx(ctx context.Context) *ServerConfiguration {
	if l, ok := ctx.Value(ctxConfigKey{}).(*ServerConfiguration); ok {
		return l
	} else if l := config; l != nil {
		return l
	}
	return NewServerConfiguration()
}

func WithCtx(ctx context.Context, l *ServerConfiguration) context.Context {
	if lp, ok := ctx.Value(ctxConfigKey{}).(*ServerConfiguration); ok {
		if lp == l {
			return ctx
		}
	}
	return context.WithValue(ctx, ctxConfigKey{}, l)
}
