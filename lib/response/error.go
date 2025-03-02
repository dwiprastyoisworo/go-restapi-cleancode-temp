package response

import "github.com/nicksnyder/go-i18n/v2/i18n"

type ResponseFormatter struct {
	i18n        *i18n.Bundle
	logger      Logger // Interface logger
	exposeError bool   // Flag untuk environment development
}

type Logger interface {
	LogError(err error, meta map[string]any)
}

func NewFormatter(i18n *i18n.Bundle, logger Logger, exposeError bool) *ResponseFormatter {
	return &ResponseFormatter{
		i18n:        i18n,
		logger:      logger,
		exposeError: exposeError,
	}
}

type Response struct {
	Code    string       `json:"code,omitempty"`
	Success bool         `json:"success"`
	Message string       `json:"message,omitempty"`
	Data    any          `json:"data,omitempty"`
	Meta    *Meta        `json:"meta,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Message string `json:"message,omitempty"`
	Debug   string `json:"debug,omitempty"` // Hanya untuk environment non-production
}
