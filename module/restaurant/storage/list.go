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

	for _, item := range moreKeys {
		db = db.Preload(item)
	}

	if cursor := paging.FakeCursor; cursor != "" {
		uid, err := common.FromBase58(cursor)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		db = db.Where("id < ?", int(uid.GetLocalID()))
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&restaurantList).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(restaurantList) > 0 {
		lastRestaurantList := restaurantList[len(restaurantList)-1]
		lastRestaurantList.Mask(false)
		paging.NextCursor = lastRestaurantList.FakeId.String()
	}

	return restaurantList, nil
}
