package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handler(req *http.Request, w *http.Response) {

}

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})
	router.Run(":8000")
}
