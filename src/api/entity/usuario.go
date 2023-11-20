package entity

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/utils"
	"github.com/google/uuid"
)

type Usuario struct {
	UserID   string
	Username string
	Senha    string
}

type IUsuarioRepository interface {
	ListarUsuarios() ([]*Usuario, error)
	ListarUsuariosPorUsername(username string) ([]*Usuario, error)
	CriarUsuario(usuario *Usuario) (*Usuario, error)
	AtualizarUsuario(usuarioAtualizado dto.UsuarioAtualizadoDto) (interface{}, error)
	ExcluirUsuario(username, senha string) (interface{}, error)
}

func NovoUsuario(usuario dto.UsuarioInputDto) *Usuario {
	return &Usuario{
		UserID:   uuid.New().String(),
		Username: usuario.Username,
		Senha:    utils.Sha256(usuario.Senha),
	}
}
