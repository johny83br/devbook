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

func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("cookie-devbook")
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string)
	if erro = s.Decode("cookie-devbook", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}
