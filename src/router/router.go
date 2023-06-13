package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

//gerar vai retornar um router com as rotas da api
func Gerar() *mux.Router{
	r := mux.NewRouter()
	return rotas.Configurar(r)
}

