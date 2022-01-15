package v1

import (
	"github.com/ZimoBoy/gohub/app/models/user"
	"github.com/ZimoBoy/gohub/pkg/auth"
	"github.com/ZimoBoy/gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseApiController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	data := user.All()
	response.Data(c, data)
}
