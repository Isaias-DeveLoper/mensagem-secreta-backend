package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "os"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/web/controllers"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/web/middleware"
	"github.com/go-chi/chi/v5"
)

var r *chi.Mux

func IniciarRotas(authController *controllers.AuthController,
	usuarioController *controllers.UsuarioController,
	grupoController *controllers.GrupoController,
	mensagemController *controllers.MensagemController) {
	r = chi.NewRouter()

	r.Use(middleware.ContentType)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"mensagem": "Seja bem vindo!"})
	})

	r.Route("/v1", func(r chi.Router) {

		r.Post("/login",authController.Login)

		r.Route("/user", func(r chi.Router) {
			r.Get("/users", usuarioController.GetUsers)
			r.Get("/users_u", usuarioController.GetUserUsername)
			r.Put("/atualizar", usuarioController.PutUser)
			r.Post("/criar", usuarioController.PostUser)
			r.Delete("/excluir", usuarioController.DeleteUser)
		})

		r.Route("/group", func(r chi.Router) {
			r.Get("/grupos", grupoController.GetGrupos)
			r.Get("/grupos_p", grupoController.GetGruposPropietario)
			r.Post("/criar", grupoController.PostGrupo)
		})

		r.Route("/mensagem", func(r chi.Router) {
			r.Get("/ler", mensagemController.GetMensagens)
			r.Post("/enviar", mensagemController.PostMensagem)
		})
	})
}

func Run() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":"+"8080"), r))
}
