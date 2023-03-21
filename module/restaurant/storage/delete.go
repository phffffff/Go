package restaurantStorage

import (
	"RestAPI/common"
	restaurantModel "RestAPI/module/restaurant/model"
	"context"
	"gorm.io/gorm"
)

func (sql *sqlStore) Delete(c context.Context, id int) error {
	if err := sql.db.Table(restaurantModel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrRecordNotFound(restaurantModel.EntityName, err)
		}
		return common.ErrDB(err)
	}
	return nil
}
