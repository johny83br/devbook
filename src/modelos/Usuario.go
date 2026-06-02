package modelos

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

	if erro := u.formatar(etapa); erro != nil {
		return erro
	}

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

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("O email inserido é inválido")
	}

	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("A senha é obrigatória e não pode ser vazio")
	}

	return nil
}

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if etapa == "cadastro" {
		u.Senha = strings.TrimSpace(u.Senha)
		senhaComHash, erro := security.Hash(u.Senha)
		if erro != nil {
			return erro
		}

		u.Senha = string(senhaComHash)
	}

	return nil
}