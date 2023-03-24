package ginUser

import (
	"RestAPI/common"
	"RestAPI/component/appContext"
	"RestAPI/component/hasher"
	"RestAPI/component/tokenProvider/jwt"
	userBusiness "RestAPI/module/user/business"
	userModel "RestAPI/module/user/model"
	userStorage "RestAPI/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData userModel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(err)
		}

		db := appCtx.GetMyDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userStorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userBusiness.NewLoginBiz(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccesResponse(account))

	}
}
