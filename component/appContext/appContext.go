package appContext

import (
	"RestAPI/component/uploadProvider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMyDBConnection() *gorm.DB
	UploadProvider() uploadProvider.UploadProvider
	SecretKey() string
}

type appContext struct {
	db             *gorm.DB
	uploadProvider uploadProvider.UploadProvider
	secretKey      string
}

func NewAppCtx(db *gorm.DB, uploadProvider uploadProvider.UploadProvider, secretKey string) *appContext {
	return &appContext{db: db, uploadProvider: uploadProvider, secretKey: secretKey}
}

func (appCtx *appContext) GetMyDBConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) UploadProvider() uploadProvider.UploadProvider {
	return appCtx.uploadProvider
}
func (appCtx *appContext) SecretKey() string {
	return appCtx.secretKey
}
