package restaurantLikeStorage

import "gorm.io/gorm"

type sqlModel struct {
	db *gorm.DB
}

func NewSQLModel(db *gorm.DB) *sqlModel {
	return &sqlModel{db: db}
}
