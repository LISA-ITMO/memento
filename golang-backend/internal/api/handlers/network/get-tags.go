package network

import (
	"context"
	"github.com/gin-gonic/gin"
	"memento/internal/models"
	"net/http"
)

type GetTagsResp struct {
	Tags []models.Tag `json:"tags"`
}

func (h *Handler) GetAllTags(c *gin.Context) {
	res, err := h.s.GetTags(context.Background())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp := GetTagsResp{Tags: res.Tags}
	c.AbortWithStatusJSON(http.StatusOK, &resp)
}
