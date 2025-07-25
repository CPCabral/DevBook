package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublicacoes = []Rota{
	{
		URI:    "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}/curtir",
		Metodo: http.MethodPost,
		Funcao: controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}/descurtir",
		Metodo: http.MethodPost,
		Funcao: controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}/atualizar",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPaginaAtualizacaoDePublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.ExcluirPublicacao,
		RequerAutenticacao: true,
	},
}