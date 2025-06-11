package controller

import (
	"github.com/gin-gonic/gin"
	"lampung_trip/exception"
	"lampung_trip/helper"
	"lampung_trip/model/service"
	"lampung_trip/model/web"
	"lampung_trip/model/web/karyawan"
	"net/http"
	"strconv"
)

type KaryawanContImpl struct {
	Serv service.KaryawanServ
}

func NewKaryawanContImpl(serv service.KaryawanServ) *KaryawanContImpl {
	return &KaryawanContImpl{Serv: serv}
}

func (cont *KaryawanContImpl) UpdateData(context *gin.Context) {
	// Parse Request Body
	var request karyawan.UpdateDataRequest
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Serv.UpdateData(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *KaryawanContImpl) GetSingleByUserId(context *gin.Context) {
	id := context.Param("id")
	idFinal, _ := strconv.Atoi(id)

	// Call Service
	results, errServ := cont.Serv.GetSingleByUserId(idFinal)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   results,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *KaryawanContImpl) GetAll(context *gin.Context) {
	// Call Service
	results, errServ := cont.Serv.GetAll()
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   results,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}
