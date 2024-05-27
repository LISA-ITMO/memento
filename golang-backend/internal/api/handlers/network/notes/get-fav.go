package notes

import (
	"github.com/gin-gonic/gin"
	"memento/internal/dto"
	"net/http"
)

type GetFavsResp struct {
}

func (h *Handler) GetFavourites(c *gin.Context) {
	user := c.GetInt("user_id")
	var _ = user

	in := dto.GetFavsIn{}
	res, err := h.s.GetFavourites(in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp := makeFavResp(res)
	c.JSON(http.StatusOK, resp)
}

func makeFavResp(in dto.GetFavsOut) GetFavsResp {
	return GetFavsResp{}
}
