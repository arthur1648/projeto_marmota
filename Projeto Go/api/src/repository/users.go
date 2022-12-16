package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) CriarUsuario(usuario model.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (usuario, senha) values (?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Usuario, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorio Usuarios) BuscarUsuario(usuario string) ([]model.Usuario, error) {
	usuario = fmt.Sprintf("%%%s%%", usuario)

	linhas, erro := repositorio.db.Query(
		"select id, usuario, senha from usuarios where usuario LIKE ?",
		usuario,
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
			&usuario.Usuario,
			&usuario.Senha,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorUsuario busca um usuário pelo usuário e retorna o seu ID e senha com hash
func (repositorio Usuarios) BuscarPorUsuario(usuario string) (model.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where usuario = ?", usuario)
	if erro != nil {
		return model.Usuario{}, erro
	}
	defer linha.Close()

	var modelUsuario model.Usuario

	if linha.Next() {
		if erro = linha.Scan(&modelUsuario.ID, &modelUsuario.Senha); erro != nil {
			return model.Usuario{}, erro
		}
	}
	return modelUsuario, nil
}
