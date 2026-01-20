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

From the project root:

```bash
docker compose up --build
