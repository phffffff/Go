package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	Email     string `json:"email" gorm:"column:email;"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	//Phone     string `json:"phone" gorm:"column:phone;"`
}

func (SimpleUser) TableName() string { return "users" }

func (simpleUser *SimpleUser) Mark(isAdminOrOwner bool) {
	simpleUser.GenUID(DbTypeUser)
}
