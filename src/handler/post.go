package handler

import (
	"gin-gorm-basic/src/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) createPost(ctx *gin.Context) {
	var postBody entity.PostBody

	if err := h.BindBody(ctx, &postBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	var post entity.Post

	post.Title = postBody.Title
	post.Content = postBody.Content

	if err := h.db.Create(&post).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "Successfully created new post", nil)
}
