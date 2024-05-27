package notes

import (
	"github.com/gin-gonic/gin"
	"memento/internal/dto"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteFav(c *gin.Context) {
	user := c.GetInt("user_id")
	note, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var _, _ = user, note

	input := dto.DeleteFavIn{}
	if err := h.s.DeleteFav(input); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
