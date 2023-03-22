package restaurantBussines

import (
	"RestAPI/common"
	restaurantModel "RestAPI/module/restaurant/model"
	"context"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		c context.Context,
		filter *restaurantModel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantModel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	c context.Context,
	filter *restaurantModel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantModel.Restaurant, error) {

	data, err := biz.store.ListDataWithCondition(c, filter, paging)
	if err != nil {
		return nil, common.ErrCannotCRUDEmpty(restaurantModel.EntityName, common.ListConstant, err)
	}
	return data, nil

}
