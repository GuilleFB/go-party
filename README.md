# 🚀 CRUD de Usuarios y Tareas en GO 🎉

¡Bienvenido/a a mi proyecto de aprendizaje para un CRUD de Usuarios y Tareas! 🤩✨ Este proyecto está desarrollado en el lenguaje de programación GO. 🐹🚀

## 🛠️ Tecnologías Utilizadas

- [**Go**](https://go.dev/doc/install) 🐹
- [**Gorilla Mux**](https://github.com/gorilla/mux) 🦍 
- [**Gin**](https://gin-gonic.com/docs/) 🍸
- [**GORM**](https://gorm.io/docs/) 🗃️
- [**JWT (JSON Web Tokens)**](https://golang-jwt.github.io/jwt/#supported-go-versions) 🔐
- **Go Crypto** 🛡️
- [**Air**](https://github.com/cosmtrek/air) ☁️ Live reload for Go apps
- **PostgreSQL** 🐘
- **Docker** 🐳

## 🌟 Características
- ✅ Signup, Login y Perfil de usuario con middleware que comprueba Autenticación :accessibility:
- ✅ Crear, leer, actualizar y eliminar usuarios 👤
- ✅ Crear, leer, actualizar y eliminar tareas 📋
- ✅ Asociación entre usuarios y tareas 🔗

## 🚀 Empezando

Sigue estos sencillos pasos para poner en marcha el proyecto en tu entorno local 🏡:

### 1. Clonar el repositorio 📂

```bash
git clone https://github.com/GuilleFB/go-party.git
cd go-party
```

### 2. Configurar las variables de entorno 🛠️

Crea un archivo `.env` en el directorio raíz del proyecto y añade tus configuraciones:

```bash
cp .env.template .env
```

```env
DB_HOST=localhost
DB_PORT=tu_port
POSTGRES_USER=tu_usuario
POSTGRES_PASSWORD=tu_contraseña
POSTGRES_DB=tu_basedatos
SECRET=tu_secreto
```

### 3. Construir y ejecutar el contenedor Docker 🐳

```bash
docker compose up -d
```
Con esto levantaremos la base de datos en PostgreSQL.

### 4. Inicializar Air

```bash
go install github.com/cosmtrek/air@latest
```
```bash
air
```

### 4. Acceder a la API 🚀

Una vez que el contenedor esté en funcionamiento, puedes acceder a la API en `http://localhost:8000`.

## 🧩 Endpoints de la API


### Usuarios 👤

- **Signup**: `POST gin/auth/signup`
- **Login**: `POST /gin/auth/login`
- **Profile**: `GET /gin/user/profile`

### Usuarios 👤

- **Crear usuario**: `POST /mux/user/create`
- **Obtener todos los usuarios**: `GET /mux/users`
- **Obtener usuario por ID**: `GET /mux/user/{id}`
- **Actualizar usuario**: `PATCH /mux/user/edit/{id}`
- **Eliminar usuario**: `DELETE /mux/user/delete/{id}`

### Tareas 📋

- **Crear tarea**: `POST /gin/tasks`
- **Obtener todas las tareas**: `GET /gin/tasks`
- **Obtener tarea por ID**: `GET /gin/task/:id`
- **Actualizar tarea**: `PATCH /gin/tasks/edit/:id`
- **Eliminar tarea**: `DELETE /gin/tasks/delete/:id`


## 🤝 Contribuciones

¡Las contribuciones son bienvenidas! 🙌 Si tienes alguna idea o encuentras algún error, por favor abre un issue o envía un pull request. 🛠️🔧

## 📬 Contacto

Si tienes alguna pregunta o sugerencia, no dudes en comentar algo en las [discusiones](https://github.com/GuilleFB/go-party/discussions).

---

¡Gracias por pasarte por aquí y feliz programación! 💻🎉🚀
