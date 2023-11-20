package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	auth_usecase "github.com/Isaias-Developer/mensagem-secreta-backend/src/api/usecases/auth"
)

type AuthController struct {
	LoginUseCase *auth_usecase.LoginUseCase
}

func AuthControllerConstruct(loginUseCase *auth_usecase.LoginUseCase) *AuthController {
	return &AuthController{
		LoginUseCase: loginUseCase,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var usuario dto.UsuarioInputDto
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, e := c.LoginUseCase.Execute(usuario)
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(result)
}
