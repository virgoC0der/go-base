package main

import (
	"go.uber.org/zap"
	"os"
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	. "github.com/virgoC0der/go-base/logging"
	"github.com/virgoC0der/go-base/rbac"
	"github.com/virgoC0der/go-base/rbac/internal/handlers"
)

func main() {
	InitLog()

	if err := rbac.Init(); err != nil {
		Logger.Error("init rbac err", zap.Error(err))
		os.Exit(1)
	}

	r := gin.New()
	r.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(Logger, true))

	apiV1 := r.Group("api/v1")
	apiV1.Use(rbac.DomainsHandler())
	{
		apiV1.POST("/casbin", handlers.CasbinTest)
		apiV1.POST("/casbin/approve", handlers.CasbinTest)
	}

	r.Run(":8080")
}
