package grupo_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type ListarGruposPorPropietarioUseCase struct {
	GrupoRepository entity.IGrupoRepository
}

func ListarGruposPropietarioConstruct(grupoRepository entity.IGrupoRepository) *ListarGruposPorPropietarioUseCase {
	return &ListarGruposPorPropietarioUseCase{
		GrupoRepository: grupoRepository,
	}
}

func (u *ListarGruposPorPropietarioUseCase) Execute(usuario_id string) ([]*dto.GrupoOutputDto, error) {
	var grupos []*dto.GrupoOutputDto

	r, err := u.GrupoRepository.ListarGruposPorPropietario(usuario_id)
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
