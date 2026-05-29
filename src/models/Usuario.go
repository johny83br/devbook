package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID 				uint64 			`json:"id,omitempty"`
	Nome 			string 			`json:"nome,omitempty"`
	Nick 			string 			`json:"nick,omitempty"`
	Email 		string 			`json:"email,omitempty"`
	Senha 		string 			`json:"senha,omitempty"`
	CriadoEm 	time.Time 	`json:"criadoEm,omitempty"`
}

func (u *Usuario) Preparar(etapa string) error {
	if err := u.validar(etapa); err != nil {
		return err
	}

	u.formatar()

	return nil
}

func (u *Usuario) validar(etapa string) error {

	if u.Nome == "" {
		return errors.New("O nome é obrigatório e não pode ser vazio")
	}

	if u.Nick == "" {
		return errors.New("O nickname é obrigatório e não pode ser vazio")
	}

	if u.Email == "" {
		return errors.New("O email é obrigatório e não pode ser vazio")
	}

	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("A senha é obrigatória e não pode ser vazio")
	}

	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	u.Senha = strings.TrimSpace(u.Senha)
}