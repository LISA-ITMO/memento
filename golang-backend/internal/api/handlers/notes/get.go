package notes

import (
	"context"
	"github.com/gin-gonic/gin"
	"memento/internal/dto"
	"net/http"
	"strconv"
)

type GetNoteResp struct{}

func (h *Handler) GetNote(c *gin.Context) {
	user := c.GetInt("user_id")
	note, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	input := dto.GetNoteIn{
		NoteID: note,
		UserID: user,
	}

	result, err := h.s.Get(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp := makeResp(result)
	c.JSON(http.StatusOK, resp)
}

func makeResp(in dto.GetNoteOut) GetNoteResp {
	return GetNoteResp{}
}
