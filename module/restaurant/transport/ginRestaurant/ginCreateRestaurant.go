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

func CreateRestaurant(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		var data restaurantModel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		//test trash middleware
		//arr := []int{}
		//log.Println(arr[1])

		store := restaurantStorage.NewSQLStore(db)
		biz := restaurantBussines.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurantBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccesResponse(data.Id))
	}
}
