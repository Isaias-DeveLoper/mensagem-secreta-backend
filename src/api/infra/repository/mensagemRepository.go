package repository

import (
	"database/sql"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/utils"
)

type MensagemRepository struct {
	Conn *sql.DB
}

func MensagemRepositoryConstruct(conn *sql.DB) *MensagemRepository {
	return &MensagemRepository{
		Conn: conn,
	}
}

func (r *MensagemRepository) EnviarMensagem(mensagem *entity.Mensagem) (*entity.Mensagem, error) {
	var credencias dto.KeyNonce

	err := r.Conn.QueryRow(`SELECT keyw,nonce FROM grupos WHERE grupo_id =?`, mensagem.GroupID).Scan(&credencias.Key, &credencias.Nonce)
	if err != nil {
		return nil, err
	}

	stmt, e := r.Conn.Prepare(`INSERT INTO mensagens (grupo_id,texto) VALUES (?,?)`)
	if e != nil {
		return nil, err
	}

	_, er := stmt.Exec(mensagem.GroupID, utils.Encrypt(credencias.Key, credencias.Nonce, mensagem.Texto))
	if er != nil {
		return nil, er
	}

	return mensagem, nil
}

func (r *MensagemRepository) LerMensagens(grupo,senha string) ([]*entity.Mensagem, error) {

	var mensagens []*entity.Mensagem
	var senhaPropietario string
	var informacoesGrupo entity.Grupo

	rows, err := r.Conn.Query(`SELECT grupo_id,texto FROM mensagens WHERE grupo_id=?`, grupo)
	if err != nil {
		return nil, err
	}

	r.Conn.QueryRow(`SELECT grupo_id,nome,propietario,
	keyw,nonce FROM grupos WHERE grupo_id =?`, grupo).Scan(&informacoesGrupo.GrupoID,
		&informacoesGrupo.Nome, &informacoesGrupo.Propietario,
		&informacoesGrupo.Key, &informacoesGrupo.Nonce)

	r.Conn.QueryRow(`SELECT senha FROM usuarios WHERE username=?`, informacoesGrupo.Propietario).Scan(&senhaPropietario)

	if senhaPropietario != utils.Sha256(senha) {
		for rows.Next() {
			var mensagem entity.Mensagem

			err := rows.Scan(&mensagem.GroupID,&mensagem.Texto)
			if err != nil {
				return nil,err
			}
			mensagens = append(mensagens, &entity.Mensagem{
				GroupID: mensagem.GroupID,
				Texto:   mensagem.Texto,
			})
		}
	} else {
		for rows.Next() {
			var mensagem entity.Mensagem
			err := rows.Scan(&mensagem.GroupID,&mensagem.Texto)
			if err != nil {
				return nil,err
			}
			mensagens = append(mensagens, &entity.Mensagem{
				GroupID: mensagem.GroupID,
				Texto:   utils.Decrypt(informacoesGrupo.Key, informacoesGrupo.Nonce, mensagem.Texto),
			})
		}
	}

	return mensagens, nil
}
