package route

import (
	apiController "chat/app/http/controller/api"
	"github.com/gin-gonic/gin"
)

func APIRouter(r *gin.Engine) {

	api := r.Group("/api")

	api.POST("/chat", apiController.Chat)
}
