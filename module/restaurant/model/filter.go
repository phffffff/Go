package restaurantModel

type Filter struct {
	Status []int `json:"-"`
	CityId int   `json:"city_id" form:"city_id"`
}
