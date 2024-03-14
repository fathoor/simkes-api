package usecase

import (
	"github.com/fathoor/simkes-api/internal/model"
)

type ShiftService interface {
	Create(request *model.ShiftRequest) model.ShiftResponse
	GetAll() []model.ShiftResponse
	GetByNama(nama string) model.ShiftResponse
	Update(nama string, request *model.ShiftRequest) model.ShiftResponse
	Delete(nama string)
}
