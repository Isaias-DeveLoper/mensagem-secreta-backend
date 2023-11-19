package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/infra/database"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/infra/repository"
	usuario_usecase "github.com/Isaias-Developer/mensagem-secreta-backend/src/api/usecases/usuario"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/web/controllers"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/web/routes"
)

func Run() {
	database.ConectarBancoDeDados()

	usuarioRepository := repository.Construct(database.DB)

	criarUsuario := usuario_usecase.CriarUsuarioConstruct(usuarioRepository)
	listarUsuario := usuario_usecase.ListarUsuarioConstruct(usuarioRepository)
	listarUsuarioUsername := usuario_usecase.ListarPorUsernameConstruct(usuarioRepository)
	atualizarUsuario := usuario_usecase.AtualizarUsuarioConstruct(usuarioRepository)
	excluirUsuario := usuario_usecase.ExcluirUsuarioConstruct(usuarioRepository)

	usuarioController := controllers.UsuarioConstruct(criarUsuario,listarUsuario,listarUsuarioUsername,atualizarUsuario,excluirUsuario)

	log.Println(fmt.Sprintf("Aplicação rodando na porta :" + os.Getenv("PORT")))
	routes.IniciarRotas(usuarioController)
	routes.Run()
}

