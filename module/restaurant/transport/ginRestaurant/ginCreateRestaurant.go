package ginRestaurant

import (
	"RestAPI/common"
	"RestAPI/component/appContext"
	restaurantBussines "RestAPI/module/restaurant/business"
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

		user, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(nil)
		}

		data.OwnerId = user.GetUserId()

		//test trash middleware
		//arr := []int{}
		//log.Println(arr[1])

		store := restaurantStorage.NewSQLStore(db)
		biz := restaurantBussines.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurantBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccesResponse(data.FakeId.String()))
	}
}
