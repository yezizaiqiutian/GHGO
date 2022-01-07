package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	test := r.Group("test")
	{
		test.GET("/get", get)
		test.POST("/post", post)
	}
	r.Run(":8888")
}

func post(context *gin.Context) {
	index :=context.PostForm("index")
	context.JSON(http.StatusOK,"post请求 index:"+index)
}

func get(context *gin.Context) {
	index, _ :=context.GetQuery("index")
	context.JSON(http.StatusOK,"get请求 index:"+index)
}