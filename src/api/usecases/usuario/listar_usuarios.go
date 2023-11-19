package usuario_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type ListarUsuariosUseCase struct {
	UsuarioRepository entity.IUsuarioRepository
}

func ListarUsuarioConstruct(usuarioRepository entity.IUsuarioRepository) *ListarUsuariosUseCase {
	return &ListarUsuariosUseCase{
		UsuarioRepository: usuarioRepository,
	}
}

func (u *ListarUsuariosUseCase) Execute() ([]*dto.UsuarioOutputDto, error) {
	var usuarios []*dto.UsuarioOutputDto

	r, err := u.UsuarioRepository.ListarUsuarios()

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
