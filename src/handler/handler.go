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
}

func (h *handler) ping(ctx *gin.Context) {
	h.SuccessResponse(ctx, http.StatusOK, "pong", nil, nil)
}

func (h *handler) Run() {
	h.http.Run(fmt.Sprintf(":%s", h.config.Get("PORT")))
}
