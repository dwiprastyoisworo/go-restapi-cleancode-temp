package response

import (
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/constants"
	"net/http"
)

type APIError struct {
	Type        constants.ErrorType // Jenis error
	Err         error               // Error original
	MessageCode string              // Kode pesan untuk i18n
	Meta        map[string]any      // Metadata tambahan
}

func NewAPIError(errorType constants.ErrorType, messageCode string, err error, meta map[string]any) *APIError {
	return &APIError{
		Type:        errorType,
		Err:         err,
		MessageCode: messageCode,
		Meta:        meta,
	}
}

func (e *APIError) HTTPStatus() int {
	switch e.Type {
	case constants.ErrorNotFoundType:
		return http.StatusNotFound
	case constants.ErrorInternalServerType:
		return http.StatusInternalServerError
	case constants.ErrorBadRequestType:
		return http.StatusBadRequest
	case constants.ErrorUnauthorizedType:
		return http.StatusUnauthorized
	case constants.ErrorForbiddenType:
		return http.StatusForbidden
	case constants.ErrorConflictType:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
