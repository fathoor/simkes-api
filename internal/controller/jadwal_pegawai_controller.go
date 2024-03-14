package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type JadwalPegawaiController struct {
	JadwalPegawaiUseCase *usecase.JadwalPegawaiUseCase
}

func NewJadwalPegawaiController(i *do.Injector) (*JadwalPegawaiController, error) {
	return &JadwalPegawaiController{
		JadwalPegawaiUseCase: do.MustInvoke[*usecase.JadwalPegawaiUseCase](i),
	}, nil
}

func (c *JadwalPegawaiController) Create(ctx *fiber.Ctx) error {
	var request web.JadwalPegawaiRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.JadwalPegawaiUseCase.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
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

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	case tahun != 0 && bulan != 0:
		response := c.JadwalPegawaiUseCase.GetByTahunBulan(int16(tahun), int16(bulan))

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	default:
		response := c.JadwalPegawaiUseCase.GetAll()

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (c *JadwalPegawaiController) GetByPK(ctx *fiber.Ctx) error {
	nip := ctx.Params("nip")

	tahun, err := ctx.ParamsInt("tahun")
	exception.PanicIfError(err)

	bulan, err := ctx.ParamsInt("bulan")
	exception.PanicIfError(err)

	hari, err := ctx.ParamsInt("hari")
	exception.PanicIfError(err)

	response := c.JadwalPegawaiUseCase.GetByPK(nip, int16(tahun), int16(bulan), int16(hari))

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *JadwalPegawaiController) Update(ctx *fiber.Ctx) error {
	var request web.JadwalPegawaiRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	nip := ctx.Params("nip")

	tahun, err := ctx.ParamsInt("tahun")
	exception.PanicIfError(err)

	bulan, err := ctx.ParamsInt("bulan")
	exception.PanicIfError(err)

	hari, err := ctx.ParamsInt("hari")
	exception.PanicIfError(err)

	response := c.JadwalPegawaiUseCase.Update(nip, int16(tahun), int16(bulan), int16(hari), &request)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *JadwalPegawaiController) Delete(ctx *fiber.Ctx) error {
	nip := ctx.Params("nip")

	tahun, err := ctx.ParamsInt("tahun")
	exception.PanicIfError(err)

	bulan, err := ctx.ParamsInt("bulan")
	exception.PanicIfError(err)

	hari, err := ctx.ParamsInt("hari")
	exception.PanicIfError(err)

	c.JadwalPegawaiUseCase.Delete(nip, int16(tahun), int16(bulan), int16(hari))

	return ctx.SendStatus(fiber.StatusNoContent)
}
