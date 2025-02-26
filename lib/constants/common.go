package constants

type ErrorType int

const (
	ErrorNotFoundType ErrorType = iota + 1
	ErrorInternalServerType
	ErrorBadRequestType
	ErrorUnauthorizedType
	ErrorForbiddenType
	ErrorConflictType
)
