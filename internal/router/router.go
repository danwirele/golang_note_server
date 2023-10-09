package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	engine := gin.Default()
	return engine
}
