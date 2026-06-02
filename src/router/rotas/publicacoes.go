package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                "/api/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes/{publicacaoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/api/publicacoes",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true,
	},
}
