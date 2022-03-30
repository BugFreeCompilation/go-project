package error

import (
	"errors"
)

type ServiceError struct {
	Message string `json:"message"`
}

var (
	ErrDatabaseOpen    = errors.New("error_open_db")
	ErrDatabaseQuery   = errors.New("error_query_db")
	ErrDatabasePrepare = errors.New("error_preparing_db")
	ErrDatabaseEntry   = errors.New("entry_not_exist")
)
