# 📝 Primera API REST en Go

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License](https://img.shields.io/badge/Licencia-MIT-green.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Estado-En%20desarrollo-yellow)]()
[![Made with](https://img.shields.io/badge/Made%20with-Go-blue?logo=go)](https://go.dev/)
[![Contributions welcome](https://img.shields.io/badge/Contribuciones-Bienvenidas-brightgreen.svg)]()

---

Esta es mi primera **API REST** construida con el lenguaje de programación **Go** 🦦.  
El objetivo del proyecto es **aprender los fundamentos del desarrollo backend con Go**, incluyendo:

- Creación de un **servidor web básico**.  
- Manejo de **rutas y peticiones HTTP** (`GET` y `POST`).  
- Gestión de **recursos en memoria**.  
- Uso del paquete `context` para **controlar el ciclo de vida** de las peticiones.

---

## 🚀 Despliegue local

### 🔧 Prerrequisitos
Asegúrate de tener instalado [Go](https://go.dev/doc/install) en tu sistema.

### ▶️ Pasos para ejecutar

```bash
# Clonar el repositorio
git clone https://github.com/vgutierrezz/API-GO.git

# Navegar al directorio del proyecto
cd API-GO

# Ejecutar el servidor
go run main.go
```
La API estará disponible en http://localhost:8080 🖥️

--- 
## ⚙️ Endpoints de la API

### 📍 GET /users
Devuelve una lista de todos los usuarios almacenados en memoria.

Ejemplo de respuesta (200 OK):

``` bash
{
  "status": 200,
  "data": [
    { "id": 1, "first_name": "John", "last_name": "Doe", "email": "john.doe@example.com" },
    { "id": 2, "first_name": "Jane", "last_name": "Smith", "email": "jane.smith@example.com" },
    { "id": 3, "first_name": "Alice", "last_name": "Johnson", "email": "alice.johnson@example.com" }
  ]
}
```

### 📨 POST /users
Crea un nuevo usuario.

Cuerpo de la petición:

``` bash
{
  "first_name": "Nuevo",
  "last_name": "Usuario",
  "email": "nuevo.usuario@example.com"
}
```

Ejemplo de respuesta (201 Created):

``` bash
{
  "status": 201,
  "data": {
    "id": 4,
    "first_name": "Nuevo",
    "last_name": "Usuario",
    "email": "nuevo.usuario@example.com"
    }
}

```

Validación de errores (400 Bad Request):

``` bash
{
  "status": 400,
  "message": "first name is required"
}
```

### 📦 Uso del paquete context
El archivo interpkgs/context/main.go incluye un ejemplo independiente del uso del paquete context, fundamental para:

- *Controlar* el ciclo de vida de las peticiones.

- *Cancelar* operaciones automáticamente cuando expiran los plazos.

- *Pasar valores específicos* entre funciones durante la ejecución.

#### Ejemplos incluidos:
- context.WithValue → almacena y recupera valores en el contexto.

- context.WithTimeout → define un tiempo límite para una operación.

Ejecutar el ejemplo:

bash
``` bash
cd interpkgs/context
go run main.go
```

💻 Tecnologías utilizadas
| Tecnología | Descripción |
| ---------- | ----------- |
| 🐹 Go |	Lenguaje principal del proyecto
| 🌐 net/http |	Paquete estándar para crear servidores web
| 📦 encoding/json |	Manejo de datos JSON
| ⏱️ context |	Control del ciclo de vida de las peticiones

--- 

### 🤝 Contribuciones
¡Las contribuciones, sugerencias y reportes de errores son bienvenidos!
Podés abrir un issue o enviar un pull request 💡

--- 

### ✍️ Autor
#### Valentina Gutierrez
---
Curso Fundamentos de GO - Go Web