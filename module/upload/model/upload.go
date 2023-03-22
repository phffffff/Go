package uploadModel

import (
	"RestAPI/common"
	"errors"
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
		errors.New(common.MsgFileTooLarge),
		common.MsgFileTooLarge,
		common.ErrFileTooLarge,
	)
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err, common.MsgFileIsNotImage, common.ErrFileIsNotImage)
}

func CanNotServerSave(err error) *common.AppError {
	return common.NewCustomError(err, common.MsgCanNotSaveFile, common.ErrCanNotSaveFile)
}
