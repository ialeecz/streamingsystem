package usuarios

import "errors"

type Usuario struct {
	nombre      string
	suscripcion string
	historial   []string
}

func NuevoUsuario(nombre string, suscripcion string) (*Usuario, error) {
	if nombre == "" {
		return nil, errors.New("el nombre no puede estar vacío")
	}

	return &Usuario{
		nombre:      nombre,
		suscripcion: suscripcion,
		historial:   []string{},
	}, nil
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) PuedeAcceder() bool {
	return u.suscripcion == "premium"
}

func (u *Usuario) AgregarHistorial(titulo string) {
	u.historial = append(u.historial, titulo)
}
