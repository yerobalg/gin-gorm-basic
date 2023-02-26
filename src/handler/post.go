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

	//find related categories
	var categories []entity.Category
	if err := h.db.Find(&categories, postBody.CategoriesID).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "categories not found", nil)
		return
	}

	if err := h.db.Create(&post).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	//add categories for post
	if err := h.db.Model(&post).Association("Categories").Append(categories); err != nil{
		h.ErrorResponse(ctx, http.StatusInternalServerError, "categories not added", nil)
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

func (h *handler) getPost(ctx *gin.Context){
	var postParam entity.PostParam

	if err := h.BindParam(ctx, &postParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	// var PostPayload entity.PostBody
	// if err := h.BindBody(ctx, &PostPayload); err != nil {
	// 	h.ErrorResponse(ctx, http.StatusBadRequest, "bad request", nil)
	// 	return
	// } 

	var post entity.Post 
	if err := h.db.Model(&post).Where(&postParam).First(&post).Error; err != nil{
		h.ErrorResponse(ctx, http.StatusInternalServerError, "couldn't get post", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "post found", post, nil)
}

func (h *handler) updatePost(ctx *gin.Context) {
	var postParam entity.PostParam
	if err := h.BindParam(ctx, &postParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	var postBody entity.PostBody
	if err := h.BindBody(ctx, &postBody); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad request", nil)
		return
	}

	var post entity.Post
	post.ID = uint(postParam.PostID)
	post.Title = postBody.Title
	post.Content = postBody.Content

	if err := h.db.Model(post).Where(postParam).Updates(&post).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "post update failed", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "update post success", post, nil)
}

func (h *handler) deletePost(ctx *gin.Context) {
	var postParam entity.PostParam
	if err := h.BindParam(ctx, &postParam); err != nil {
		h.ErrorResponse(ctx, http.StatusBadRequest, "bad param", nil)
		return
	}

	var post entity.Post
	post.ID = uint(postParam.PostID)
	if err := h.db.Delete(&post).Error; err != nil {
		h.ErrorResponse(ctx, http.StatusInternalServerError, "delete post failed", nil)
		return
	}

	h.SuccessResponse(ctx, http.StatusOK, "delete post success", nil, nil)
}