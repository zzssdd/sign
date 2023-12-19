package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"sign/conf"
	"sign/pkg/errmsg"
	"sign/utils"
	"strings"
)

func JwtMW() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.Abort()
			c.JSON(http.StatusOK, errmsg.GetErrMsg(errmsg.TokenNotExist))
			return
		}
		checkToken := strings.Split(auth, " ")
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, errmsg.GetErrMsg(errmsg.TokenFormatError))
			c.Abort()
			return
		}
		email, id, err := utils.ParseToken(checkToken[1], conf.GlobalConfig.JwtSecret)
		if err != nil {
			c.JSON(http.StatusOK, errmsg.GetErrMsg(errmsg.TokenInValid))
			c.Abort()
			return
		}
		c.Set("email", email)
		c.Set("id", id)
		c.Next(ctx)
	}
}
