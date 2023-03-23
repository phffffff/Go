package userModel

import "RestAPI/common"

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (User) TableName() string {
	return "users"
}

func (user *User) Mask() {
	user.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (UserCreate) TableName() string { return User{}.TableName() }

func (userCreate *UserCreate) Mask(isAdminOrOwned bool) {
	userCreate.GenUID(common.DbTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

func ErrEmailExisted(err error) *common.AppError {
	return common.NewCustomError(err, MsgUserEmailExited, ErrUserEmailExisted)
}

func ErrUserDisabled(err error) *common.AppError {
	return common.NewCustomError(err, MsgUserDisabled, ErrUserIsDisable)
}

func ErrorEmailOrPasswordInvalid(err error) *common.AppError {
	return common.NewCustomError(err, MsgEmailOrPasswordInvalid, ErrEmailOrPasswordInvalid)
}

const (
	ErrUserEmailExisted = "ErrEmailExisted"
	MsgUserEmailExited  = "email existed"

	MsgUserDisabled  = "user disabled"
	ErrUserIsDisable = "ErrUserIsDisable"

	ErrEmailOrPasswordInvalid = "ErrEmailOrPasswordInvalid"
	MsgEmailOrPasswordInvalid = "email or password invalid!"
)
