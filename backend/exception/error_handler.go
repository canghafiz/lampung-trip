package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lampung_trip/helper"
	"lampung_trip/model/web"
	"net/http"
	"strings"
)

func ErrorHandler(c *gin.Context, err interface{}) {
	if newError(c, err) {
		return
	}
	internalServerError(c, err)
}

func newError(c *gin.Context, err interface{}) bool {
	var errMsg string

	switch e := err.(type) {
	case error:
		errMsg = e.Error()
	case string:
		errMsg = e
	default:
		errMsg = fmt.Sprintf("%v", err)
	}

	lowered := strings.ToLower(errMsg)

	errorResponses := map[string]int{
		"user already exists": http.StatusConflict,
		"user not found":      http.StatusUnauthorized,
		"password wrong":      http.StatusUnauthorized,
	}

	statusCode, exists := errorResponses[lowered]
	if !exists {
		statusCode = http.StatusBadRequest
	}

	webResponse := web.ApiResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   errMsg,
	}

	_ = helper.WriteToResponseBody(c, statusCode, webResponse)
	return true
}

// internalServerError handles all uncaught errors as HTTP 500.
func internalServerError(c *gin.Context, err interface{}) {
	webResponse := web.ApiResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   fmt.Sprintf("%v", err),
	}
	_ = helper.WriteToResponseBody(c, http.StatusInternalServerError, webResponse)
}
