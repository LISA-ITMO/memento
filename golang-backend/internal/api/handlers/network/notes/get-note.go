package notes

import (
	"github.com/gin-gonic/gin"
	"memento/internal/dto"
	"memento/internal/models"
	"net/http"
	"strconv"
)

type GetNotesResp struct {
	Questions []models.Question
}

func (h *Handler) GetNotes(c *gin.Context) {
	tag, err := strconv.Atoi(c.Param("tag"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	input := dto.GetPublicNotesIn{Tag: tag}

	result, err := h.s.GetPublicNotes(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp := GetNotesResp{Questions: result.Questions}
	c.JSON(http.StatusOK, &resp)
}

