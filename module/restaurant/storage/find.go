package restaurantStorage

import (
	"RestAPI/common"
	restaurantModel "RestAPI/module/restaurant/model"
	"context"
	"gorm.io/gorm"
)

func (sql *sqlStore) FindDataWithCondition(
	c context.Context,
	cond map[string]interface{},
	moreKeys ...string) (*restaurantModel.Restaurant, error) {

	var data restaurantModel.Restaurant

	if err := sql.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(restaurantModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
