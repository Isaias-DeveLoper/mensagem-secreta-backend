package auth_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type LoginUseCase struct {
	AuthRepository entity.IAuthRepository
}

func LoginConstruct(authRepository entity.IAuthRepository) *LoginUseCase {
	return &LoginUseCase{
		AuthRepository: authRepository,
	}
}

func (u *LoginUseCase) Execute(input dto.UsuarioInputDto) (interface{}, error) {
	usuario := entity.NovoUsuario(input)
	r, err := u.AuthRepository.Login(usuario)
	if err != nil {
		return nil, err
	}
	return r, nil
}
