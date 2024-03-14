package usecase

import (
	"github.com/fathoor/simkes-api/internal/model"
)

type DepartemenService interface {
	Create(request *model.DepartemenRequest) model.DepartemenResponse
	GetAll() []model.DepartemenResponse
	GetByDepartemen(d string) model.DepartemenResponse
	Update(d string, request *model.DepartemenRequest) model.DepartemenResponse
	Delete(d string)
}
