package main

import (
	"github.com/alitaso345-sandbox/newsfeeder/httpd/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	r.Run()
}