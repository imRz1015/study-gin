/*
 * @Author: tangqimin
 * @Date: 2021-11-08 16:23:08
 * @Description:
 * @LastEditTime: 2021-12-11 17:18:26
 * @LastEditors: tangqimin
 * @FilePath: \study-gin\main.go
 */
package main

import (
	"StudyGin/pkg/setting"
	"StudyGin/routers"
	"fmt"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
