package ginUser

import (
	"RestAPI/common"
	"RestAPI/component/appContext"
	hasher "RestAPI/component/hasher"
	userBusiness "RestAPI/module/user/business"
	userModel "RestAPI/module/user/model"
	userStorage "RestAPI/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Regiter(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()

		var data userModel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		hasher := hasher.NewMd5Hash()
		store := userStorage.NewSqlStore(db)
		biz := userBusiness.NewRegisterBiz(store, hasher)

		if err := biz.RegisterUser(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.IndentedJSON(http.StatusOK, common.SimpleSuccesResponse(
			data.FakeId.String(),
			//data,
		))
	}
}
