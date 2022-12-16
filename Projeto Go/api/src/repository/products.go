package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
)

type Produtos struct {
	db *sql.DB
}

func NovoRepositorioDeProdutos(db *sql.DB) *Produtos {
	return &Produtos{db}
}

// Criar insere um produto no banco de dados
func (repositorio Produtos) Criar(produto model.Produto) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into produtos (name, code, price) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(produto.Name, produto.Code, produto.Price)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz todos os produtos que atendem um filtro de nome ou nick
func (repositorio Produtos) Buscar(nome string) ([]model.Produto, error) {
	nome = fmt.Sprintf("%%%s%%", nome)

	linhas, erro := repositorio.db.Query(
		"select id, name, code, price from produtos where name LIKE ?",
		nome,
	)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var produtos []model.Produto

	for linhas.Next() {
		var produto model.Produto

		if erro = linhas.Scan(
			&produto.ID,
			&produto.Name,
			&produto.Code,
			&produto.Price,
		); erro != nil {
			return nil, erro
		}

		produtos = append(produtos, produto)
	}

	return produtos, nil
}

// BuscarPorID traz um produto do banco de dados
func (repositorio Produtos) BuscarPorID(ID uint64) (model.Produto, error) {
	linhas, erro := repositorio.db.Query(
		"select id, name, code, price from produtos where id = ?",
		ID,
	)
	if erro != nil {
		return model.Produto{}, erro
	}
	defer linhas.Close()

	var produto model.Produto

	if linhas.Next() {
		if erro = linhas.Scan(
			&produto.ID,
			&produto.Name,
			&produto.Code,
			&produto.Price,
		); erro != nil {
			return model.Produto{}, erro
		}
	}

	return produto, nil
}

// Atualizar altera as informações de um produto no banco de dados
func (repositorio Produtos) Atualizar(ID uint64, produto model.Produto) error {
	statement, erro := repositorio.db.Prepare(
		"update produtos set name = ?, code = ?, price = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(produto.Name, produto.Code, produto.Price, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de um produto no banco de dados
func (repositorio Produtos) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from produtos where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
