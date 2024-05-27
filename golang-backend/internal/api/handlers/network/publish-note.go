package network

import (
	"context"
	"github.com/gin-gonic/gin"
	"memento/internal/dto"
	"net/http"
	"strconv"
)

type PublishNoteReq struct {
	Name  string `json:"name"`
	Descr string `json:"descr"`
	Tags  []int  `json:"tags"`
}

func (h *Handler) PublishNote(c *gin.Context) {
	user := c.GetInt("user_id")
	note, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var req PublishNoteReq
	err = c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	input := dto.PublishNoteIn{
		UserID: user,
		NoteID: note,
		Name:   req.Name,
		Descr:  req.Descr,
		Tags:   req.Tags,
	}

	if err = h.s.PublishNote(context.Background(), input); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
