package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/samber/do"
)

type FileController struct {
	FileUseCase *usecase.FileUseCase
	Log         *zerolog.Logger
}

func NewFileController(i *do.Injector) (*FileController, error) {
	return &FileController{
		FileUseCase: do.MustInvoke[*usecase.FileUseCase](i),
		Log:         do.MustInvoke[*zerolog.Logger](i),
	}, nil
}

func (c *FileController) Upload(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		c.Log.Error().Err(err).Msg("No file uploaded")
		panic(exception.BadRequestError{
			Message: "No file uploaded",
		})
	}

	fileType := ctx.FormValue("type")

	request := model.FileRequest{
		File: file,
		Type: fileType,
	}

	response := c.FileUseCase.Upload(&request)

	if err := ctx.SaveFile(file, response.Path); err != nil {
		c.Log.Error().Err(err).Msg("Failed to save file")
		panic(exception.InternalServerError{
			Message: "Failed to save file",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *FileController) Download(ctx *fiber.Ctx) error {
	fileType := ctx.Params("filetype")
	fileName := ctx.Params("filename")

	filePath := c.FileUseCase.Get(fileType, fileName)

	return ctx.Download(filePath)
}

func (c *FileController) View(ctx *fiber.Ctx) error {
	fileType := ctx.Params("filetype")
	fileName := ctx.Params("filename")

	filePath := c.FileUseCase.Get(fileType, fileName)

	return ctx.SendFile(filePath)
}

func (c *FileController) Delete(ctx *fiber.Ctx) error {
	fileType := ctx.Params("filetype")
	fileName := ctx.Params("filename")

	c.FileUseCase.Delete(fileType, fileName)

	return ctx.SendStatus(fiber.StatusNoContent)
}
