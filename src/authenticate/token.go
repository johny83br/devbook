package authenticate

import (
	"api/src/config"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = usuarioID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}

func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	if tokenString == "" {
		return jwt.ErrSignatureInvalid
	}
	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return err
	}
	fmt.Println(token)
	return nil
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(token) > 7 && token[:7] == "Bearer " {
		return token[7:]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}
	return config.SecretKey, nil
}