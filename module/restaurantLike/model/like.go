package restaurantLikeModel

import (
	"RestAPI/common"
	"time"
)

const (
	ErrCannotLikeRestaurant    = "ErrCannotLikeRestaurant"
	MsgErrCannotLikeRestaurant = "cannot like this restaurant"

	ErrCannotUnLikeRestaurant    = "ErrCannotUnLikeRestaurant"
	MsgErrCannotUnLikeRestaurant = "cannot unlike this restaurant"
)

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreateAt     *time.Time         `json:"create_at" gorm:"column:create_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false; foreignKey:UserId"`
}

func (Like) TableName() string { return "restaurant_likes" }
func ErrorCannotLikeRestautant(err error) *common.AppError {
	return common.NewCustomError(err, MsgErrCannotLikeRestaurant, ErrCannotLikeRestaurant)
}
func ErrorCannotUnLikeRestautant(err error) *common.AppError {
	return common.NewCustomError(err, MsgErrCannotUnLikeRestaurant, ErrCannotUnLikeRestaurant)
}
