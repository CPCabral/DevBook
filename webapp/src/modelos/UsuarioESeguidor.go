package modelos

func UsuarioESeguidor(usuarioLogadoID uint64, seguidores []Usuario) bool{
	for _, seguidor := range seguidores {
		if seguidor.ID == usuarioLogadoID {
			return true
		}
	}
	return false
}