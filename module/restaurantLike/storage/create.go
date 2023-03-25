package restaurantLikeStorage

import (
	"RestAPI/common"
	restaurantLikeModel "RestAPI/module/restaurantLike/model"
	"context"
)

func (sql *sqlModel) Create(c context.Context, data *restaurantLikeModel.Like) error {
	if err := sql.db.Table(data.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
