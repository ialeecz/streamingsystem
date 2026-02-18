\# 🎬 Sistema de Gestión de Streaming



\## 📖 Descripción del Proyecto



El presente proyecto corresponde al desarrollo de un Sistema de Gestión de Streaming implementado en el lenguaje Go.



El sistema fue desarrollado aplicando los principios de:



\- Programación Orientada a Objetos

\- Encapsulación

\- Interfaces

\- Manejo de errores

\- Estructuras de datos

\- Arquitectura modular por paquetes



---



\## 🧠 Objetivo Académico



Aplicar estructuras de datos en la resolución de problemas, implementando adecuadamente interfaces dentro de un entorno de programación orientado a objetos.



---



\## 🏗 Arquitectura del Sistema



El proyecto está organizado en paquetes independientes:



streamingsystem/

│ main.go

│ go.mod

│ README.md

│

├── usuarios/

│ └── usuario.go

│

├── contenido/

│ └── contenido.go

│

├── streaming/

│ └── plataforma.go



---



\## 🗂 Estructuras de Datos Utilizadas



\- `map\[string]\*Usuario` → Para almacenar usuarios registrados.

\- `\[]\*Contenido` → Para gestionar el catálogo de contenidos.

\- `\[]string` → Para almacenar el historial de reproducción.



Estas estructuras permiten eficiencia en búsqueda y almacenamiento dinámico.



---



\## 🔐 Encapsulación



Los atributos de las estructuras fueron declarados privados (iniciando con letra minúscula).



El acceso se realiza mediante métodos públicos como:



\- `GetNombre()`

\- `PuedeAcceder()`



Esto garantiza control sobre la información interna del sistema.



---



\## 🧩 Interfaces Implementadas



Se definió la interfaz:



```go

type Reproducible interface {

&nbsp;   Reproducir() string

}



