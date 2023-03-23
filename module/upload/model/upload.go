package uploadModel

import (
	"RestAPI/common"
	"errors"
)

const (
	ErrFileUploadTooLarge = "ErrFileTooLarge"
	MsgFileTooLarge       = "file too large"

	ErrFileUploadIsNotImage = "ErrFileIsNotImage"
	MsgFileIsNotImage       = "file is not image"

	ErrCanNotSaveFile = "ErrCanNotSaveFile"
	MsgCanNotSaveFile = "can not save file"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New(MsgFileTooLarge),
		MsgFileTooLarge,
		ErrFileUploadTooLarge,
	)
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err, MsgFileIsNotImage, ErrFileUploadIsNotImage)
}

func CanNotServerSave(err error) *common.AppError {
	return common.NewCustomError(err, MsgCanNotSaveFile, ErrCanNotSaveFile)
}
