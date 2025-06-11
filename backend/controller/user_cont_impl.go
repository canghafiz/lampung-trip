package controller

import (
	"github.com/gin-gonic/gin"
	"lampung_trip/exception"
	"lampung_trip/helper"
	"lampung_trip/model/service"
	"lampung_trip/model/web"
	"lampung_trip/model/web/user"
	"net/http"
	"strconv"
)

type UserContImpl struct {
	Serv service.UserServ
}

func NewUserContImpl(serv service.UserServ) *UserContImpl {
	return &UserContImpl{Serv: serv}
}

func (cont *UserContImpl) LoginAdminKaryawan(context *gin.Context) {
	// Parse Request Body
	var request user.LoginRequest
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	token, errServ := cont.Serv.LoginAdminKaryawan(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   token,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UserContImpl) Login(context *gin.Context) {
	// Parse Request Body
	var request user.LoginRequest
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	token, errServ := cont.Serv.Login(request)
	if errServ != nil {
		exception.ErrorHandler(context, errServ)
		return
	}

	// Response
	api := web.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   token,
	}

	errResponse := helper.WriteToResponseBody(context, api.Code, api)
	if errResponse != nil {
		exception.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *UserContImpl) UpdatePw(context *gin.Context) {
	// Parse Request Body
	var request user.UpdatePwRequest
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Serv.UpdatePw(request)
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

func (cont *UserContImpl) RegisterAdmin(context *gin.Context) {
	// Parse Request Body
	var request user.RegisterAdminRequest
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Serv.RegisterAdmin(request)
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

func (cont *UserContImpl) RegisterKaryawan(context *gin.Context) {
	// Parse Request Body
	var request user.RegisterKaryawanRequest
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Serv.RegisterKaryawan(request)
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

func (cont *UserContImpl) RegisterCustomer(context *gin.Context) {
	// Parse Request Body
	var request user.RegisterCustomerRequest
	errParse := helper.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exception.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.Serv.RegisterCustomer(request)
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

func (cont *UserContImpl) DeleteKaryawan(context *gin.Context) {
	id := context.Param("id")
	idFinal, _ := strconv.Atoi(id)

	// Call Service
	errServ := cont.Serv.DeleteKaryawan(idFinal)
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

func (cont *UserContImpl) DeleteAdmin(context *gin.Context) {
	id := context.Param("id")
	idFinal, _ := strconv.Atoi(id)

	// Call Service
	errServ := cont.Serv.DeleteAdmin(idFinal)
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
