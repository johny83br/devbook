package repositories

import (
	"api/src/models"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = "%" + nomeOuNick + "%"

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio usuarios) BuscarPorID(usuarioID uint64) (models.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?",
		usuarioID,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio usuarios) Atualizar(usuarioID uint64, usuario models.Usuario) error {
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) Deletar(usuarioID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuarioID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"SELECT id, senha FROM usuarios WHERE email = ?",
		email,
	)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuarioID, seguidorID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuarioID, seguidorID)
	if erro != nil {
		return erro
	}

	return nil
}