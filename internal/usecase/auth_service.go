package usecase

import (
	"github.com/fathoor/simkes-api/internal/model"
)

type AuthService interface {
	Login(request *model.AuthRequest) model.AuthResponse
}
