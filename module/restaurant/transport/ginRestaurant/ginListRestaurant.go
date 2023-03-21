package ginRestaurant

import (
	"RestAPI/common"
	"RestAPI/component/appContext"
	restaurantBussines "RestAPI/module/restaurant/Business"
	restaurantModel "RestAPI/module/restaurant/model"
	restaurantStorage "RestAPI/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(ctx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMyDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.FulFill()

		var filter restaurantModel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1}
		filter.CityId = 1

		store := restaurantStorage.NewSQLStore(db)
		biz := restaurantBussines.NewListRestaurantBiz(store)

		data, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.IndentedJSON(http.StatusOK, common.NewSuccesResponse(data, pagingData, filter))

	}
}
