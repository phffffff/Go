package userStorage

import (
	"RestAPI/common"
	userModel "RestAPI/module/user/model"
	"context"
)

func (sql *sqlStore) Create(c context.Context, data *userModel.UserCreate) error {
	if err := sql.db.Table(data.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
