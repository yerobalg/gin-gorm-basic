package handler

import (
	"fmt"
	"gin-gorm-basic/sdk/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	http   *gin.Engine
	config config.Interface
	db     *gorm.DB
}

func Init(config config.Interface, db *gorm.DB) *handler {
	rest := handler{
		http:   gin.Default(),
		config: config,
		db:     db,
	}

	rest.registerRoutes()

	return &rest
}

func (h *handler) registerRoutes() {
	h.http.GET("/", h.ping)

	v1 := h.http.Group("api/v1")

	// Post
	v1.POST("/post", h.createPost)
	v1.GET("/post", h.getListPost)
	v1.GET("/post/:post_id", h.getPost)
	v1.PUT("/post/:post_id", h.updatePost) // 1 -> aku mau update post yang id nya 1
	v1.DELETE("/post/:post_id", h.deletePost)

	// Comment
	// v1.POST("/comment", h.createComment)
	// v1.GET("/comment", h.getListComment)
	// v1.GET("/comment/:comment_id", h.getComment)
	// v1.PUT("/comment/:comment_id", h.updateComment) // 1 -> aku mau update post yang id nya 1
	// v1.DELETE("/comment/:comment_id", h.deleteComment)
}

func (h *handler) ping(ctx *gin.Context) {
	h.SuccessResponse(ctx, http.StatusOK, "pong", nil, nil)
}

func (h *handler) Run() {
	h.http.Run(fmt.Sprintf(":%s", h.config.Get("PORT")))
}
