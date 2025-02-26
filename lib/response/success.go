package response

import "net/http"

type APISuccess struct {
	MessageCode string
	Data        any
	Meta        *Meta
	HTTPCode    int
}

type Meta struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	Total int `json:"total,omitempty"`
}

func NewSuccess(messageCode string, data any, meta *Meta, httpCode int) *APISuccess {
	if httpCode == 0 {
		httpCode = http.StatusOK
	}

	return &APISuccess{
		MessageCode: messageCode,
		Data:        data,
		Meta:        meta,
		HTTPCode:    httpCode,
	}
}
