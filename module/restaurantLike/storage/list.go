package restaurantLikeStorage

import (
	"RestAPI/common"
	restaurantLikeModel "RestAPI/module/restaurantLike/model"
	"context"
)

func (sql *sqlModel) GetRestaurantLikes(
	c context.Context,
	ids []int,
	filter *restaurantLikeModel.Filter,
) (map[int]int /*map[restaurant_id]like_count*/, error) {

	result := make(map[int]int)

	type likeStruct struct {
		RestaurantId int `json:"-" gorm:"culumn:restaurant_id"`
		LikeCount    int `json:"-" gorm:"column:count"`
	}

	var listLike []likeStruct

	db := sql.db.Table(restaurantLikeModel.Like{}.TableName())
	if f := filter; f != nil {
		if f.UserId >= 0 {
			db = db.Where("user_id in (?)", f.UserId)
		}
		if f.RestaurantId >= 0 {
			db = db.Where("restaurant_id in (?)", f.RestaurantId)
		}
	}
	if err := db.Select("restaurant_id", "count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikeCount
	}

	return result, nil
}
