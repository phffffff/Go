package restaurantStorage

import (
	"RestAPI/common"
	restaurantModel "RestAPI/module/restaurant/model"
	"context"
)

func (sql *sqlStore) ListDataWithCondition(
	c context.Context,
	filter *restaurantModel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantModel.Restaurant, error) {
	var restaurantList []restaurantModel.Restaurant
	db := sql.db

	if err := db.Error; err != nil {
		//nếu có lỗi thì data là nil và lỗi là db, db trả về error
		return nil, common.ErrDB(err)
	}
	if f := filter; f != nil {
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
		if f.CityId >= 0 {
			db = db.Where("city_id = ?", f.CityId)
		}
	}

	if err := db.
		Table(restaurantModel.Restaurant{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&restaurantList).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return restaurantList, nil
}
