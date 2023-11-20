package grupo_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type CriarGrupoUseCase struct {
	GrupoRepository entity.IGrupoRepository
}

func CriarGrupoConstruct(grupoRepository entity.IGrupoRepository) *CriarGrupoUseCase {
	return &CriarGrupoUseCase{
		GrupoRepository: grupoRepository,
	}
}

func (u *CriarGrupoUseCase) Execute(input dto.GrupoInputDto) (*dto.GrupoOutputDto, error) {
	grupo := entity.NovoGrupo(input)

	r, err := u.GrupoRepository.CriarGrupo(grupo)

	if err != nil {
		return nil, err
	}
	return &dto.GrupoOutputDto{
		GrupoID:     r.GrupoID,
		Nome:        r.Nome,
		Propietario: r.Propietario,
	}, nil
}
