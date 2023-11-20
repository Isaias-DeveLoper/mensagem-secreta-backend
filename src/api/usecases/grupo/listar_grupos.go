package grupo_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type ListarGrupoUseCase struct {
	GrupoRepository entity.IGrupoRepository
}

func ListarGrupoConstruct(grupoRepository entity.IGrupoRepository) *ListarGrupoUseCase {
	return &ListarGrupoUseCase{
		GrupoRepository: grupoRepository,
	}
}

func (u *ListarGrupoUseCase) Execute() ([]*dto.GrupoOutputDto, error) {
	var grupos []*dto.GrupoOutputDto

	r, err := u.GrupoRepository.ListarGrupos()
	if err != nil {
		return nil, err
	}
	for _, grupo := range r {
		grupos = append(grupos, &dto.GrupoOutputDto{
			GrupoID:     grupo.GrupoID,
			Nome:        grupo.Nome,
			Propietario: grupo.Propietario,
		})
	}
	return grupos, nil
}
