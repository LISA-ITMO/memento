package folders

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"memento/internal/dto"
	"memento/internal/models"
	"net/http"
)

type GetMainResp struct {
	Folders []models.Folder `json:"folders"`
	Notes   []models.Note   `json:"notes"`
}

func (h *Handler) GetMain(c *gin.Context) {
	const op = "internal/api/handlers/folders/get-main"
	user := c.GetInt("user_id")
	input := makeGetMainDTOIn(user)

	result, err := h.s.GetMainFolder(context.Background(), input)
	if err != nil {
		log.Printf(op+": %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp := makeGetMainResp(result)
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func makeGetMainDTOIn(user int) dto.GetMainFolderIn {
	return dto.GetMainFolderIn{UserID: user}
}

func makeGetMainResp(in dto.GetMainFolderOut) GetMainResp {
	return GetMainResp{
		Folders: in.Folders,
		Notes:   in.Notes,
	}
}
