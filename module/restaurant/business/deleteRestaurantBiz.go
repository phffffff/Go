package restaurantBussines

import (
	"RestAPI/common"
	restaurantModel "RestAPI/module/restaurant/model"
	"context"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		c context.Context,
		cond map[string]interface{},
		moreKeys ...string) (*restaurantModel.Restaurant, error)
	Delete(c context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(c context.Context, id int) error {
	olddata, err := biz.store.FindDataWithCondition(c, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrRecordNotFound(restaurantModel.EntityName, err)
	}
	if olddata.Status == 0 {
		return common.ErrEntityDeleted(restaurantModel.EntityName, nil)
	}

	if err := biz.store.Delete(c, id); err != nil {
		return common.ErrCannotCRUDEntity(restaurantModel.EntityName, common.DeleteConstant, err)
	}
	return nil
}
