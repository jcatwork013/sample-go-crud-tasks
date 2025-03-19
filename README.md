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

## 📌 **4. Building for Production**

### **1. Build the application**
To build the application for production, use the following command:
```bash
go build -o dist/todo-api cmd/server/main.go
```

This will create a binary file in the `dist/` directory.

### **2. Build for specific environments**
You can build the application for different operating systems and architectures using the `GOOS` and `GOARCH` environment variables.

#### **Build for Linux**
```bash
GOOS=linux GOARCH=amd64 go build -o dist/todo-api-linux cmd/server/main.go
```

#### **Build for Windows**
```bash
GOOS=windows GOARCH=amd64 go build -o dist/todo-api.exe cmd/server/main.go
```

#### **Build for macOS**
```bash
GOOS=darwin GOARCH=amd64 go build -o dist/todo-api-macos cmd/server/main.go
```

---

### **3. Running the built application**
After building the application, you can run it directly from the `dist/` directory.

#### **Run the application**
```bash
cd dist
./todo-api
```

#### **Ensure `.env` is present**
Make sure to copy the `.env` file to the `dist/` directory before running the application:
```bash
cp .env dist/
cd dist
./todo-api
```

---

## 📌 **5. Deploying with Docker**

### **1. Build Docker image**
To deploy the application using Docker, build the Docker image:
```bash
docker build -t todo-api .
```

### **2. Run the Docker container**
Run the container with the following command:
```bash
docker run -d -p 8080:8080 --name todo-api --env-file .env todo-api
```

- **`-p 8080:8080`**: Maps port 8080 of the container to port 8080 on the host machine.
- **`--env-file .env`**: Loads environment variables from the `.env` file.

### **3. Push Docker image to a registry**
If you want to deploy the application to a remote server, push the Docker image to a registry like Docker Hub:
```bash
docker tag todo-api your-dockerhub-username/todo-api
docker push your-dockerhub-username/todo-api
```

### **4. Pull and run the image on the server**
On the production server, pull the image and run the container:
```bash
docker pull your-dockerhub-username/todo-api
docker run -d -p 8080:8080 --env-file .env your-dockerhub-username/todo-api
```

---

## 📌 **6. Environment Variables**
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

## 📌 **7. Future Improvements**
- Add role-based authorization
- Implement pagination for large datasets
- Add more modules (e.g., Projects, Comments)
- Write unit tests for all modules
