package userBusiness

import (
	"RestAPI/common"
	userModel "RestAPI/module/user/model"
	"context"
)

type RegisterStore interface {
	Create(c context.Context, data *userModel.UserCreate) error
	FindDataWithCondition(c context.Context, cond map[string]interface{}, moreKeys ...string) (*userModel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	store  RegisterStore
	hasher Hasher
}

func NewRegisterBiz(store RegisterStore, hasher Hasher) *registerBiz {
	return &registerBiz{store: store, hasher: hasher}
}

func (biz *registerBiz) RegisterUser(c context.Context, data *userModel.UserCreate) error {
	user, _ := biz.store.FindDataWithCondition(c, map[string]interface{}{"email": data.Email})
	if user != nil {
		if user.Status == 0 {
			return userModel.ErrUserDisabled(nil)
		}

		return userModel.ErrEmailExisted(nil)
	}

	salt := common.GetSalt(50)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" //hard code
	//data.Status = 1

	if err := biz.store.Create(c, data); err != nil {
		return common.ErrCannotCRUDEntity(userModel.EntityName, common.CreateConstant, err)
	}
	return nil
}
