package restaurantModel

import (
	"RestAPI/common"
	"errors"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"addr" gorm:"column:addr;"`
	Status          int    `json:"status" gorm:"column:status;"`
	CityId          int    `json:"city_id" gorm:"column:city_id;"`
}

func (Restaurant) TableName() string {
	return "restaurant"
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
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"addr" gorm:"column:addr;"`
	CityId          int    `json:"city_id" gorm:"column:city_id;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
