package handler

import (
	"gin-gorm-basic/src/business/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *handler) SuccessResponse(c *gin.Context, code int64, message string, data interface{}) {
	c.JSON(int(code), entity.HTTPResponse{
		Message:   message,
		IsSuccess: true,
		Data:      data,
	})
}

func (h *handler) ErrorResponse(c *gin.Context, code int64, message string, data interface{}) {
	c.JSON(int(code), entity.HTTPResponse{
		Message:   message,
		IsSuccess: false,
		Data:      data,
	})
}

func (h *handler) BindBody(c *gin.Context, body interface{}) interface{} {
	return c.ShouldBindWith(body, binding.JSON)
}
