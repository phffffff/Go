package middleware

import (
	"RestAPI/common"
	"RestAPI/component/appContext"
	"RestAPI/component/tokenProvider/jwt"
	userModel "RestAPI/module/user/model"
	userStorage "RestAPI/module/user/storage"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	ErrWrongAuthHeader = "ErrWrongAuthHeader"
	MsgWrongAuthHeader = "wrong authen header"
)

func ErrorWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(err, MsgWrongAuthHeader, ErrWrongAuthHeader)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) > 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrorWrongAuthHeader(nil)
	}
	return parts[1], nil
}

func RequiredAuth(appCtx appContext.AppContext) gin.HandlerFunc {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMyDBConnection()
		store := userStorage.NewSqlStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		user, err := store.FindDataWithCondition(
			c.Request.Context(),
			map[string]interface{}{"id": payload.UserId},
		)

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(userModel.ErrUserDisabled(nil))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
