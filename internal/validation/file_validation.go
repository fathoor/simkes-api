package validation

import (
	"fmt"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/helper"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/go-playground/validator/v10"
	"path"
)

func ValidateFileRequest(validator *validator.Validate, request *model.FileRequest) (string, error) {
	if valid := validator.Struct(request); valid != nil {
		return "", &exception.BadRequestError{
			Message: "Invalid request data",
		}
	}

	switch request.Type {
	case "image":
		ext := path.Ext(request.File.Filename)
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			return "", &exception.BadRequestError{
				Message: fmt.Sprintf("Invalid file extension: %s is not supported", ext),
			}
		}

		if request.File.Size > 2*1024*1024 {
			return "", &exception.BadRequestError{
				Message: "File size exceeds limit",
			}
		}
	case "doc":
		ext := path.Ext(request.File.Filename)
		if ext != ".pdf" && ext != ".doc" && ext != ".docx" {
			return "", &exception.BadRequestError{
				Message: fmt.Sprintf("Invalid file extension: %s is not supported", ext),
			}
		}

		if request.File.Size > 5*1024*1024 {
			return "", &exception.BadRequestError{
				Message: "File size exceeds limit",
			}
		}
	default:
		return "", &exception.BadRequestError{
			Message: fmt.Sprintf("Invalid file type: %s is not supported", request.Type),
		}
	}

	fileType := request.Type
	fileExt := path.Ext(request.File.Filename)
	filePath := helper.GenerateFile(fileType, fileExt)

	return filePath, nil
}
