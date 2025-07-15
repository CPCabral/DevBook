package controllers

import (
	"net/http"
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
	utils.ExecutarTemplates(w, "home.html", nil)
}