package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasProdutos = []Rota{
	{
		URI:                "/produtos",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarProduto,
		RequerAutenticacao: false,
	},
	{
		URI:                "/produtos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarProdutos,
		RequerAutenticacao: false,
	},
	{
		URI:                "/produtos/{produtoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarProduto,
		RequerAutenticacao: false,
	},
	{
		URI:                "/produtos/{produtoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarProduto,
		RequerAutenticacao: false,
	},
	{
		URI:                "/produtos/{produtoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarProduto,
		RequerAutenticacao: false,
	},
}
