package ginUser

import (
	"RestAPI/component/appContext"
	"RestAPI/component/hasher"
	jwtTokenProvider "RestAPI/component/tokenProvider/jwt"
	userBusiness "RestAPI/module/user/business"
	userModel "RestAPI/module/user/model"
	userStorage "RestAPI/module/user/storage"
	"github.com/gin-gonic/gin"
)

func Login(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData userModel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(err)
		}
		db := appCtx.GetMyDBConnection()
		tokenProvider := jwtTokenProvider.NewJwtProvider("kien")

		store := userStorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userBusiness.NewLoginBiz(store, md5, tokenProvider, 60*60*24*30)

	}
}
