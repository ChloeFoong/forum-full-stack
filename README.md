# Forum Full Stack Application

A full-stack forum web application built with:

- **Backend:** Go (Gin, GORM, SQLite, JWT)
- **Frontend:** React
- **Containerization:** Docker & Docker Compose

This project supports user authentication, topics, posts, and protected routes using JWT.

---
## Features
- User registration & login
- JWT authentication
- Create and view forum topics and posts
- Protected API routes
- SQLite database
- Fully dockerized frontend and backend

---
## Prerequisites

Make sure you have installed:

- **Docker**
- **Docker Compose**

---
## How to Run the Application

### 1. Clone the repository

```bash
git clone https://github.com/ChloeFoong/forum-full-stack.git
cd forum-full-stack
```

### 2. Build and start application
```bash
docker compose up --build
```
This will:

- Build backend and frontend Docker images
- Start backend on http://localhost:8080
- Start frontend on http://localhost:3000

### 3. Access the application

Frontend (React app): http://localhost:3000
Backend API: http://localhost:8080

### 4. Stop the application

```bash
docker compose down
```
## Use of AI
- Research on Go libraries
- Research on databases (PostgreSQL, SQLite)
- Research on JWT authentication
- Review code on JWT authentication
- CSS syntax and styling
- Check if routing connection are correct
- Docker issues


