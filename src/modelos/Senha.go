package modelos

// Senha representa o formato da requisição de senha para atualizar a senha de um usuário no banco de dados
type Senha struct {
	Atual string `json:"atual"`
	Nova   string `json:"nova"`
}