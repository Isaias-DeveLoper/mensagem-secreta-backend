package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/dto"
	grupo_usecase "github.com/Isaias-Developer/mensagem-secreta-backend/src/api/usecases/grupo"
)

type GrupoController struct {
	ListarGruposUseCase            *grupo_usecase.ListarGrupoUseCase
	ListarGruposPropietarioUseCase *grupo_usecase.ListarGruposPorPropietarioUseCase
	CriarGrupoUseCase              *grupo_usecase.CriarGrupoUseCase
}

func GrupoControllerConstruct(listarGruposUseCase *grupo_usecase.ListarGrupoUseCase,
	listarGruposPropietarioUseCase *grupo_usecase.ListarGruposPorPropietarioUseCase,
	criarGrupoUseCase *grupo_usecase.CriarGrupoUseCase) *GrupoController {
	return &GrupoController{
		ListarGruposUseCase:            listarGruposUseCase,
		ListarGruposPropietarioUseCase: listarGruposPropietarioUseCase,
		CriarGrupoUseCase:              criarGrupoUseCase,
	}
}

func (c *GrupoController) GetGrupos(w http.ResponseWriter, r *http.Request) {
	grupos, err := c.ListarGruposUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(grupos)
}

func (c *GrupoController) PostGrupo(w http.ResponseWriter, r *http.Request) {
	var grupo dto.GrupoInputDto
	if err := json.NewDecoder(r.Body).Decode(&grupo); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := c.CriarGrupoUseCase.Execute(grupo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (c *GrupoController) GetGruposPropietario(w http.ResponseWriter, r *http.Request) {

	propietario := r.URL.Query().Get("propietario")
	grupos, err := c.ListarGruposPropietarioUseCase.Execute(propietario)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(grupos)
}
