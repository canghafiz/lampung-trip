package controller

import "github.com/gin-gonic/gin"

type FileCont interface {
	UploadFile(context *gin.Context)
	DeleteFile(context *gin.Context)
}
