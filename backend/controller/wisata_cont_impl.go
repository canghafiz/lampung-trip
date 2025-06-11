package controller

import (
	"github.com/gin-gonic/gin"
	"lampung_trip/exception"
	"lampung_trip/helper"
	"lampung_trip/model/service"
	"lampung_trip/model/web"
	"lampung_trip/model/web/wisata"
	"net/http"
	"strconv"
)

type WisataContImpl struct {
	Serv service.WisataServ
}

func NewWisataContImpl(serv service.WisataServ) *WisataContImpl {
	return &WisataContImpl{Serv: serv}
}

func (cont *WisataContImpl) Create(context *gin.Context) {
	// Parse Request Body
	var request wisata.CreateRequest
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Serv.Create(request)
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

func (cont *WisataContImpl) Update(context *gin.Context) {
	// Parse Request Body
	var request wisata.UpdateRequest
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Serv.Update(request)
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

func (cont *WisataContImpl) GetAll(context *gin.Context) {
	// Call Service
	servResult, errServ := cont.Serv.GetAll()
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   servResult,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *WisataContImpl) Delete(context *gin.Context) {
	id := context.Param("id")
	idFinal, _ := strconv.Atoi(id)

	// Call Service
	errServ := cont.Serv.Delete(idFinal)
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
