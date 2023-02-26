package handler

import (
	"fmt"
	"gin-gorm-basic/sdk/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	http   *gin.Engine
	config config.Interface
}

func Init(config config.Interface) *handler {
	rest := handler{
		http:   gin.Default(),
		config: config,
	}

	rest.registerRoutes()

	return &rest
}

func (h *handler) registerRoutes() {
	h.http.GET("/", h.ping)
}

func (h *handler) ping(ctx *gin.Context) {
	h.SuccessResponse(ctx, http.StatusOK, "pong", nil)
}

func (h *handler) Run() {
	h.http.Run(fmt.Sprintf(":%s", h.config.Get("PORT")))
}
