package entity

import "github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"

type Mensagem struct {
	GroupID string
	Texto   string
}

type IMensagemRepository interface {
	LerMensagens(grupo,senha string)([]*Mensagem,error)
	EnviarMensagem(mensagem *Mensagem) (*Mensagem, error)
}

func NovaMensagem(mensagem dto.MensagemInputDto) *Mensagem{
	return &Mensagem{
		GroupID: mensagem.GroupID,
		Texto: mensagem.Texto,
	}
}