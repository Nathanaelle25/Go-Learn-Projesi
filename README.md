GoLearn API
Welcome to GoLearn API — a production-ready, high-performance Learning Management System (LMS) backend built with Go.

Author: NAME: NATHANAELLE 
Numbeer 24080410150

🚀 Features
Authentication & Security: JWT-based authentication with Bearer tokens, Role-Based Access Control (RBAC) ensuring Teachers and Students are separated securely.
RESTful Architecture: Full standard CRUD for Courses, Lessons, and Quizzes with automatic JSON structuring.
Quiz Engine: Built-in auto-grading functionality determining pass/fail thresholds (70%) based on dynamically matched correct answers.
Real-Time Websockets: Multi-room WebSocket endpoints acting as virtual "classrooms," automatically broadcasting "joined" and "left" states, with full mutex collision protection.
Database: Fully automated SQLite ORM execution via GORM with automatic schema migrations.
Rate-Limiting & CORS: Token bucket IP-based request throttling protecting the server from spam, plus configured CORS headers.
Swagger Documentation: Auto-generated, highly readable interactive API documentation available natively on the server.
🛠️ Technology Stack
Language: Go 1.22
Framework: Gin (github.com/gin-gonic/gin)
Database: SQLite (Pure-Go github.com/glebarez/sqlite)
ORM: GORM (gorm.io/gorm)
Security: golang.org/x/crypto/bcrypt & github.com/golang-jwt/jwt/v5
Rate Limiting: golang.org/x/time/rate
Live Classrooms: Gorilla WebSockets (github.com/gorilla/websocket)
API Docs: Swaggo (github.com/swaggo/gin-swagger)
🏗️ Structure Overview
golearn/
├── database/        # Database initialization & SQLite GORM configurations
├── docs/            # Auto-generated Swagger documentation mappings
├── handlers/        # API Controller Logic (Auth, Course, Lesson, Quiz, Progress, Websockets)
├── middleware/      # Interceptors (CORS, JWT Validations, RBAC, Rate Limiter)
├── models/          # Core LMS Entities (User, Course, Lesson, Quiz, Question, Progress)
├── Dockerfile       # Standardized multi-stage container deployment instructions
├── docker-compose   # Zero-Config environment simulation
├── e2e_test.go      # Concurrent end-to-end testing suite (22 full validation tests)
└── main.go          # Application entrypoint & Gin route assembly
🏃 Getting Started
1. Running Locally (Native)
Ensure you have Go installed, then start the server:

# Install dependencies
go mod tidy

# Run the project
go run .
The server will automatically start on http://localhost:8090.

2. Docker Execution
If you prefer running isolated instances without needing Go installed locally:

# Build and run the image detached
docker compose up --build -d
📖 API Documentation (Swagger)
Once the server is running, all REST endpoints can be viewed, simulated, and tested directly inside your browser: 👉 http://localhost:8090/swagger/index.html

🧪 Running Validation Tests
The project features an automated End-to-End suite consisting of 22 complex scenarios covering every aspect of the LMS (RBAC constraints, authentication handshakes, rate limit overloads, etc.).

To trigger the evaluations manually:

go test -v -count=1 .
