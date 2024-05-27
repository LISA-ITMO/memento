package notes

import (
	"context"
	"github.com/gin-gonic/gin"
	"memento/internal/dto"
	"net/http"
	"strconv"
)

func (h *Handler) RenameNote(c *gin.Context) {
	user := c.GetInt("user_id")
	note, err := strconv.Atoi(c.Param("id"))
	newName := c.Param("name")

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	input := dto.RenameNoteIn{
		NoteID:  note,
		UserID:  user,
		NewName: newName,
	}

	if err := h.s.Rename(context.Background(), input); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
