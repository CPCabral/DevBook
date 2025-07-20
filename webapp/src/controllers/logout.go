package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// FazerLogout remove os dados de autenticação do cookie e redireciona para a página de login
func FazerLogout(w http.ResponseWriter, r *http.Request) {	
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}