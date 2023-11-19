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
	ListarUsuariosPorUsername(username string)([]*Usuario,error)
	CriarUsuario(usuario *Usuario) (*Usuario, error)
	AtualizarUsuario(usuarioAtualizado dto.UsuarioAtualizadoDto)(interface{},error)
	ExcluirUsuario(username,senha string) (interface{},error)
}

func NovoUsuario(username, senha string) *Usuario {
	id := uuid.New()

	return &Usuario{
		UserID:   id.String(),
		Username: username,
		Senha:    utils.Sha256(senha),
	}
}
