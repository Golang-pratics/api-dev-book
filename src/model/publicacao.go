package model

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicação feita por um usuario
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorID,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	AutorNick string    `json:"AutorNick,omitempty"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}


func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil{
		return erro
	}

	publicacao.formatar()

	return nil
}
// validar verifica se os campos da publicação estão preenchidos
func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == ""{
		return errors.New("O titulo é obrigatório e não pode estar em branco")
	}

	if publicacao.Conteudo == ""{
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}
	return nil
}
// formatar irá formatar os campos da publicação
func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}