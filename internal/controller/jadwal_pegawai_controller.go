package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/samber/do"
)

type JadwalPegawaiController struct {
	JadwalPegawaiUseCase *usecase.JadwalPegawaiUseCase
	Log                  *zerolog.Logger
}

func NewJadwalPegawaiController(i *do.Injector) (*JadwalPegawaiController, error) {
	return &JadwalPegawaiController{
		JadwalPegawaiUseCase: do.MustInvoke[*usecase.JadwalPegawaiUseCase](i),
		Log:                  do.MustInvoke[*zerolog.Logger](i),
	}, nil
}

func (c *JadwalPegawaiController) Create(ctx *fiber.Ctx) error {
	var request model.JadwalPegawaiRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		c.Log.Error().Err(parse).Msg("Invalid request body")
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.JadwalPegawaiUseCase.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *JadwalPegawaiController) Get(ctx *fiber.Ctx) error {
	nip := ctx.Query("nip")
	tahun := ctx.QueryInt("tahun")
	bulan := ctx.QueryInt("bulan")

	switch {
	case nip != "":
		response := c.JadwalPegawaiUseCase.GetByNIP(nip)

		return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	case tahun != 0 && bulan != 0:
		response := c.JadwalPegawaiUseCase.GetByTahunBulan(int16(tahun), int16(bulan))

		return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	default:
		response := c.JadwalPegawaiUseCase.GetAll()

		return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (c *JadwalPegawaiController) GetByPK(ctx *fiber.Ctx) error {
	nip := ctx.Params("nip")

	tahun, err := ctx.ParamsInt("tahun")
	if err != nil {
		c.Log.Info().Str("tahun", ctx.Params("tahun")).Msg("Invalid tahun")
		panic(exception.BadRequestError{
			Message: "Invalid tahun",
		})
	}

	bulan, err := ctx.ParamsInt("bulan")
	if err != nil {
		c.Log.Info().Str("bulan", ctx.Params("bulan")).Msg("Invalid bulan")
		panic(exception.BadRequestError{
			Message: "Invalid bulan",
		})
	}

	hari, err := ctx.ParamsInt("hari")
	if err != nil {
		c.Log.Info().Str("hari", ctx.Params("hari")).Msg("Invalid hari")
		panic(exception.BadRequestError{
			Message: "Invalid hari",
		})
	}

	response := c.JadwalPegawaiUseCase.GetByPK(nip, int16(tahun), int16(bulan), int16(hari))

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *JadwalPegawaiController) Update(ctx *fiber.Ctx) error {
	var request model.JadwalPegawaiRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		c.Log.Error().Err(parse).Msg("Invalid request body")
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	nip := ctx.Params("nip")

	tahun, err := ctx.ParamsInt("tahun")
	if err != nil {
		c.Log.Info().Str("tahun", ctx.Params("tahun")).Msg("Invalid tahun")
		panic(exception.BadRequestError{
			Message: "Invalid tahun",
		})
	}

	bulan, err := ctx.ParamsInt("bulan")
	if err != nil {
		c.Log.Info().Str("bulan", ctx.Params("bulan")).Msg("Invalid bulan")
		panic(exception.BadRequestError{
			Message: "Invalid bulan",
		})
	}

	hari, err := ctx.ParamsInt("hari")
	if err != nil {
		c.Log.Info().Str("hari", ctx.Params("hari")).Msg("Invalid hari")
		panic(exception.BadRequestError{
			Message: "Invalid hari",
		})
	}

	response := c.JadwalPegawaiUseCase.Update(nip, int16(tahun), int16(bulan), int16(hari), &request)

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *JadwalPegawaiController) Delete(ctx *fiber.Ctx) error {
	nip := ctx.Params("nip")

	tahun, err := ctx.ParamsInt("tahun")
	if err != nil {
		c.Log.Info().Str("tahun", ctx.Params("tahun")).Msg("Invalid tahun")
		panic(exception.BadRequestError{
			Message: "Invalid tahun",
		})
	}

	bulan, err := ctx.ParamsInt("bulan")
	if err != nil {
		c.Log.Info().Str("bulan", ctx.Params("bulan")).Msg("Invalid bulan")
		panic(exception.BadRequestError{
			Message: "Invalid bulan",
		})
	}

	hari, err := ctx.ParamsInt("hari")
	if err != nil {
		c.Log.Info().Str("hari", ctx.Params("hari")).Msg("Invalid hari")
		panic(exception.BadRequestError{
			Message: "Invalid hari",
		})
	}

	c.JadwalPegawaiUseCase.Delete(nip, int16(tahun), int16(bulan), int16(hari))

	return ctx.SendStatus(fiber.StatusNoContent)
}
