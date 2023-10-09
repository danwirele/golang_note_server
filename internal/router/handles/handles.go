package handles

import (
	"notes_server/internal/domain/note"

	"github.com/gin-gonic/gin"
)

type handles struct {
	repo note.Repository
}

type Handles interface {
	CreateNote(*gin.Context)
	ReadNote(*gin.Context)
}

func NewHandles(repo note.Repository) Handles {
	return &handles{repo: repo}
}

func ImplHandles(ge *gin.Engine, h Handles) {
	ge.POST("/notes", h.CreateNote)
	ge.GET("/notes/:part1/:part2", h.ReadNote)
}
