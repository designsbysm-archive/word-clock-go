package v1

import (
	"github.com/designsbysm/word-clock-server-go/api/v1/grid"
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	group := r.Group("/v1")
	{
		grid.AddRoute(group)
	}
}
