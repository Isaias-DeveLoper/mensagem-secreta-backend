package repository

import (
	"database/sql"
	"errors"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type GrupoRepository struct {
	Conn *sql.DB
}

func GrupoRepositoryConstruct(conn *sql.DB) *GrupoRepository {
	return &GrupoRepository{
		Conn: conn,
	}
}

func (r *GrupoRepository) CriarGrupo(grupo *entity.Grupo) (*entity.Grupo, error) {

	var count int

	errSelectCount := r.Conn.QueryRow(`SELECT COUNT(*) FROM grupos WHERE nome =? `, grupo.Nome).Scan(&count)

	if errSelectCount != nil {
		panic(errSelectCount.Error())
	}

	if count > 0 {
		return nil, errors.New("Grupo já existente!")
	}

	stmt, err := r.Conn.Prepare(`INSERT INTO grupos (grupo_id,nome,propietario,keyw,nonce) VALUES (?,?,?,?,?)`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(grupo.GrupoID, grupo.Nome, grupo.Propietario, grupo.Key, grupo.Nonce)

	if err != nil {
		return nil, err
	}

	return grupo, nil
}

func (r *GrupoRepository) ListarGrupos() ([]*entity.Grupo, error) {
	var grupos []*entity.Grupo

	rows, err := r.Conn.Query(`SELECT grupo_id,nome,propietario FROM grupos`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var grupo entity.Grupo

		err = rows.Scan(&grupo.GrupoID, &grupo.Nome, &grupo.Propietario)
		grupos = append(grupos, &grupo)
	}
	return grupos, nil
}

func (r *GrupoRepository) ListarGruposPorPropietario(propietario string) ([]*entity.Grupo, error) {
	var grupos []*entity.Grupo

	rows, err := r.Conn.Query(`SELECT grupo_id,nome,propietario FROM grupos WHERE propietario =?`,propietario)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var grupo entity.Grupo

		err = rows.Scan(&grupo.GrupoID, &grupo.Nome, &grupo.Propietario)
		grupos = append(grupos, &grupo)
	}
	return grupos, nil
}
