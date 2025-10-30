📝 Primera API REST en Go
Esta es mi primera API REST construida con el lenguaje de programación Go. El objetivo de este proyecto es aprender a crear un servidor web básico, manejar rutas, procesar peticiones HTTP (GET y POST) y gestionar el estado de los recursos en memoria. Además, se incluye un ejemplo de uso del paquete context para la gestión de ciclos de vida de peticiones. 
🚀 Despliegue local
Para ejecutar esta API en tu entorno local, sigue estos sencillos pasos:
Prerrequisitos
Asegúrate de tener instalado Go en tu sistema.
Pasos
Clona el repositorio.
Navega hasta el directorio del proyecto.
Ejecuta el servidor con el siguiente comando:
bash
go run main.go
 

La API estará disponible en http://localhost:8080. 
⚙️ Endpoints de la API
/users
GET /users: Devuelve una lista de todos los usuarios almacenados en la memoria.
Respuesta de ejemplo (200 OK):
json
{
  "status": 200,
  "data": [
    {
      "id": 1,
      "first_name": "John",
      "last_name": "Doe",
      "email": "john.doe@example.com"
    },
    {
      "id": 2,
      "first_name": "Jane",
      "last_name": "Smith",
      "email": "jane.smith@example.com"
    },
    {
      "id": 3,
      "first_name": "Alice",
      "last_name": "Johnson",
      "email": "alice.johnson@example.com"
    }
  ]
}

POST /users: Crea un nuevo usuario.
Cuerpo de la petición:
json
{
  "first_name": "Nuevo",
  "last_name": "Usuario",
  "email": "nuevo.usuario@example.com"
}

Respuesta de ejemplo (201 Created):
json
{
  "status": 201,
  "data": {
    "id": 4,
    "first_name": "Nuevo",
    "last_name": "Usuario",
    "email": "nuevo.usuario@example.com"
  }
}

Validación de errores: Si falta algún campo (nombre, apellido o email), la API devolverá un error.
Respuesta de ejemplo (400 Bad Request):
json
{
  "status": 400,
  "message": "first name is required"
}

 
📦 Uso de context en Go
El archivo interpkgs/context/main.go incluye un ejemplo independiente de cómo utilizar el paquete context de Go. Este paquete es fundamental para gestionar el ciclo de vida de las peticiones, pasando valores específicos a través de las funciones y cancelando operaciones cuando expiran los plazos o el trabajo ya no es necesario. 
El código muestra dos ejemplos principales:
context.WithValue: Permite almacenar y recuperar datos específicos de una petición dentro del contexto.
context.WithTimeout: Configura una fecha límite para una operación. Si la operación no finaliza antes de la fecha límite, el contexto se cancela automáticamente, lo que puede ser útil para evitar que las operaciones de larga duración bloqueen los recursos indefinidamente. 
Para ejecutar este ejemplo, navega a la carpeta y ejecuta:
bash
cd interpkgs/context
go run main.go

💻 Tecnologías utilizadas
Go: El lenguaje de programación principal.
Net/http: Paquete estándar de Go para construir servidores web.
Encoding/json: Paquete estándar para codificar y decodificar JSON.
Context: Paquete estándar para gestionar el ciclo de vida de las peticiones. 
🤝 Contribuciones
Las contribuciones, los informes de errores y las sugerencias de mejora son bienvenidos. 
✍️ Autor
Tu Nombre - Tu perfil de GitHub