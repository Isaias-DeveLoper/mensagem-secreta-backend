package dto

type UsuarioInputDto struct {
	Username string `json:"username"`
	Senha    string `json:"senha"`
}

type UsuarioOutputDto struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Senha    string `json:"senha"`
}

type UsuarioAtualizadoDto struct {
	UsernameAtual string `json:"username_atual"`
	UsernameNovo  string `json:"username_novo"`
	SenhaAtual    string `json:"senha_atual"`
	SenhaNova     string `json:"senha_nova"`
}
