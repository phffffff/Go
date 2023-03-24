package ginUser

import (
	"RestAPI/common"
	"RestAPI/component/appContext"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser)

		c.IndentedJSON(http.StatusOK, common.SimpleSuccesResponse(user))
	}
}
