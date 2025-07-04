package repositorios

import (
	"api/src/modelos"

	"database/sql"
	"fmt"
)

// Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
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

// Buscar traz todos os usuarios que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string ) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios =  append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarPorID traz um usuario do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil{
			return modelos.Usuario{}, erro
		}		
	}
	return usuario, nil
}

// BuscarPorEmail busca um usuario por email e retorna o seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
		if erro != nil {
			return modelos.Usuario{}, erro
		}
		defer linha.Close()

		var usuario modelos.Usuario

		if linha.Next() {
			if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
				return modelos.Usuario{}, erro
			}			
		}

		return usuario, nil
}

// Atualizar altera as informações do usuario no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error{
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil{
		return erro
	}
	return nil
}

// Deletera exclui as informações de um usuario no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

// Seguir permite que um usuario siga outro
func (repositorio Usuarios) Seguir(usuarioID, seguidorID  uint64) error{
	statement, erro := repositorio.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",	
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

// PararDeSeguir permite que um usuario siga outro
func (repositorio Usuarios) PararDeSeguir(usuarioID, seguidorID  uint64) error{
	statement, erro := repositorio.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",	
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

// BuscarSeguidores mostra todos os usuarios que estão seguindo um usuario id
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.CriadoEm
	from usuarios u inner join seguidores s on u.id = s.seguidor_id
	where s.usuario_id = ?
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	
	var seguidores []modelos.Usuario
	for linhas.Next() {
		var seguidor modelos.Usuario
		if erro = linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, seguidor)
	}
	return seguidores, nil
}

// BuscarSeguindo mostra todos os usuarios que um usuario id está seguindo
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.CriadoEm
	from usuarios u inner join seguidores s on u.id = s.usuario_id
	where s.seguidor_id = ?
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	
	var seguidores []modelos.Usuario
	for linhas.Next() {
		var seguidor modelos.Usuario
		if erro = linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, seguidor)
	}
	return seguidores, nil
}

// BuscarSenha traz a senha de um usuario pelo seu ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro			
		}
	}

	return usuario.Senha, nil
}

func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}
	return nil

}