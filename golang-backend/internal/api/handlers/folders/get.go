package folders

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"memento/internal/dto"
	"memento/internal/models"
	"net/http"
	"strconv"
)

type GetFolderResponse struct {
	Folders []models.Folder `json:"folders"`
	Notes   []models.Note   `json:"notes"`
}

func (h *Handler) Get(c *gin.Context) {
	const op = "internal/api/handlers/folders/Get"
	user := c.GetInt("user_id")
	folder, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf(op+": %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	input := makeGetFolderDTOIn(folder, user)
	result, err := h.s.GetFolder(context.Background(), input)
	if err != nil {
		log.Printf(op+": %v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp := makeResp(result)
	c.JSON(http.StatusOK, resp)
	return
}

func makeGetFolderDTOIn(folderID, userID int) dto.GetFolderIn {
	return dto.GetFolderIn{
		ID:     folderID,
		UserID: userID,
	}
}

func makeResp(in dto.GetFolderOut) GetFolderResponse {
	return GetFolderResponse{
		Folders: in.Folders,
		Notes:   in.Notes,
	}
}
