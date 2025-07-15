package modelos

// DadosAutenticacao contém o token e o id do usuario autenticado
type DadosAutenticacao struct {
	ID string `json: "id"`	
	Token string `json: "Token"`
}
