package controllers

import "net/http"

// CarregarTelaDeLogin carrega a tela de login para o usuario
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de Login"))
}