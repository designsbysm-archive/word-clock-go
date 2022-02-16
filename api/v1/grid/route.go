package grid

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/grid")
	{
		group.GET("/", grid)
	}
}
