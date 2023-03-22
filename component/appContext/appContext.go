package appContext

import (
	"RestAPI/component/uploadProvider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMyDBConnection() *gorm.DB
	UploadProvider() uploadProvider.UploadProvider
}

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadProvider.UploadProvider
}

func NewAppCtx(db *gorm.DB, uploadProvider uploadProvider.UploadProvider) *appContext {
	return &appContext{db: db, uploadProvider: uploadProvider}
}

func (appCtx *appContext) GetMyDBConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) UploadProvider() uploadProvider.UploadProvider {
	return appCtx.uploadProvider
}
