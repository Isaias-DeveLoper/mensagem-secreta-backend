package usuario_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type ExcluirUsuarioUseCase struct {
	UsuarioRepository entity.IUsuarioRepository
}

func ExcluirUsuarioConstruct(usuarioRepository entity.IUsuarioRepository) *ExcluirUsuarioUseCase {
	return &ExcluirUsuarioUseCase{
		UsuarioRepository: usuarioRepository,
	}
}

func (u *ExcluirUsuarioUseCase) Execute(usuario dto.UsuarioInputDto) (interface{}, error) {
	r, err := u.UsuarioRepository.ExcluirUsuario(usuario.Username, usuario.Senha)
	if err != nil {
		return nil, err
	}
	return r, nil
}
