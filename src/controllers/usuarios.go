package controllers

import (
	"api/src/banco"
	"api/src/model"
	"api/src/repository"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)

	usuarioId, erro := repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	idInseridoStr := fmt.Sprintf("Id inserido: %d", usuarioId)
	respostas.JSON(w, http.StatusCreated, idInseridoStr)

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	usuariosEncontrados, erro:=  repositorio.BuscaPorNomeOuNick(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuariosEncontrados)

}

func Buscarusuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando um usuario"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando usuario"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deletando usuario"))
}
