package helpers

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
func NewErrorResponse(code, messageStatus, messageError string) *Response {
	return &Response{
		Success: false,
		Message: messageStatus,
		Error: &ResponseError{
			Code:    code,
			Message: messageError,
		},
	}
}
