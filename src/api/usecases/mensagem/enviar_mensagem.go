package mensagem_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type EnviarMensagemUseCase struct {
	MensagemRepository entity.IMensagemRepository
}

func EnviarMensagemConstruct(mensagemRepository entity.IMensagemRepository) *EnviarMensagemUseCase {
	return &EnviarMensagemUseCase{
		MensagemRepository: mensagemRepository,
	}
}

func (u *EnviarMensagemUseCase) Execute(mensagem dto.MensagemInputDto) (*dto.MensagemInputDto, error) {
	novaMensagem := entity.NovaMensagem(mensagem)
	r, err := u.MensagemRepository.EnviarMensagem(novaMensagem)
	if err != nil {
		return nil, err
	}
	return &dto.MensagemInputDto{
		GroupID: r.GroupID,
		Texto:   r.Texto,
	}, nil
}
