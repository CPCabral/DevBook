package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicação representa uma publicação feita por um usuário
type Publicacao struct {
	ID        uint64 `json:"id,omitempty"`
	Titulo    string `json:"titulo,omitempty"`
	Conteudo  string `json:"conteudo,omitempty"`
	AutorID   uint64 `json:"autorId,omitempty"`
	AutorNick string `json:"autorNick,omitempty"`
	Curtidas  uint64 `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm"`
}

func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.Validar(); erro != nil {
		return erro
	}
	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) Validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o título da publicação não pode estar vazio")
	}
	if publicacao.Conteudo == "" {
		return errors.New("o conteúdo da publicação não pode estar vazio")
	}
	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
	publicacao.CriadaEm = time.Now()	
}