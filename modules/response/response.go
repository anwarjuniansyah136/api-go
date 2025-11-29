package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type BaseResponse struct {
	Meta  Meta `json:"meta"`
	Data  any  `json:"data,omitempty"`
	Error any  `json:"error,omitempty"`
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, BaseResponse{
		Meta: Meta{
			Status: "SUCCESS",
			Message: "OK",
		},
		Data: data,
	})
}

func Error(ctx *gin.Context, statusCode int, message string, err any) {
	ctx.JSON(statusCode, BaseResponse{
		Meta : Meta{
			Status : "ERROR",
			Message : message,
		},
		Error: err,
	})
}