package entity

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/utils"
	"github.com/google/uuid"
)

type Grupo struct {
	GrupoID     string
	Nome        string
	Propietario string
	Key         string
	Nonce       string
}

type IGrupoRepository interface {
	ListarGrupos() ([]*Grupo, error)
	ListarGruposPorPropietario(propietario string) ([]*Grupo, error)
	CriarGrupo(grupo *Grupo) (*Grupo, error)
}

func NovoGrupo(grupo dto.GrupoInputDto) *Grupo {
	return &Grupo{
		GrupoID:     uuid.New().String(),
		Nome:        grupo.Nome,
		Propietario: grupo.Propietario,
		Key:         utils.GenerateKey(),
		Nonce:       utils.GenerateNonce(),
	}
}
