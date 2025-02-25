package helpers

import (
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/configs"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"log"
	"net/http"
)

type AppError struct {
	Error       error
	MessageCode string
	Code        int
	Meta        map[string]string
}

func NewNotFoundError(messageCode string, err error) *AppError {
	return &AppError{
		Error:       err,
		MessageCode: messageCode,
		Code:        http.StatusNotFound,
	}
}

func NewInternalServerError(messageCode string, err error) *AppError {
	return &AppError{
		Error:       err,
		MessageCode: messageCode,
		Code:        http.StatusInternalServerError,
	}
}

func NewBadRequestError(messageCode string, err error) *AppError {
	return &AppError{
		Error:       err,
		MessageCode: messageCode,
		Code:        http.StatusBadRequest,
	}
}

func NewUnauthorizedError(messageCode string, err error) *AppError {
	return &AppError{
		Error:       err,
		MessageCode: messageCode,
		Code:        http.StatusUnauthorized,
	}
}

func NewForbiddenError(messageCode string, err error) *AppError {
	return &AppError{
		Error:       err,
		MessageCode: messageCode,
		Code:        http.StatusForbidden,
	}
}

func NewConflictError(messageCode string, err error) *AppError {
	return &AppError{
		Error:       err,
		MessageCode: messageCode,
		Code:        http.StatusConflict,
	}
}

// ResponseError digunakan untuk menyimpan detail error jika terjadi kegagalan.
type ResponseError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// Meta digunakan untuk informasi tambahan seperti paging.
type Meta struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	Total int `json:"total,omitempty"`
}

// Response adalah struktur standar untuk response API.
type Response struct {
	Success bool           `json:"success"`
	Message string         `json:"message,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
	Meta    *Meta          `json:"meta,omitempty"`
	Error   *ResponseError `json:"error,omitempty"`
}

// NewSuccessResponse membuat response yang sukses.
func NewSuccessResponse(data interface{}, message string, meta *Meta) *Response {
	return &Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

// NewErrorResponse membuat response error.
func NewErrorResponse(response *AppError, i18n *i18n.Bundle) *Response {
	log.Print(response.Error)
	message := configs.Translate(i18n, "en", response.MessageCode, response.Meta)
	return &Response{
		Success: false,
		Error: &ResponseError{
			Code:    response.MessageCode,
			Message: message,
		},
	}
}
