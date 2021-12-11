/*
 * @Author: tangqimin
 * @Date: 2021-11-12 11:12:19
 * @Description:
 * @LastEditTime: 2021-12-11 17:18:59
 * @LastEditors: tangqimin
 * @FilePath: \study-gin\middleware\jwt\jwt.go
 */
package jwt

import (
	"StudyGin/pkg/e"
	"StudyGin/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		var data interface{}
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
