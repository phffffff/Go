package userStorage

import (
	"RestAPI/common"
	userModel "RestAPI/module/user/model"
	"context"
	"gorm.io/gorm"
)

func (sql *sqlStore) FindDataWithCondition(
	c context.Context,
	cond map[string]interface{},
	moreKeys ...string) (*userModel.User, error) {

	db := sql.db.Table(userModel.User{}.TableName())

	//for i := range moreKeys {
	//	db = sql.db.Preload(moreKeys[i])
	//}

	var user userModel.User
	if err := db.Where(cond).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(userModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &user, nil
}
