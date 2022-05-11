package handlers

import "github.com/gin-gonic/gin"

func CasbinTest(c *gin.Context) {
	resp := struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{
		0,
		"success",
	}
	c.JSON(200, &resp)
}
