package ginUser

import (
	"RestAPI/common"
	"RestAPI/component/appContext"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.MustGet(common.CurrentUser).(common.Requester)

		if !ok {
			panic(nil)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccesResponse(user))
	}
}
