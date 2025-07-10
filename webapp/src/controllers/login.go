package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaDeLogin carrega a tela de login para o usuario
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplates(w, "login.html", nil)
}