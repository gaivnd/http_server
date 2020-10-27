package types

import (
	"errors"
)

var (
	ErrProcessNotFound        = errors.New("process not found")
	ErrGroupNotFound          = errors.New("group not found")
	ErrQueryNotFound	  = errors.New("no query results")
	ErrNameAlreadyExists      = errors.New("process name already exists")
	ErrIdAlreadyExists        = errors.New("process id already exists")
	ErrNameAndUrlDoesNotMatch = errors.New("name and url does not match")
	ErrNotReady		= errors.New("some processes are not ready")
	ErrInternal		= errors.New("internal error")
	ErrAvailSvcNotFound	= errors.New("no availabe service node found")
)
