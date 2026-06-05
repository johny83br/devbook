package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"webapp/src/respostas"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	nome := r.FormValue("nome")
	email := r.FormValue("email")
	nick := r.FormValue("nick")
	senha := r.FormValue("senha")
	confirmarSenha := r.FormValue("confirmar_senha")

	if senha != confirmarSenha {
		log.Fatal(errors.New("As senhas não conferem"))
		return
	}

	usuario, erro := json.Marshal(map[string]string{
		"nome":  nome,
		"email": email,
		"nick":  nick,
		"senha": senha,
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	response, erro := http.Post("http://localhost:5000/api/usuarios", "application/json", bytes.NewBuffer(usuario))

	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, response)
		return
	}

	respostas.JSON(w, http.StatusCreated, nil)

}
