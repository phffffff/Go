package restaurantBussines

import (
	"RestAPI/common"
	restaurantModel "RestAPI/module/restaurant/model"
	"context"
)

type CreateRestaurantStore interface {
	Create(c context.Context, data *restaurantModel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurantBiz(c context.Context, data *restaurantModel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return data.ErrCannotName(err)
	}

	if err := biz.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(restaurantModel.EntityName, common.CreateConstant, err)
	}
	return nil
}
