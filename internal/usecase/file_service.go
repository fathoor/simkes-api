package usecase

import (
	"github.com/fathoor/simkes-api/internal/model"
)

type FileService interface {
	Upload(request *model.FileRequest) model.FileResponse
	Get(filetype, filename string) string
	Delete(filetype, filename string)
}
