package controller

import (
	"github.com/gin-gonic/gin"
)

type UserCont interface {
	LoginAdminKaryawan(context *gin.Context)
	Login(context *gin.Context)
	UpdatePw(context *gin.Context)
	RegisterAdmin(context *gin.Context)
	RegisterKaryawan(context *gin.Context)
	RegisterCustomer(context *gin.Context)
	DeleteKaryawan(context *gin.Context)
	DeleteAdmin(context *gin.Context)
}
