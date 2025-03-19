# 🚀 Todo API (Golang + PostgreSQL)

This is a **Todo List API** built with **Golang**, using **Gin**, **GORM**, and **PostgreSQL**.  
It supports **CRUD operations** for managing tasks and users and is **containerized with Docker**.

---

## 📌 **1. Tech Stack**
- **Golang** (Gin Framework) - Web API framework
- **PostgreSQL** - Database
- **GORM** - ORM for PostgreSQL
- **Docker & Docker Compose** - Containerization
- **godotenv** - Environment variable management
- **JWT** - Authentication and Authorization

---

## 📌 **2. Modules**
### **Task Module**
- Manage tasks with the following endpoints:
  - Create a task
  - Retrieve all tasks
  - Retrieve a task by ID
  - Update a task
  - Mark a task as completed
  - Delete a task

### **User Module**
- Manage users with the following endpoints:
  - Create a user
  - Retrieve all users
  - Retrieve a user by ID

### **Authentication Module**
- Manage authentication with the following endpoints:
  - Login to get a JWT token

---

## 📌 **3. Running the app**

### **1. Initialize the project**
Before running the application, ensure that the Go module is initialized and all dependencies are installed.

#### **Step 1: Initialize `go.mod`**
Run the following command in the root directory of the project to initialize the Go module:
```bash
go mod init todo-api
```

#### **Step 2: Install dependencies**
Run the following command to install all required libraries:
```bash
go get ./...
```

---

### **2. Run the server**
```bash
# Start the server
$ go run cmd/server/main.go
```

---

### **3. Manual database operations**
You can manually manage the database using the following commands:

#### **Setup the database**
Run setup to create the necessary tables in the database:
```bash
$ go run cmd/server/main.go setup
```

#### **Seed the database**
Insert default data into the `tasks` and `users` tables:
```bash
$ go run cmd/server/main.go seed
```

#### **Drop the database**
Drop all tables (`tasks` and `users`) from the database:
```bash
$ go run cmd/server/main.go drop
```

---

## 📌 **4. API Endpoints**

### **Authentication Module**
| Method | Endpoint   | Description               |
|--------|------------|---------------------------|
| POST   | `/login`   | Login to get a JWT token  |

#### **Example Request**
```json
{
  "email": "admin@example.com",
  "password": "password"
}
```

#### **Example Response**
```json
{
  "token": "your_jwt_token"
}
```

---

### **Task Module**
| Method | Endpoint               | Description               |
|--------|-------------------------|---------------------------|
| POST   | `/tasks`               | Create a new task         |
| GET    | `/tasks`               | Retrieve all tasks        |
| GET    | `/tasks/:id`           | Retrieve a task by ID     |
| PUT    | `/tasks/:id`           | Update an existing task   |
| PATCH  | `/tasks/:id/completed` | Mark a task as completed  |
| DELETE | `/tasks/:id`           | Delete a task             |

### **User Module**
| Method | Endpoint       | Description           |
|--------|-----------------|-----------------------|
| POST   | `/users`       | Create a new user     |
| GET    | `/users`       | Retrieve all users    |
| GET    | `/users/:id`   | Retrieve a user by ID |

---

## 📌 **5. Environment Variables**
The application uses environment variables to configure the database connection. Create a `.env` file in the root directory with the following content:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database
JWT_SECRET=your_secret_key
```

---

## 📌 **6. Docker Support**
You can run the application using Docker and Docker Compose.

### **Build and run the app**
```bash
# Build the Docker image
$ docker-compose build

# Start the containers
$ docker-compose up
```

### **Stop the app**
```bash
$ docker-compose down
```

---

## 📌 **7. Future Improvements**
- Add role-based authorization
- Implement pagination for large datasets
- Add more modules (e.g., Projects, Comments)
- Write unit tests for all modules
