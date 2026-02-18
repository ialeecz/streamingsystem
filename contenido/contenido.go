package contenido

type Reproducible interface {
	Reproducir() string
}

type Contenido struct {
	titulo    string
	categoria string
}

func NuevoContenido(titulo string, categoria string) *Contenido {
	return &Contenido{
		titulo:    titulo,
		categoria: categoria,
	}
}

func (c *Contenido) Reproducir() string {
	return "Reproduciendo: " + c.titulo
}

func (c *Contenido) GetTitulo() string {
	return c.titulo
}
