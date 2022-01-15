package v1

import (
	"github.com/ZimoBoy/gohub/app/models/category"
	"github.com/ZimoBoy/gohub/app/requests"
	"github.com/ZimoBoy/gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseApiController
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}
