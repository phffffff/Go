package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

//type AeppError *AppError

func NewErrorResponse(root error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, rootError error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    rootError,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, isOk := e.RootErr.(*AppError); isOk {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		MsgErrDb,
		err.Error(),
		ErrDBKey)
}
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		MsgErrSv,
		err.Error(),
		ErrInternalKey)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, MsgInvalidReq, err.Error(), ErrInvalidRequestKey)
}

func ErrCannotCRUDEmpty(entity string, crud string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot %s %s", crud, strings.ToLower(entity)),
		fmt.Sprintf("ErrCannot%s%s", strings.ToTitle(entity)),
	)
}

func ErrRecordNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", strings.ToTitle(entity)),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted", entity),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}
