package restaurantRepo

import (
	"RestAPI/common"
	restaurantModel "RestAPI/module/restaurant/model"
	restaurantLikeModel "RestAPI/module/restaurantLike/model"
	"context"
	"log"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		c context.Context,
		filter *restaurantModel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantModel.Restaurant, error)
}

type LikeRestaurantStore interface {
	GetRestaurantLikes(
		c context.Context,
		ids []int,
		filter *restaurantLikeModel.Filter,
	) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore LikeRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store, likeStore: likeStore}
}

func (repo *listRestaurantRepo) ListRestaurant(
	c context.Context,
	filter *restaurantModel.Filter,
	paging *common.Paging) ([]restaurantModel.Restaurant, error) {

	data, err := repo.store.ListDataWithCondition(c, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotCRUDEntity(restaurantModel.EntityName, common.ListConstant, err)
	}

	ids := make([]int, len(data))

	for i := range ids {
		ids[i] = data[i].Id
	}

	likeMap, err := repo.likeStore.GetRestaurantLikes(c, ids, nil)

	if err != nil {
		log.Println(err)
		return data, nil
	}

	for i, item := range data {
		data[i].LikeCount = likeMap[item.Id]
	}

	return data, nil

}
