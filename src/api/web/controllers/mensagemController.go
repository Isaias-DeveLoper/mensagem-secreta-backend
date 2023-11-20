package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	mensagem_usecase "github.com/Isaias-Developer/mensagem-secreta-backend/src/api/usecases/mensagem"
)

type MensagemController struct {
	EnviarMensagemUseCase *mensagem_usecase.EnviarMensagemUseCase
	LerMensagensUseCase   *mensagem_usecase.LerMensagensUseCase
}

func MensagemControllerConstruct(enviarMensagemUseCase *mensagem_usecase.EnviarMensagemUseCase,
	lerMensagensUseCase *mensagem_usecase.LerMensagensUseCase) *MensagemController {
	return &MensagemController{
		EnviarMensagemUseCase: enviarMensagemUseCase,
		LerMensagensUseCase:   lerMensagensUseCase,
	}
}

func (c *MensagemController) PostMensagem(w http.ResponseWriter, r *http.Request) {
	var mensagem dto.MensagemInputDto
	if err := json.NewDecoder(r.Body).Decode(&mensagem); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := c.EnviarMensagemUseCase.Execute(mensagem)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (c *MensagemController) GetMensagens(w http.ResponseWriter, r *http.Request) {
	grupo := r.URL.Query().Get("grupo")
	senha := r.URL.Query().Get("senha")

	mensagens, err := c.LerMensagensUseCase.Execute(grupo,senha)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(mensagens)
}
