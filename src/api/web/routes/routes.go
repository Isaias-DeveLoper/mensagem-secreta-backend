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

func IniciarRotas(usuarioController *controllers.UsuarioController) {
	r = chi.NewRouter()

	r.Use(middleware.ContentType)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"mensagem": "Seja bem vindo!"})
	})
	
	r.Route("/v1", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Get("/users", usuarioController.GetUsers)
			r.Get("/users_u", usuarioController.GetUserUsername)
			r.Put("/atualizar", usuarioController.PutUser)
			r.Post("/criar", usuarioController.PostUser)
			r.Delete("/excluir", usuarioController.DeleteUser)
		})
	})
}

func Run() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":"+"8080"), r))
}
