package dto

type GrupoInputDto struct {
	Nome        string `json:"nome"`
	Propietario string `json:"propietario"`
}

type GrupoOutputDto struct {
	GrupoID     string `json:"grupo_id"`
	Nome        string `json:"nome"`
	Propietario string `json:"propietario"`
}

type KeyNonce struct {
	Key   string `json:"keyw"`
	Nonce string `json:"nonce"`
}
