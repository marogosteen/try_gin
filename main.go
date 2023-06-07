package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("site/templates/*")
	router.GET("/noma", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name":    "noma",
			"message": "終わりだよ。",
		})
	})
	router.GET("/sugimoto", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name":    "杉本",
			"message": "泣いてるよ。",
		})
	})
	router.Run()
}
