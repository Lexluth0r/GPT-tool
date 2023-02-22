package route

import "github.com/gin-gonic/gin"

func WebRouter(r *gin.Engine) {

	r.Static("/assets", "./resource/dist/assets")
	r.StaticFile("/", "./resource/dist/index.html")
	r.StaticFile("/favicon.ico", "./resource/dist/favicon.ico")
}
