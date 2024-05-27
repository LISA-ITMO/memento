package folders

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"memento/internal/dto"
	"net/http"
	"strconv"
)

func (h *Handler) Delete(c *gin.Context) {
	const op = "internal/api/handlers/folders/delete.go/Delete()"
	user := c.GetInt("user_id")
	folder, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf(op+": %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	in := makeDelFolderDTOIn(folder, user)
	if err = h.s.DeleteFolder(context.Background(), in); err != nil {
		log.Printf(op+": %v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func makeDelFolderDTOIn(folderID, userID int) dto.DeleteFolderIn {
	return dto.DeleteFolderIn{FolderID: folderID, UserID: userID}
}
