package controller

import "github.com/gin-gonic/gin"

type CustomerCont interface {
	UpdateData(context *gin.Context)
	GetSingleByUserId(context *gin.Context)
	GetAll(context *gin.Context)
}
