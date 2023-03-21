package restaurantStorage

import (
	"RestAPI/common"
	restaurantModel "RestAPI/module/restaurant/model"
	"context"
)

func (sql *sqlStore) Create(c context.Context, data *restaurantModel.RestaurantCreate) error {
	if err := sql.db.Table(data.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
