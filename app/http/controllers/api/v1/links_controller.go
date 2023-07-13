package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lihc520/gohub/app/models/link"
	"github.com/lihc520/gohub/pkg/response"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	response.Data(c, link.AllCached())
}
