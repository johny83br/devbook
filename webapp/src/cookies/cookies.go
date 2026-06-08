package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configurar() {
	s = securecookie.New(config.HASH_KEY, config.BLOCK_KEY)
}

func Salvar(w http.ResponseWriter, nome string, valor string) error {
	dados := map[string]string{
		"id":    nome,
		"token": valor,
	}

	encoded, erro := s.Encode("cookie-devbook", dados)

	if erro != nil {
		return erro
	}

	cookie := &http.Cookie{
		Name:     "cookie-devbook",
		Value:    encoded,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)

	return nil
}
