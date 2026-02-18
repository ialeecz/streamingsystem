package main

import (
	"fmt"
	"streamingsystem/contenido"
	"streamingsystem/streaming"
	"streamingsystem/usuarios"
)

func main() {

	plataforma := streaming.NuevaPlataforma()

	usuario, err := usuarios.NuevoUsuario("Kevin", "premium")
	if err != nil {
		fmt.Println(err)
		return
	}

	plataforma.RegistrarUsuario(usuario)

	video := contenido.NuevoContenido("Curso de Go", "Educativo")
	plataforma.AgregarContenido(video)

	resultado, err := plataforma.Reproducir("Kevin", "Curso de Go")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resultado)
}
