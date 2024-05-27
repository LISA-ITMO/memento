package notes

import (
	"context"
	"github.com/gin-gonic/gin"
	"memento/internal/dto"
	"net/http"
	"strconv"
)

type CreateQuestionReq struct {
	Text   string `json:"text" binding:"required"`
	Answer string `json:"answer" binding:"required"`
}

func (h *Handler) CreateQuestion(c *gin.Context) {
	// Парсим необходимые данные из запроса
	user := c.GetInt("user_id")
	note, err := strconv.Atoi(c.Param("id"))
	var _, _ = user, note
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Извлекаем тело запроса
	var req CreateQuestionReq
	_ = c.BindJSON(&req)

	// Обращаемся к БД
	input := dto.CreateQuestionIn{
		NoteID: note,
		UserID: user,
		Body:   req.Text,
		Answer: req.Answer,
	}

	err = h.s.CreateQuestion(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
