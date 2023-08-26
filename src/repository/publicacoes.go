package repository

import (
	"api/src/model"
	"database/sql"
)

// Publicacoes representa um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}

// novo repositorios de publicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// criar insere uma publicação no banco de dados
func(repositorio Publicacoes) Criar( publicacao model.Publicacao) (uint64, error){
	query := "INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES($1, $2, $3) RETURNING id"
	statement := repositorio.db.QueryRow(query, publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)

	// idInserido, erro := resultado.LastInsertId()
	// if erro != nil{
	// 	return 0, erro
	// }

	// return uint64(idInserido), nil


	var idInserido uint64
	err := statement.Scan(&idInserido)
	if err != nil {
		return 0, err
	}

	return uint64(idInserido), nil
}


func (repositorios Publicacoes) BuscarPorID(publicacaoID uint64) (model.Publicacao, error){
	linha, erro := repositorios.db.Query(`
		select p.*, u.nick from publicacoes p inner join usuarios u
		on u.id = p.autor_id where p.id = $1`,
		publicacaoID)

	if erro != nil{
		return model.Publicacao{}, erro
	}

	defer linha.Close()

	var publicacao model.Publicacao

	if linha.Next(){
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.AutorID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.CriadaEm,
			&publicacao.Curtidas,
			&publicacao.AutorNick,
		); erro != nil{
			return model.Publicacao{}, erro
		}
	}

	return publicacao, nil
}