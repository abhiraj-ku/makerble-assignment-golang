# 🏥 Health App Backend

A backend service for managing users, patients, authentication, and email processing in a healthcare application. Built using Go, PostgreSQL, Redis, and Gin.

---

## 🚀 Features

- **User Authentication** with JWT
- **Role-Based Access Control** (RBAC)
- **Patient Management** (CRUD)
- **Background Email Processing** using Redis
- **RESTful API** with Gin
- **Dockerized** with Docker and Docker Compose

---

## 🛠️ Tech Stack

- **Go** – Programming language
- **Gin** – Web framework
- **PostgreSQL** – Primary database
- **Redis** – Used for caching and background job queue
- **JWT** – For secure authentication
- **Docker + Compose** – Containerization and orchestration

---

## 📁 Project Structure

```
.
├── cmd/
│   └── health-app/
│       └── main.go       # Main entry point of the app
├── config/               # App configuration loader (reads from .env + config)
├── internal/
│   ├── auth/             # JWT auth, middleware, role-based access
│   ├── handler/          # HTTP handlers (Gin routes/controllers)
│   ├── repository/
│   │   └── db/
│   │       └── migration/ # SQL migration files
│   ├── service/          # Business logic (user, patient services)
│   └── worker/           # Background workers (email queue w/ Redis)
├── Dockerfile            # Docker build config
├── docker-compose.yml    # Runs health-app + Redis + PostgreSQL
├── Makefile              # Task runner for migrations, build, docker, etc.
├── go.mod
├── go.sum

```

---

## ⚙️ Configuration

Configuration is loaded from environment variables. Example variables:

```env
POSTGRES_URI=postgres://user:password@db:5432/health_app?sslmode=disable
REDIS_ADDR=redis:6379
SERVER_PORT=:8080
JWT_SECRET=your_secret_key
```

---

## 🐳 Running with Docker Compose

1. **Clone the repo**

   ```bash
   git clone https://github.com/yourusername/health_app.git
   cd health_app
   ```

2. **Start all services**
   ```bash
   docker-compose up --build
   ```

This will spin up the Go app, PostgreSQL, and Redis containers.

---

## 🧪 API Endpoints

### 🔐 Auth

- `POST /login` – User login, returns JWT

### 🩺 Patient (protected with RBAC)

- `GET /patients` – List all patients
- `POST /patients` – Create new patient
- `PUT /patients/:id` – Update patient
- `DELETE /patients/:id` – Delete patient

> Note: Routes are protected with `auth.JWTMiddleware()` and `auth.RequireRole`.

---

## 📬 Email Worker

The email worker listens for new email jobs in Redis and sends emails asynchronously.

Starts automatically:

```go
go emailWorker.ProcessEmailQueue()
```

---

## ▶️ Getting Started (without Docker)

1. **Set up PostgreSQL and Redis manually**

2. **Run the server**
   ```bash
   go run main.go
   ```

---

## 🧹 Future Improvements

- Add unit and integration tests

---

## 🧑‍💻 Author

Developed by [Abhiraj Ku](https://github.com/abhiraj-ku)

---

## 📄 License

MIT License
