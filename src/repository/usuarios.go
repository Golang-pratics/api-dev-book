package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// cria um repositorio de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario model.Usuario) (uint64, error) {
	query := "INSERT INTO usuarios (nome, nick, email, senha) VALUES($1, $2, $3, $4) RETURNING id"
	row := repositorio.db.QueryRow(query, usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	var idInserido uint64
	err := row.Scan(&idInserido)
	if err != nil {
		return 0, err
	}

	return uint64(idInserido), nil

}

// busca todos os usuários que atendem pelo filto de nome ou nick
func (repositorio usuarios) BuscaPorNomeOuNick(nomeOuNick string) ([]model.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeounick%

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE $1 OR nick LIKE $2",
		nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []model.Usuario
	for linhas.Next() {
		var usuario model.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, nil

}

// Buscar por id traz um usuário do banco de dados
func (repositorio usuarios) BuscarPorID(ID uint64) (model.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = $1",
		ID,
	)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario model.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return model.Usuario{}, erro
		}
	}

	return usuario, nil

}

func (repositorio usuarios) AtualizarUsuario(ID uint64, usuario model.Usuario) error {
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET nome = $1, nick = $2, email = $3 WHERE id = $4")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil

}

func (repositorio usuarios)Deletar(ID uint64) error{
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _,erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios)BuscaPorEmail(email string) (model.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = $1", email)
	if erro != nil {
		 return model.Usuario{}, erro
	}

	defer linha.Close()

	var usuario model.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return model.Usuario{}, erro
		}
	}

	return usuario, nil
}
// seguir permite que um usuario siga outro
func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert into seguidores (usuario_id, seguidor_id) values ($1, $2) ON CONFLICT DO NOTHING;",
	)
	if erro != nil {
		return erro
   }
   defer statement.Close()

   if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
   }

   return nil

}