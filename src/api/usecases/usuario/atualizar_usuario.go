package usuario_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type AtualizarUsuarioUseCase struct {
	UsuarioRepository entity.IUsuarioRepository
}

func AtualizarUsuarioConstruct(usuarioRepository entity.IUsuarioRepository) *AtualizarUsuarioUseCase {
	return &AtualizarUsuarioUseCase{
		UsuarioRepository: usuarioRepository,
	}
}

func(u *AtualizarUsuarioUseCase) Execute(usuarioAtualizado dto.UsuarioAtualizadoDto) (interface{}, error) {
	r, err := u.UsuarioRepository.AtualizarUsuario(usuarioAtualizado)
	if err != nil {
		return nil, err
	}
	return r, nil
}
