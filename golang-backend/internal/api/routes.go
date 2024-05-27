package api

import "memento/internal/api/handlers/folders"

func (s *Server) setupRoutes() {
	foldersGroup := s.router.Group("/:id")
	{
		h := folders.New(s.storage)
		foldersGroup.GET("", h.Get)
		foldersGroup.POST("", h.CreateObject)
		foldersGroup.PATCH("", h.Rename)
		foldersGroup.DELETE("", h.Delete)
	}
}

