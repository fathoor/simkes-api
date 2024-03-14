package usecase

import (
	"github.com/fathoor/simkes-api/internal/entity"
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/repository"
	"github.com/fathoor/simkes-api/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/samber/do"
	"time"
)

type CutiUseCase struct {
	CutiRepository *repository.CutiRepository
	Log            *zerolog.Logger
	Validator      *validator.Validate
}

func NewCutiUseCase(i *do.Injector) (*CutiUseCase, error) {
	return &CutiUseCase{
		CutiRepository: do.MustInvoke[*repository.CutiRepository](i),
		Log:            do.MustInvoke[*zerolog.Logger](i),
		Validator:      do.MustInvoke[*validator.Validate](i),
	}, nil
}

func (u *CutiUseCase) Create(request *model.CutiCreateRequest) model.CutiResponse {
	validation.ValidateCutiCreateRequest(u.Validator, u.Log, request)

	tanggalMulai, err := time.Parse("2006-01-02", request.TanggalMulai)
	if err != nil {
		u.Log.Info().Str("tanggal_mulai", request.TanggalMulai).Msg("Invalid date format")
		panic(exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	tanggalSelesai, err := time.Parse("2006-01-02", request.TanggalSelesai)
	if err != nil {
		u.Log.Info().Str("tanggal_selesai", request.TanggalSelesai).Msg("Invalid date format")
		panic(exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	cuti := entity.Cuti{
		ID:             uuid.New(),
		NIP:            request.NIP,
		TanggalMulai:   tanggalMulai,
		TanggalSelesai: tanggalSelesai,
		Keterangan:     request.Keterangan,
	}

	if err := u.CutiRepository.Insert(&cuti); err != nil {
		u.Log.Error().Err(err).Msg("Failed to insert cuti")
		panic(exception.InternalServerError{
			Message: "Failed to insert cuti",
		})
	}

	return model.CutiResponse{
		ID:             cuti.ID.String(),
		NIP:            cuti.NIP,
		TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
		Keterangan:     cuti.Keterangan,
		Status:         cuti.Status,
	}
}

func (u *CutiUseCase) GetAll() []model.CutiResponse {
	cuti, err := u.CutiRepository.FindAll()
	if err != nil {
		u.Log.Error().Err(err).Msg("Failed to get cuti")
		panic(exception.InternalServerError{
			Message: "Failed to get cuti",
		})
	}

	response := make([]model.CutiResponse, len(cuti))
	for i, cuti := range cuti {
		response[i] = model.CutiResponse{
			ID:             cuti.ID.String(),
			NIP:            cuti.NIP,
			TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
			TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
			Keterangan:     cuti.Keterangan,
			Status:         cuti.Status,
		}
	}

	return response
}

func (u *CutiUseCase) GetByNIP(nip string) []model.CutiResponse {
	cuti, err := u.CutiRepository.FindByNIP(nip)
	if err != nil {
		u.Log.Info().Str("nip", nip).Msg("Cuti not found")
		panic(exception.NotFoundError{
			Message: "Cuti not found",
		})
	}

	response := make([]model.CutiResponse, len(cuti))
	for i, cuti := range cuti {
		response[i] = model.CutiResponse{
			ID:             cuti.ID.String(),
			NIP:            cuti.NIP,
			TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
			TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
			Keterangan:     cuti.Keterangan,
			Status:         cuti.Status,
		}
	}

	return response
}

func (u *CutiUseCase) GetByID(id string) model.CutiResponse {
	cutiID, err := uuid.Parse(id)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Invalid UUID")
		panic(exception.BadRequestError{
			Message: "Invalid UUID",
		})
	}

	cuti, err := u.CutiRepository.FindByID(cutiID)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Cuti not found")
		panic(exception.NotFoundError{
			Message: "Cuti not found",
		})
	}

	return model.CutiResponse{
		ID:             cuti.ID.String(),
		NIP:            cuti.NIP,
		TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
		Keterangan:     cuti.Keterangan,
		Status:         cuti.Status,
	}
}

func (u *CutiUseCase) Update(id string, request *model.CutiUpdateRequest) model.CutiResponse {
	validation.ValidateCutiUpdateRequest(u.Validator, u.Log, request)

	cutiID, err := uuid.Parse(id)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Invalid UUID")
		panic(exception.BadRequestError{
			Message: "Invalid UUID",
		})
	}

	cuti, err := u.CutiRepository.FindByID(cutiID)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Cuti not found")
		panic(exception.NotFoundError{
			Message: "Cuti not found",
		})
	}

	tanggalMulai, err := time.Parse("2006-01-02", request.TanggalMulai)
	if err != nil {
		u.Log.Info().Str("tanggal_mulai", request.TanggalMulai).Msg("Invalid date format")
		panic(exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	tanggalSelesai, err := time.Parse("2006-01-02", request.TanggalSelesai)
	if err != nil {
		u.Log.Info().Str("tanggal_selesai", request.TanggalSelesai).Msg("Invalid date format")
		panic(exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	cuti.TanggalMulai = tanggalMulai
	cuti.TanggalSelesai = tanggalSelesai
	cuti.Keterangan = request.Keterangan

	if err := u.CutiRepository.Update(&cuti); err != nil {
		u.Log.Error().Err(err).Msg("Failed to update cuti")
		panic(exception.InternalServerError{
			Message: "Failed to update cuti",
		})
	}

	return model.CutiResponse{
		ID:             cuti.ID.String(),
		NIP:            cuti.NIP,
		TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
		Keterangan:     cuti.Keterangan,
		Status:         cuti.Status,
	}
}

func (u *CutiUseCase) UpdateStatus(id string, request *model.CutiUpdateRequest) model.CutiResponse {
	validation.ValidateCutiUpdateRequest(u.Validator, u.Log, request)

	cutiID, err := uuid.Parse(id)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Invalid UUID")
		panic(exception.BadRequestError{
			Message: "Invalid UUID",
		})
	}

	cuti, err := u.CutiRepository.FindByID(cutiID)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Cuti not found")
		panic(exception.NotFoundError{
			Message: "Cuti not found",
		})
	}

	tanggalMulai, err := time.Parse("2006-01-02", request.TanggalMulai)
	if err != nil {
		u.Log.Info().Str("tanggal_mulai", request.TanggalMulai).Msg("Invalid date format")
		panic(exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	tanggalSelesai, err := time.Parse("2006-01-02", request.TanggalSelesai)
	if err != nil {
		u.Log.Info().Str("tanggal_selesai", request.TanggalSelesai).Msg("Invalid date format")
		panic(exception.BadRequestError{
			Message: "Invalid date format",
		})
	}

	cuti.TanggalMulai = tanggalMulai
	cuti.TanggalSelesai = tanggalSelesai
	cuti.Keterangan = request.Keterangan
	cuti.Status = request.Status

	if err := u.CutiRepository.Update(&cuti); err != nil {
		u.Log.Error().Err(err).Msg("Failed to update cuti")
		panic(exception.InternalServerError{
			Message: "Failed to update cuti",
		})
	}

	return model.CutiResponse{
		ID:             cuti.ID.String(),
		NIP:            cuti.NIP,
		TanggalMulai:   cuti.TanggalMulai.Format("2006-01-02"),
		TanggalSelesai: cuti.TanggalSelesai.Format("2006-01-02"),
		Keterangan:     cuti.Keterangan,
		Status:         cuti.Status,
	}
}

func (u *CutiUseCase) Delete(id string) {
	cutiID, err := uuid.Parse(id)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Invalid UUID")
		panic(exception.BadRequestError{
			Message: "Invalid UUID",
		})
	}

	cuti, err := u.CutiRepository.FindByID(cutiID)
	if err != nil {
		u.Log.Info().Str("id", id).Msg("Cuti not found")
		panic(exception.NotFoundError{
			Message: "Cuti not found",
		})
	}

	if err := u.CutiRepository.Delete(&cuti); err != nil {
		u.Log.Err(err).Msg("Failed to delete cuti")
		panic(exception.InternalServerError{
			Message: "Failed to delete cuti",
		})
	}
}
