package usuario_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type CriarUsuarioUseCase struct {
	UsuarioRepository entity.IUsuarioRepository
}

func CriarUsuarioConstruct(usuarioRepository entity.IUsuarioRepository) *CriarUsuarioUseCase {
	return &CriarUsuarioUseCase{
		UsuarioRepository: usuarioRepository,
	}
}

func (u *CriarUsuarioUseCase) Execute(input dto.UsuarioInputDto) (*dto.UsuarioOutputDto, error) {
	novoUsuario := entity.NovoUsuario(input.Username, input.Senha)
	usuario, err := u.UsuarioRepository.CriarUsuario(novoUsuario)
	if err != nil {
		panic(err.Error())
	}
	return &dto.UsuarioOutputDto{
		UserID:   usuario.UserID,
		Username: usuario.Username,
		Senha:    usuario.Senha,
	}, nil
}
