package rbac

import (
	"github.com/gin-gonic/gin"
	. "github.com/virgoC0der/go-base/logging"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		// TODO: read role from session
		sub := "admin"
		success, err := Enforcer(sub, obj, act)
		if err != nil {
			Logger.Warn("enforce err")
			return
		}

		if success {
			c.Next()
			return
		}

		Logger.Warn("无权限访问")
		resp := struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{
			1001,
			"no access",
		}
		c.AbortWithStatusJSON(200, &resp)
	}
}
