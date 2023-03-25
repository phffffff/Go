package restaurantStorage

import (
	"RestAPI/common"
	restaurantModel "RestAPI/module/restaurant/model"
	"context"
)

func (sql *sqlStore) Create(c context.Context, data *restaurantModel.RestaurantCreate) error {
	db := sql.db.Begin()
	//defer func() {
	//	if err := recover(); err != nil {
	//		db.Rollback()
	//	}
	//}()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
