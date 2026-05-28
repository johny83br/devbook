package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	urlConexao := config.StringConexaoBanco
	db, err := abrirConexao(urlConexao)
	return db, err
}

func abrirConexao(url string) (*sql.DB, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}