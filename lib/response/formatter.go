package response

import (
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/configs"
	"github.com/gin-gonic/gin"
)

func (f *ResponseFormatter) FormatSuccess(ctx *gin.Context, success *APISuccess) *Response {
	message := f.translateMessage(ctx, success.MessageCode, nil)

	return &Response{
		Success: true,
		Message: message,
		Data:    success.Data,
		Meta:    success.Meta,
	}
}

func (f *ResponseFormatter) FormatError(ctx *gin.Context, apiError *APIError) *Response {
	// Log error dengan metadata
	f.logger.LogError(apiError.Err, map[string]any{
		"type":         apiError.Type,
		"message_code": apiError.MessageCode,
		"meta":         apiError.Meta,
	})

	// Terjemahkan pesan error
	translatedMsg := f.translateMessage(ctx, apiError.MessageCode, apiError.Meta)

	response := &Response{
		Success: false,
		Error: &ErrorDetail{
			Code:    apiError.MessageCode,
			Message: translatedMsg,
		},
	}

	// Tambahkan debug info jika diperlukan
	if f.exposeError && apiError.Err != nil {
		response.Error.Debug = apiError.Err.Error()
	}

	return response
}

func (f *ResponseFormatter) translateMessage(ctx *gin.Context, code string, params map[string]any) string {
	lang := ctx.GetHeader("Accept-Language")
	return configs.Translate(f.i18n, lang, code, params)
}
