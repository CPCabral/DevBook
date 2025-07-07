package modelos

import "time"

// Publicação representa uma publicação feita por um usuário
type Publicação struct {
	ID        uint64 `json:"id,omitempty"`
	Título    string `json:"título,omitempty"`
	Conteudo  string `json:"conteudo,omitempty"`
	AutorID   uint64 `json:"autorId,omitempty"`
	AutorNick string `json:"autor_nick,omitempty"`
	Curtidas  uint64 `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm"`
}