package handles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateNote implements Handles.
func (h *handles) CreateNote(c *gin.Context) {
	type jsonStructBody struct {
		Data string `json:"data"`
	}
	jsonBody := &jsonStructBody{}
	if err := c.BindJSON(jsonBody); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "internal server error"})
	}

	note, err := h.repo.CreateNote(jsonBody.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error2"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Your special code and note uniqueId": note.Code + " ||| " + note.Uid})
}

// ReadNote implements Handles.
func (h *handles) ReadNote(c *gin.Context) {
	type jsonStructBody struct {
		Uid string `json:"uid"`
	}
	jsonBody := &jsonStructBody{}
	if err := c.BindJSON(jsonBody); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "internal server error"})
	}

	code := c.Param("part1") + "/" + c.Param("part2")
	note, err := h.repo.ReadNote(code, jsonBody.Uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Your note data": note.Data})
}
