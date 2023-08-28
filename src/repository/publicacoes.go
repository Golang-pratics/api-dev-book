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
func (repositorio Publicacoes) Criar(publicacao model.Publicacao) (uint64, error) {
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

func (repositorios Publicacoes) BuscarPorID(publicacaoID uint64) (model.Publicacao, error) {
	linha, erro := repositorios.db.Query(`
		select p.*, u.nick from publicacoes p inner join usuarios u
		on u.id = p.autor_id where p.id = $1`,
		publicacaoID)

	if erro != nil {
		return model.Publicacao{}, erro
	}

	defer linha.Close()

	var publicacao model.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.AutorID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.CriadaEm,
			&publicacao.Curtidas,
			&publicacao.AutorNick,
		); erro != nil {
			return model.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar traz as publicacoes dos usuarios seguidos e tambpem do usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]model.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select distinct p.*, u.nick from publicacoes p
	inner join usuarios u on u.id = p.autor_id
	inner join seguidores s on p.autor_id = s.usuario_id
	where u.id = $1 or s.seguidor_id = $2
	order by 1 desc`,
		usuarioID, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []model.Publicacao

	for linhas.Next() {
		var publicacao model.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.AutorID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.CriadaEm,
			&publicacao.Curtidas,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil

}

// Atualizar altera os dados de uma publicação no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao model.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = $1, conteudo = $2 where id = $3")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil

}

// Deletar exclui uma publicação do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacoes where id = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil

}

// BuscarPorUsuario traz as publicações de um usuário específico do banco de dados
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]model.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p 
		join usuarios u on u.id = p.autor_id
		where p.autor_id = $1	
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []model.Publicacao

	for linhas.Next() {
		var publicacao model.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.AutorID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.CriadaEm,
			&publicacao.Curtidas,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil

}


func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}
// Descurtir subtrai a quantidade de curtidas de uma publicação
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
		update publicacoes set curtidas =	
		CASE 
			WHEN curtidas > 0 THEN curtidas - 1 
			ELSE 0 
		END
		where id = $1
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}