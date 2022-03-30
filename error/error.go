package error

import (
	"errors"
)

type ServiceError struct {
	Message string `json:"message"`
}

var (
	ErrDatabaseOpen    = errors.New("open_db")
	ErrDatabaseQuery   = errors.New("query_db")
	ErrDatabasePrepare = errors.New("preparing_db")
	ErrDatabaseEntry   = errors.New("entry_not_exist")
)
