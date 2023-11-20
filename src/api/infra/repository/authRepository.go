package repository

import (
	"database/sql"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/utils"
)

type AuthRepository struct {
	Conn *sql.DB
}

func AuthRepositoryConstruct(conn *sql.DB) *AuthRepository {
	return &AuthRepository{
		Conn: conn,
	}
}

func (r *AuthRepository) Login(usuario *entity.Usuario) (interface{}, error) {
	var count int
	var usuario_id string

	err := r.Conn.QueryRow(`SELECT COUNT(*) FROM usuarios WHERE username=? AND senha=?`, usuario.Username, usuario.Senha).Scan(&count)
	if err != nil {
		return nil, err
	}
	err = r.Conn.QueryRow(`SELECT user_id FROM usuarios WHERE username=? AND senha=?`, usuario.Username,usuario.Senha).Scan(&usuario_id)

	if count == 0 {
		return map[string]interface{}{"isAuthenticated": false, "usuario_id": "", "token": ""}, nil
	}
	return map[string]interface{}{"isAuthenticated": true, "usuario_id": usuario_id, "token": utils.GenerateToken()}, nil
}
