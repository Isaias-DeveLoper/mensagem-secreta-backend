package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/infra/database"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/infra/repository"
	grupo_usecase "github.com/Isaias-Developer/mensagem-secreta-backend/src/api/usecases/grupo"
	mensagem_usecase "github.com/Isaias-Developer/mensagem-secreta-backend/src/api/usecases/mensagem"
	usuario_usecase "github.com/Isaias-Developer/mensagem-secreta-backend/src/api/usecases/usuario"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/web/controllers"
	"github.com/Isaias-Developer/mensagem-secreta-backend/src/api/web/routes"
)

func Run() {
	database.ConectarBancoDeDados()

	usuarioRepository := repository.UsuarioRepositoryConstruct(database.DB)
	grupoRepository := repository.GrupoRepositoryConstruct(database.DB)
	mensagemRepository := repository.MensagemRepositoryConstruct(database.DB)

	criarUsuario := usuario_usecase.CriarUsuarioConstruct(usuarioRepository)
	listarUsuario := usuario_usecase.ListarUsuarioConstruct(usuarioRepository)
	listarUsuarioUsername := usuario_usecase.ListarPorUsernameConstruct(usuarioRepository)
	atualizarUsuario := usuario_usecase.AtualizarUsuarioConstruct(usuarioRepository)
	excluirUsuario := usuario_usecase.ExcluirUsuarioConstruct(usuarioRepository)

	listarGrupos := grupo_usecase.ListarGrupoConstruct(grupoRepository)
	listarGruposPropietario := grupo_usecase.ListarGruposPropietarioConstruct(grupoRepository)
	criarGrupo := grupo_usecase.CriarGrupoConstruct(grupoRepository)

	enviarMensagem := mensagem_usecase.EnviarMensagemConstruct(mensagemRepository)
	lerMensagens := mensagem_usecase.LerMensagensConstruct(mensagemRepository)

	usuarioController := controllers.UsuarioControllerConstruct(criarUsuario, listarUsuario, listarUsuarioUsername, atualizarUsuario, excluirUsuario)
	grupoController := controllers.GrupoControllerConstruct(listarGrupos, listarGruposPropietario, criarGrupo)
	mensagemController := controllers.MensagemControllerConstruct(enviarMensagem,lerMensagens)

	log.Println(fmt.Sprintf("Aplicação rodando na porta :" + os.Getenv("PORT")))
	
	routes.IniciarRotas(usuarioController, grupoController, mensagemController)

	routes.Run()
}
