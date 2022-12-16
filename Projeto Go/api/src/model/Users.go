package model

import (
	"api/src/security"
	"errors"
	"strings"
)

type Usuario struct {
	ID      uint64 `json:"id:omitempty`
	Usuario string `json:"usuario:omitempty`
	Senha   string `json:"senha:omitempty`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) PrepararUsuario(etapa string) error {
	if erro := usuario.validarUsuario(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatarUsuario(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validarUsuario(etapa string) error {
	if usuario.Usuario == "" {
		return errors.New("O usuário é obrigatório e não pode estar em branco")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatarUsuario(etapa string) error {
	usuario.Usuario = strings.TrimSpace(usuario.Usuario)

	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
