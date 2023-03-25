package restaurantBussines

import (
	"RestAPI/common"
	restaurantModel "RestAPI/module/restaurant/model"
	"context"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		c context.Context,
		filter *restaurantModel.Filter,
		paging *common.Paging) ([]restaurantModel.Restaurant, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) ListRestaurant(
	c context.Context,
	filter *restaurantModel.Filter,
	paging *common.Paging) ([]restaurantModel.Restaurant, error) {
	data, err := biz.repo.ListRestaurant(c, filter, paging)

	if err != nil {
		return nil, common.ErrCannotCRUDEntity(restaurantModel.EntityName, common.ListConstant, err)
	}

	return data, nil

}
