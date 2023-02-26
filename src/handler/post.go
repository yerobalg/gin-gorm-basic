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

	h.SuccessResponse(ctx, http.StatusOK, "Successfully created new post", nil, nil)
}

func (h *handler) getListPost(ctx *gin.Context) {
	var postParam entity.PostParam

	if err := h.BindParam(ctx, &postParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	postParam.FormatPagination()

	var posts []entity.Post

	if err := h.db.
		Model(entity.Post{}).
		Limit(int(postParam.Limit)).
		Offset(int(postParam.Offset)).
		Find(&posts).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var totalElements int64

	if err := h.db.
		Model(entity.Post{}).
		Limit(int(postParam.Limit)).
		Offset(int(postParam.Offset)).
		Count(&totalElements).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	postParam.ProcessPagination(totalElements)

	h.SuccessResponse(ctx, http.StatusOK, "Successfully get list post", posts, &postParam.PaginationParam)
}
