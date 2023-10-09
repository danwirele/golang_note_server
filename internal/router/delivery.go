package router

import (
	"net/http"
	"notes_server/internal/domain/note"
	"notes_server/internal/router/handles"

	"github.com/gin-gonic/gin"
)

func RouterProvider(repo note.Repository) *gin.Engine {
	engine := NewRouter()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	handler := handles.NewHandles(repo)
	handles.ImplHandles(engine, handler)

	return engine
}
