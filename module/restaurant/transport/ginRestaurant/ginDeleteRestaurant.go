package ginRestaurant

import (
	"RestAPI/common"
	"RestAPI/component/appContext"
	restaurantBussines "RestAPI/module/restaurant/Business"
	restaurantStorage "RestAPI/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantStorage.NewSQLStore(db)
		biz := restaurantBussines.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccesResponse(1))
	}
}
