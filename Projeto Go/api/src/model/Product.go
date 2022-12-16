package model

import (
	"errors"
	"strings"
)

type Produto struct {
	ID    uint64  `json: "id,omitempty"`
	Name  string  `json: "name,omitempty"`
	Code  string  `json: "code,omitempty"`
	Price float64 `json: "price,omitempty"`
}

func (produto *Produto) Preparar() error {
	if erro := produto.validar(); erro != nil {
		return erro
	}

	produto.formatar()
	return nil
}

func (produto *Produto) validar() error {
	if produto.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if produto.Code == "" {
		return errors.New("A marca é obrigatória e não pode estar em branco")
	}

	if produto.Price <= 0 {
		return errors.New("O preço não pode ser zero ou negativo")
	}

	return nil
}

func (produto *Produto) formatar() {
	produto.Name = strings.TrimSpace(produto.Name)
	produto.Code = strings.TrimSpace(produto.Code)
}
