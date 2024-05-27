package questions

import (
	"context"
	"github.com/gin-gonic/gin"
	"memento/internal/dto"
	"net/http"
	"strconv"
)

type UPDQuestionReq struct {
	Text   string `json:"text"`
	Answer string `json:"answer"`
}

func (h *Handler) UpdateQuestion(c *gin.Context) {
	// Парсим запрос
	user := c.GetInt("user_id")
	question, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Парсим тело запроса
	var req UPDQuestionReq
	_ = c.BindJSON(&req)

	// Идем в БД
	input := dto.UPDQuestionIn{
		QuestionID: question,
		UserID:     user,
		Text:       req.Text,
		Answer:     req.Answer,
	}

	if err = h.s.Update(context.Background(), input); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
