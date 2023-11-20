package repository

import (
	"database/sql"
	"errors"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/utils"
)

type UsuarioRepository struct {
	Conn *sql.DB
}

func UsuarioRepositoryConstruct(conn *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{
		Conn: conn,
	}
}

func (r *UsuarioRepository) ListarUsuarios() ([]*entity.Usuario, error) {
	var usuarios []*entity.Usuario

	rows, err := r.Conn.Query(`SELECT user_id,username,senha FROM usuarios`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var usuario entity.Usuario

		err = rows.Scan(&usuario.UserID, &usuario.Username, &usuario.Senha)
		if err != nil {
			return nil, err
		}

		usuarios = append(usuarios, &usuario)
	}

	return usuarios, nil
}

func (r *UsuarioRepository) CriarUsuario(usuario *entity.Usuario) (*entity.Usuario, error) {

	var count int

	errSelectCount := r.Conn.QueryRow(`SELECT COUNT(*) FROM usuarios WHERE username =? AND senha=?`, usuario.Username, usuario.Senha).Scan(&count)

	if errSelectCount != nil {
		panic(errSelectCount.Error())
	}

	if count > 0 {
		return nil, errors.New("Usuário já existente!")
	}

	stmt, err := r.Conn.Prepare(`INSERT INTO usuarios (user_id,username,senha) VALUES (?,?,?)`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(usuario.UserID, usuario.Username, usuario.Senha)

	if err != nil {
		return nil, err
	}

	return usuario, nil
}

func (r *UsuarioRepository) ListarUsuariosPorUsername(username string) ([]*entity.Usuario, error) {
	var usuarios []*entity.Usuario

	query := "SELECT user_id, username, senha FROM usuarios WHERE username LIKE ?"
	rows, err := r.Conn.Query(query, "%"+username+"%")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var usuario entity.Usuario

		err = rows.Scan(&usuario.UserID, &usuario.Username, &usuario.Senha)
		usuarios = append(usuarios, &usuario)
	}

	return usuarios, nil
}

func (r *UsuarioRepository) AtualizarUsuario(usuarioAtualizado dto.UsuarioAtualizadoDto) (interface{}, error) {

	stmt, err := r.Conn.Prepare(`UPDATE usuarios SET username =? , senha=? WHERE username=? AND senha=?`)
	if err != nil {
		return nil, err
	}

	_, e := stmt.Exec(usuarioAtualizado.UsernameNovo, utils.Sha256(usuarioAtualizado.SenhaNova), usuarioAtualizado.UsernameAtual, utils.Sha256(usuarioAtualizado.SenhaAtual))
	if e != nil {
		return nil, err
	}

	return map[string]string{"status": "Usuário atualizado com sucesso!"}, nil
}

func (r *UsuarioRepository) ExcluirUsuario(username,senha string) (interface{},error){
	stmt,err := r.Conn.Prepare(`DELETE FROM usuarios WHERE username=? AND senha=?`)
	if err != nil {
		return nil,err
	}

	_,e := stmt.Exec(username,utils.Sha256(senha))
	if e != nil {
		return nil,err
	}

	return map[string]string{"sucesso":"usuário excluido com sucesso!"},nil
}