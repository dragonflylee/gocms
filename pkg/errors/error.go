package errors

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	ErrorTypeMessage gin.ErrorType = 1 << 4
)

var (
	errOK = errors.New("OK")

	ErrInternal  = New("InternalError")
	ErrBinding   = New("InvalidParameter")
	ErrForbbiden = New("Forbbiden")
	ErrRequest   = New("BadRequest")
	ErrCode      = New("InvalidCode")
	ErrCapcha    = New("InvalidCapcha")
	ErrConfig    = New("NoConfig")
	ErrPassword  = New("Password")
)

func New(format string, args ...interface{}) *gin.Error {
	return &gin.Error{Err: errors.New(format),
		Type: ErrorTypeMessage, Meta: args}
}

func OK(args ...interface{}) *gin.Error {
	err := &gin.Error{Err: errOK, Type: ErrorTypeMessage}
	if len(args) == 1 {
		err.Meta = args[0]
	} else {
		err.Meta = args
	}
	return err
}
