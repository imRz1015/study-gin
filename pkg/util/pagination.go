/*
 * @Author: tangqimin
 * @Date: 2021-11-09 11:25:45
 * @Description:
 * @LastEditTime: 2021-12-11 17:20:04
 * @LastEditors: tangqimin
 * @FilePath: \study-gin\pkg\util\pagination.go
 */
package util

import (
	"StudyGin/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
