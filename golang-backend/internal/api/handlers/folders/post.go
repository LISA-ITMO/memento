package folders

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"memento/internal/dto"
	"net/http"
	"strconv"
)

type CreateObjectReq struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) CreateObject(c *gin.Context) {
	const op = "internal/api/handlers/folders/post.go/CreateObject()"

	// Парсим запрос
	user := c.GetInt("user_id")
	object := c.GetHeader("object_type")
	folder, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf(op+": %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Парсим тело запроса
	var req CreateObjectReq
	_ = c.BindJSON(&req)

	// Идем к БД
	switch object {
	case "folder":
		in := makeCreateFolderDTOIn(folder, user, req.Name)
		err = h.s.CreateFolder(context.Background(), in)
		if err != nil {
			log.Printf(op+": %v", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	case "note":
		in := makeCreateNoteDTOIn(folder, user, req.Name)
		err = h.s.CreateNote(context.Background(), in)
		if err != nil {
			log.Printf(op+": %v", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	default:
		c.JSON(http.StatusBadRequest, "bullshit")
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func makeCreateFolderDTOIn(parentFolderID, userID int, name string) dto.CreateFolderIn {
	return dto.CreateFolderIn{
		ParentFolderID: parentFolderID,
		UserID:         userID,
		Name:           name,
	}
}

func makeCreateNoteDTOIn(parentFolderID, userID int, name string) dto.CreateNoteIn {
	return dto.CreateNoteIn{
		ParentFolderID: parentFolderID,
		UserID:         userID,
		Name:           name,
	}
}
