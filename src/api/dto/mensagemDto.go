package dto

type MensagemInputDto struct {
	GroupID string `json:"group_id"`
	Texto   string `json:"texto"`
}

type MensagemOutputDto struct {
	Texto string `json:"texto"`
}
