package restaurantModel

import (
	"RestAPI/common"
	"errors"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"addr" gorm:"column:addr;"`
	Status          int                `json:"status" gorm:"column:status;"`
	CityId          int                `json:"-" gorm:"column:city_id;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"'`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	OwnerId         int                `json:"owner_id" gorm:"column:owner_id;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;foreignKey:OwnerId;"`
	LikeCount       int                `json:"like_count" gorm:"-"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	user := r.User
	if user != nil {
		user.Mark(false)
	}
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

func (data *RestaurantCreate) ErrCannotName(err error) error {
	return common.NewCustomError(
		err,
		ErrNameIsEmpty.Error(),
		"ErrCannotName",
	)
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	CityId          int            `json:"city_id" gorm:"column:city_id;"`
	OwnerId         int            `json:"-" gorm:"column:owner_id"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"'`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
	//OwnerId int     `json:"-" gorm:"column:owner_id"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
