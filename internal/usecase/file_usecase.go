package usecase

import (
	"github.com/fathoor/simkes-api/internal/config"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/helper"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/samber/do"
	"path"
)

type FileUseCase struct {
	Config    *config.Config
	Log       *zerolog.Logger
	Validator *validator.Validate
}

func NewFileUseCase(i *do.Injector) (*FileUseCase, error) {
	return &FileUseCase{
		Config:    do.MustInvoke[*config.Config](i),
		Log:       do.MustInvoke[*zerolog.Logger](i),
		Validator: do.MustInvoke[*validator.Validate](i),
	}, nil
}

func (u *FileUseCase) Upload(request *model.FileRequest) model.FileResponse {
	file, err := validation.ValidateFileRequest(u.Validator, request)
	if err != nil {
		u.Log.Error().Err(err).Msg("Validation error")
		panic(&exception.BadRequestError{
			Message: "Invalid request",
		})
	}

	storage := u.Config.Get("APP_STORAGE")

	return model.FileResponse{
		File: file,
		Path: path.Join(storage, file),
	}
}

func (u *FileUseCase) Get(filetype, filename string) string {
	file, err := helper.GetFile(filetype, filename)
	if err != nil {
		u.Log.Info().Str("filetype", filetype).Str("filename", filename).Msg("File not found")
		panic(exception.NotFoundError{
			Message: "File not found",
		})
	}

	return file
}

func (u *FileUseCase) Delete(filetype, filename string) {
	filepath, err := helper.GetFile(filetype, filename)
	if err != nil {
		u.Log.Info().Str("filetype", filetype).Str("filename", filename).Msg("File not found")
		panic(exception.NotFoundError{
			Message: "File not found",
		})
	}

	if err := helper.RemoveFile(filepath); err != nil {
		u.Log.Error().Err(err).Msg("Failed to delete file")
		panic(exception.InternalServerError{
			Message: "Failed to delete file",
		})
	}
}
