package streaming

import (
	"errors"
	"streamingsystem/contenido"
	"streamingsystem/usuarios"
)

type PlataformaStreaming struct {
	usuarios   map[string]*usuarios.Usuario
	contenidos []*contenido.Contenido
}

func NuevaPlataforma() *PlataformaStreaming {
	return &PlataformaStreaming{
		usuarios:   make(map[string]*usuarios.Usuario),
		contenidos: []*contenido.Contenido{},
	}
}

func (p *PlataformaStreaming) RegistrarUsuario(u *usuarios.Usuario) {
	p.usuarios[u.GetNombre()] = u
}

func (p *PlataformaStreaming) AgregarContenido(c *contenido.Contenido) {
	p.contenidos = append(p.contenidos, c)
}

func (p *PlataformaStreaming) Reproducir(nombreUsuario string, titulo string) (string, error) {

	u, existe := p.usuarios[nombreUsuario]
	if !existe {
		return "", errors.New("usuario no encontrado")
	}

	if !u.PuedeAcceder() {
		return "", errors.New("usuario sin suscripción premium")
	}

	for _, c := range p.contenidos {
		if c.GetTitulo() == titulo {
			u.AgregarHistorial(titulo)
			return c.Reproducir(), nil
		}
	}

	return "", errors.New("contenido no encontrado")
}
