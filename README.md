# TaskWise Backend 🚀

TaskWise is a task management system with AI-powered task prioritization. This backend is built using Golang, Gin, PostgreSQL, JWT authentication, and AI integration via TensorFlow Lite.

## 📌 Features

- ✅ **User Authentication** – Register/Login with JWT authentication.
- ✅ **Task Management** – Create, Read, Update, Delete (CRUD) tasks.
- ✅ **Task Assignment** – Assign tasks to users.
- ✅ **Task Prioritization AI** – AI-based priority ranking.
- ✅ **Task Comments** – Users can add comments to tasks.
- ✅ **Secure API** – Uses JWT-based authentication and middleware.
- ✅ **Scalable Architecture** – Modular structure for easy expansion.

## 📂 Project Structure

```bash
TASKWISE-BACKEND/
│── ai/                       # AI integration (Golang-based AI logic)
│   ├── ai_handler.go         # Handles AI-related API calls (requests from ai_python)
│── ai_python/                # AI integration (Python-based ML model)
│   ├── __pycache__/          # Python cache
│   ├── ai_server.py          # FastAPI AI server (serves predictions)
│   ├── task_priority_model.h5 # AI model file (TensorFlow trained model)
│   ├── train_model.py        # AI model training script
│── cmd/                      # Main application entry (no more routes here)
│   ├── main.go               # Main application entry point
│── config/                   # Configuration files
│   ├── db.go                 # Database connection setup
│   ├── env.go                # Loads environment variables
│   ├── logger.go             # Logger configuration
│── controllers/              # API Controllers (Previously handlers)
│   ├── auth_controller.go    # Authentication (JWT)
│   ├── task_controller.go    # Task CRUD
│   ├── comment_controller.go # Task Comments
│   ├── ai_controller.go      # AI-related API handling
│── middleware/               # Middleware (Security & Auth)
│   ├── jwt.go                # JWT authentication middleware
│   ├── cors.go               # CORS middleware (if needed)
│   ├── logger.go             # Logging middleware (for request logging)
│── models/                   # Database models
│   ├── comment.go            # Comments model
│   ├── task.go               # Tasks model
│   ├── user.go               # Users model
│── repositories/             # Database queries (Separates DB logic)
│   ├── user_repository.go    # User queries
│   ├── task_repository.go    # Task queries
│   ├── comment_repository.go # Comment queries
│── routes/                   # API Routes (All routes are here)
│   ├── routes.go             # Central route loader
│   ├── auth_routes.go        # Authentication routes
│   ├── task_routes.go        # Task management routes
│   ├── comment_routes.go     # Comment-related routes
│   ├── ai_routes.go          # AI-related routes
│── services/                 # Business logic layer
│   ├── auth_service.go       # Authentication logic
│   ├── task_service.go       # Task-related business logic
│   ├── comment_service.go    # Comment handling logic
│   ├── ai_service.go         # AI integration logic
│── storage/                  # File storage (Optional, for AI models/uploads)
│   ├── models/               # Store AI models
│   ├── uploads/              # Store uploaded files (if any)
│── utils/                    # Utility functions
│   ├── hash.go               # Password hashing
│   ├── response.go           # API response formatting
│   ├── validator.go          # Input validation helpers
│── deploy/                   # Deployment scripts and configs
│   ├── docker-compose.yml    # For running both Go and Python services
│   ├── Dockerfile            # Container configuration
│── scripts/                  # Additional scripts (e.g., database migrations)
│── .env                      # Environment variables
│── .gitignore                # Git ignore list
│── go.mod                    # Go module file
│── go.sum                    # Go dependencies checksum
│── LICENSE                   # Project license
│── README.md                 # Documentation 
```
## ⚙️ Setup & Installation

### 1️⃣ Prerequisites

Make sure you have installed:

- ✅ Go 1.19+
- ✅ PostgreSQL 14+
- ✅ Python 3.10+ (for AI)

### 2️⃣ Clone the Repository

```sh
git clone https://github.com/azka-art/taskwise-backend.git
cd taskwise-backend
```

### 3️⃣ Create a .env File

Create a `.env` file in the root directory and add:

```ini
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=taskwise
DB_PORT=5433
DB_SSLMODE=disable
JWT_SECRET=your-secret-key
```
### 4️⃣ Install Dependencies

🔹 **Install Go dependencies**

```sh
go mod tidy
```
🔹 **Install Go dependencies**
```sh
cd ai_python
pip install -r requirements.txt
```
## 🚀 Running the Application

### 1️⃣ Start PostgreSQL Database

Ensure PostgreSQL is running on your system:

```sh
sudo systemctl start postgresql  # Linux
net start postgresql-x64-14      # Windows
```
Or use Docker:
```sh
docker-compose up -d
```
### 3️⃣ Run the Go Backend
```sh
go run cmd/main.go
```
Server should now be running on http://localhost:8080 🎉

## 🛠 API Endpoints

Here are the available API endpoints:

### 🔐 Authentication

| Method | Endpoint       | Description       |
|--------|----------------|-------------------|
| POST   | /api/register  | Register new user |
| POST   | /api/login     | Login user        |

### 📌 Task Management

| Method | Endpoint          | Description       |
|--------|-------------------|-------------------|
| POST   | /api/tasks        | Create a task     |
| GET    | /api/tasks        | Get all tasks     |
| PUT    | /api/tasks/:id    | Update a task     |
| DELETE | /api/tasks/:id    | Delete a task     |

### 💬 Comment Management

| Method | Endpoint                | Description       |
|--------|-------------------------|-------------------|
| POST   | /api/tasks/:id/comments | Add a comment     |
| GET    | /api/tasks/:id/comments | Get all comments  |

### 🤖 AI Integration

| Method | Endpoint                     | Description                         |
|--------|------------------------------|-------------------------------------|
| POST   | /api/tasks/recommendations   | Get AI-based task prioritization    |

## 📦 Deployment

### 🔹 Docker Deployment

Build and run the backend using Docker:

```sh
docker build -t taskwise-backend .
docker run -p 8080:8080 taskwise-backend
```
Use `docker-compose` to run Go + AI service:

```sh
docker-compose up --build
```
🛡 Security Features
--------------------

*   ✅ **JWT Authentication** for secure access
    
*   ✅ **Password hashing** with bcrypt
    
*   ✅ **CORS Middleware** for frontend integration
    
*   ✅ **Role-based access control** (planned for future updates)
    

📅 Roadmap
----------

*   Implement OAuth2 authentication
    
*   Add role-based permissions
    
*   Improve AI task prioritization model
    
*   Enhance logging with structured logging
    

📜 License
----------

This project is licensed under the MIT License.

🙌 Acknowledgments
------------------

*   Golang Gin Framework 🦫
    
*   PostgreSQL 🐘
    
*   TensorFlow AI 🤖
    
*   Docker & Kubernetes 🐳
