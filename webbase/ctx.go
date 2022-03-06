package webbase

import (
	"time"

	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/virgoC0der/go-base/logging"
	"github.com/virgoC0der/go-base/valid"
)

type CommonResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ServeResponse(c *gin.Context, errMsg *CommonResp, args ...interface{}) {
	if nil == errMsg {
		return
	}

	data := errMsg.Data
	if len(args) > 0 {
		if nil != args[0] {
			data = args[0]
		}
	}

	result := &CommonResp{
		Code:    errMsg.Code,
		Message: errMsg.Message,
		Data:    data,
	}
	c.JSON(200, result)
}

// InitServer inits gin server
func InitServer() *gin.Engine {
	logging.InitLog()

	g := gin.New()
	g.Use(ginzap.Ginzap(logging.Logger, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(logging.Logger, true))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		valid.RegisterValidator(v)
	}

	return g
}
