package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/model"
	"api/src/repository"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarPublicacao cria uma nova publicacao no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request){
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil{
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao model.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID
	if erro = publicacao.Preparar(); erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)


}
// BuscarPublicacoes traz as publicacoes que apareceriam no feed do usuario
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request){
	

}
// BuscarPublicacao traz uma unica publicacao do banco de dados
func BuscarPublicacao(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)
	pubicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDePublicacoes(db)
	publicacao, erro := repositorio.BuscarPorID(pubicacaoID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publicacao)
	

}
// AtualizarPublicacao altera os dados de uma publicacao no banco de dados
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request){
	

}
// DeletarPublicacao exclui uma publicacao do banco de dados
func DeletarPublicacao(w http.ResponseWriter, r *http.Request){
	

}
