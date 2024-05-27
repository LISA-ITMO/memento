package questions

import (
	"context"
	"github.com/gin-gonic/gin"
	"memento/internal/dto"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteQuestion(c *gin.Context) {
	user := c.GetInt("user_id")
	question, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	input := dto.DelQuestionIn{
		QuestionID: question,
		UserID:     user,
	}

	if err = h.s.Delete(context.Background(), input); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
