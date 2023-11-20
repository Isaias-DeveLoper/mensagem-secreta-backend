package mensagem_usecase

import (
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/entity"
)

type LerMensagensUseCase struct {
	MensagemRepository entity.IMensagemRepository
}

func LerMensagensConstruct(mensagemRepository entity.IMensagemRepository) *LerMensagensUseCase {
	return &LerMensagensUseCase{
		MensagemRepository: mensagemRepository,
	}
}

func (u *LerMensagensUseCase) Execute(grupo,senha string) ([]*dto.MensagemOutputDto, error) {

	var mensagens []*dto.MensagemOutputDto

	r, err := u.MensagemRepository.LerMensagens(grupo,senha)
	if err != nil {
		return nil, err
	}
	for _, mensagem := range r {
		mensagens = append(mensagens, &dto.MensagemOutputDto{
			Texto: mensagem.Texto,
		})
	}
	return mensagens, nil
}
