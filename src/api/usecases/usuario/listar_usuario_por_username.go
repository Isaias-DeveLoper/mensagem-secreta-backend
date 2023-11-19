package usuario_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type ListarUsuarioPorUsernameUseCase struct {
	UsuarioRepository entity.IUsuarioRepository
}

func ListarPorUsernameConstruct(usuarioRepository entity.IUsuarioRepository) *ListarUsuarioPorUsernameUseCase {
	return &ListarUsuarioPorUsernameUseCase{
		UsuarioRepository: usuarioRepository,
	}
}

func (u *ListarUsuarioPorUsernameUseCase) Execute(username string) ([]*dto.UsuarioOutputDto, error) {
	var usuarios []*dto.UsuarioOutputDto

	r, err := u.UsuarioRepository.ListarUsuariosPorUsername(username)
	if err != nil {
		return nil, err
	}

	for _, usuario := range r {
		usuarios = append(usuarios, &dto.UsuarioOutputDto{
			UserID:   usuario.UserID,
			Username: usuario.Username,
			Senha:    usuario.Senha,
		})
	}

	return usuarios, nil
}
