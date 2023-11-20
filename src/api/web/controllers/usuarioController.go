package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	usuario_usecase "github.com/Isaias-Developer/mensagem-secreta-backend/src/api/usecases/usuario"
)

type UsuarioController struct {
	CriarUsuarioUseCase           *usuario_usecase.CriarUsuarioUseCase
	ListarUsuariosUseCase         *usuario_usecase.ListarUsuariosUseCase
	ListarUsuariosUsernameUseCase *usuario_usecase.ListarUsuarioPorUsernameUseCase
	AtualizarUsuarioUseCase       *usuario_usecase.AtualizarUsuarioUseCase
	ExcluirUsuarioUseCase         *usuario_usecase.ExcluirUsuarioUseCase
}

func UsuarioControllerConstruct(criarUsuarioUseCase *usuario_usecase.CriarUsuarioUseCase,
	listarUsuariosUseCase *usuario_usecase.ListarUsuariosUseCase,
	listarUsuariosUsernameUseCase *usuario_usecase.ListarUsuarioPorUsernameUseCase,
	atualizarUsuarioUseCase *usuario_usecase.AtualizarUsuarioUseCase,
	exluirUsuarioUseCase *usuario_usecase.ExcluirUsuarioUseCase) *UsuarioController {
	return &UsuarioController{
		CriarUsuarioUseCase:           criarUsuarioUseCase,
		ListarUsuariosUseCase:         listarUsuariosUseCase,
		ListarUsuariosUsernameUseCase: listarUsuariosUsernameUseCase,
		AtualizarUsuarioUseCase:       atualizarUsuarioUseCase,
		ExcluirUsuarioUseCase:         exluirUsuarioUseCase,
	}
}

func (c *UsuarioController) PostUser(w http.ResponseWriter, r *http.Request) {
	var usuario dto.UsuarioInputDto
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	novoUsuario, err := c.CriarUsuarioUseCase.Execute(usuario)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(&novoUsuario)
}

func (c *UsuarioController) GetUsers(w http.ResponseWriter, r *http.Request) {
	usuarios, err := c.ListarUsuariosUseCase.Execute()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(usuarios)
}

func (c *UsuarioController) GetUserUsername(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")

	usuarios, err := c.ListarUsuariosUsernameUseCase.Execute(username)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(usuarios)
}

func (c *UsuarioController) PutUser(w http.ResponseWriter, r *http.Request) {
	var usuario dto.UsuarioAtualizadoDto
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	result, e := c.AtualizarUsuarioUseCase.Execute(usuario)
	if e != nil {
		http.Error(w, e.Error(), 500)
	}

	json.NewEncoder(w).Encode(result)
}

func (c *UsuarioController) DeleteUser(w http.ResponseWriter, r *http.Request){
	var usuario dto.UsuarioInputDto
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, e := c.ExcluirUsuarioUseCase.Execute(usuario)
	if e != nil {
		http.Error(w, e.Error(), 500)
	}

	json.NewEncoder(w).Encode(result)
}