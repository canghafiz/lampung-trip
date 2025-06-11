package controller

import "github.com/gin-gonic/gin"

type UlasanVillaCont interface {
	Create(context *gin.Context)
	Delete(context *gin.Context)
}
