package restaurantLikeStorage

import (
	"RestAPI/common"
	restaurantLikeModel "RestAPI/module/restaurantLike/model"
	"context"
)

func (sql *sqlModel) Delete(c context.Context, userId, restaurantId int) error {
	if err := sql.db.
		Table(restaurantLikeModel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
