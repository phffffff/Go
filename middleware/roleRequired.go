package middleware

import (
	"RestAPI/common"
	"RestAPI/component/appContext"
	"github.com/gin-gonic/gin"
)

func RoleRequired(ctx appContext.AppContext, allowRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false
		for _, item := range allowRoles {
			if item == user.GetRole() {
				hasFound = true
				break
			}
		}

		if !hasFound {
			panic(common.ErrorNoPermission(nil))
		}
		c.Next()
	}
}
