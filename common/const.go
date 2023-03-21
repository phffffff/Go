package common

import "log"

const (
	CreateConstant = "CREATE"
	UpdateConstant = "UPDATE"
	DeleteConstant = "DELETE"
	ListConstant   = "LIST"

	MsgErrDb = "something went wrong with DB"
	ErrDBKey = "DB_ERROR"

	MsgErrSv       = "something went wrong with Server"
	ErrInternalKey = "ErrInternal"

	MsgInvalidReq        = "invalid request"
	ErrInvalidRequestKey = "ErrInvalidRequest"

	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Revovery Error:", err)
	}
}
