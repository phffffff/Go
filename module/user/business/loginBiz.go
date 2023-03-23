package userBusiness

import (
	"RestAPI/common"
	"RestAPI/component/hasher"
	tokenProvider "RestAPI/component/tokenProvider"
	userModel "RestAPI/module/user/model"
	"context"
)

type LoginStore interface {
	FinDataWithCondition(c context.Context, cond map[string]interface{}, moreKeys ...string) (*userModel.User, error)
}

type loginBiz struct {
	//appCtx        context.Context
	store         LoginStore
	tokenProvider tokenProvider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(
	store LoginStore,
	hasher Hasher,
	tokenProvider tokenProvider.Provider,
	expiry int,
) *loginBiz {
	return &loginBiz{
		store:         store,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(c context.Context, data *userModel.UserLogin) (*tokenProvider.Token, error) {
	user, err := biz.store.FinDataWithCondition(c, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, userModel.ErrorEmailOrPasswordInvalid(err)
	}
	passHash := hasher.NewMd5Hash().Hash(data.Password + user.Salt)
	if passHash != user.Password {
		return nil, userModel.ErrorEmailOrPasswordInvalid(err)
	}
	payload := tokenProvider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}
	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return accessToken, nil
}
