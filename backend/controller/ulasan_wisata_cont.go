package controller

import "github.com/gin-gonic/gin"

type UlasanWisataCont interface {
	Create(context *gin.Context)
	Delete(context *gin.Context)
}
