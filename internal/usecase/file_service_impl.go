package usecase

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/helper"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/validation"
	"path"
)

type fileServiceImpl struct {
	config.Config
}

func (service *fileServiceImpl) Upload(request *model.FileRequest) model.FileResponse {
	file, err := validation.ValidateFileRequest(request)
	exception.PanicIfError(err)

	storage := service.Config.Get("APP_STORAGE")

	return model.FileResponse{
		File: file,
		Path: path.Join(storage, file),
	}
}

func (service *fileServiceImpl) Get(filetype, filename string) string {
	file, err := helper.GetFile(filetype, filename)
	exception.PanicIfError(err)

	return file
}

func (service *fileServiceImpl) Delete(filetype, filename string) {
	filepath, err := helper.GetFile(filetype, filename)
	exception.PanicIfError(err)

	if err := helper.RemoveFile(filepath); err != nil {
		panic(exception.InternalServerError{
			Message: "Failed to delete file",
		})
	}
}

func NewFileServiceProvider() FileService {
	return &fileServiceImpl{
		Config: config.ProvideConfig(),
	}
}
