package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/category"
	"gohub/app/requests"
	"gohub/pkg/response"
)

type CategoriesController struct {
	BaseApiController
}

func (ctrl *CategoriesController) Index(c *gin.Context) {
	categories := category.All()
	response.Data(c, categories)
}

func (ctrl *CategoriesController) Show(c *gin.Context) {
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, categoryModel)
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
