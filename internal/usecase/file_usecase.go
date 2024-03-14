package usecase

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/helper"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/samber/do"
	"path"
)

type FileUseCase struct {
	config.Config
}

func NewFileUseCase(i *do.Injector) (*FileUseCase, error) {
	return &FileUseCase{
		do.MustInvoke[config.Config](i),
	}, nil
}

func (u *FileUseCase) Upload(request *model.FileRequest) model.FileResponse {
	file, err := validation.ValidateFileRequest(request)
	exception.PanicIfError(err)

	storage := u.Config.Get("APP_STORAGE")

	return model.FileResponse{
		File: file,
		Path: path.Join(storage, file),
	}
}

func (u *FileUseCase) Get(filetype, filename string) string {
	file, err := helper.GetFile(filetype, filename)
	exception.PanicIfError(err)

	return file
}

func (u *FileUseCase) Delete(filetype, filename string) {
	filepath, err := helper.GetFile(filetype, filename)
	exception.PanicIfError(err)

	if err := helper.RemoveFile(filepath); err != nil {
		panic(exception.InternalServerError{
			Message: "Failed to delete file",
		})
	}
}
