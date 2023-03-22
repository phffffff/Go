package uploadProvider

import (
	"RestAPI/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(c context.Context, data []byte, dst string) (*common.Image, error)
}
