package v1

import (
	"github.com/ZimoBoy/gohub/app/models/link"
	"github.com/ZimoBoy/gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseApiController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	links := link.AllCached()
	response.Data(c, links)
}
