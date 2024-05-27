package folders

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"memento/internal/dto"
	"net/http"
	"strconv"
)

func (h *Handler) Rename(c *gin.Context) {
	const op = "internal/api/handlers/folders/patch.go/Rename()"

	user := c.GetInt("user_id")
	folder, err := strconv.Atoi(c.Param("id"))
	newName := c.Param("name")

	if err != nil {
		log.Printf(op+": %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	in := makeRenameDTOIn(folder, user, newName)
	if err = h.s.RenameFolder(context.Background(), in); err != nil {
		log.Printf(op+": %v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func makeRenameDTOIn(folder, user int, newName string) dto.RenameFolderIn {
	return dto.RenameFolderIn{
		FolderID: folder,
		UserID:   user,
		NewName:  newName,
	}
}
