package controller

import (
	"github.com/gin-gonic/gin"
	"lampung_trip/exception"
	"lampung_trip/helper"
	"lampung_trip/model/service"
	"lampung_trip/model/web"
	"lampung_trip/model/web/villa"
	"net/http"
)

type UlasanVillaContImpl struct {
	Serv service.UlasanVillaServ
}

func NewUlasanVillaContImpl(serv service.UlasanVillaServ) *UlasanVillaContImpl {
	return &UlasanVillaContImpl{Serv: serv}
}

func (cont *UlasanVillaContImpl) Create(context *gin.Context) {
	// Parse Request Body
	var request villa.UlasanCreateRequest
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

func (cont *UlasanVillaContImpl) Delete(context *gin.Context) {
	// Parse Request Body
	var request villa.UlasanDeleteRequest
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Serv.Delete(request)
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
