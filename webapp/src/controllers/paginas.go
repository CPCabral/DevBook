package controllers

import (
	"fmt"
	"net/http"
	"webapp/config"
	"webapp/src/requisicoes"
	"webapp/src/utils"
)

// CarregarTelaDeLogin carrega a tela de login para o usuario
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "login.html", nil)
}

func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal carrega a pagina principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, erro)
	defer response.Body.Close()
	
	utils.ExecutarTemplates(w, "home.html", nil)
}

// CarregarPaginaDePerfil carrega a pagina de perfil do usuario