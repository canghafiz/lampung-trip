package controller

import (
	"github.com/gin-gonic/gin"
	"lampung_trip/exception"
	"lampung_trip/helper"
	"lampung_trip/model/service"
	"lampung_trip/model/web"
	"net/http"
)

type FileContImpl struct {
	Serv service.FileServ
}

func NewFileContImpl(serv service.FileServ) *FileContImpl {
	return &FileContImpl{Serv: serv}
}

func (cont *FileContImpl) UploadFile(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		exception.ErrorHandler(context, err)
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		exception.ErrorHandler(context, "no files uploaded")
		return
	}

	uploadedFiles, errServ := cont.Serv.UploadFile(files)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   uploadedFiles,
	}

	if err := helper.WriteToResponseBody(context, api.Code, api); err != nil {
		exception.ErrorHandler(context, err)
		return
	}
}

func (cont *FileContImpl) DeleteFile(context *gin.Context) {
	url := context.Query("url")

	// Call Service
	err := cont.Serv.DeleteFile(url)
	if err != nil {
		exception.ErrorHandler(context, err)
		return
	}

	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	if err := helper.WriteToResponseBody(context, api.Code, api); err != nil {
		exception.ErrorHandler(context, err)
		return
	}
}
