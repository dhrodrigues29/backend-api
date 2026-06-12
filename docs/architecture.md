# Backend API Architecture

## Goal

Build a REST API for task and project management, focused on authentication, clean architecture, PostgreSQL persistence, Docker and automated CI.

## Core entities

- User
- Project
- Task

## Main features

- User registration
- User login
- JWT authentication
- Create projects
- List projects
- Create tasks
- Update task status
- Assign tasks to users
- Delete tasks

## Non-goals

- Frontend
- Real-time chat
- Payments
- Complex notifications

## Project Structure

backend-api/
├── cmd/
│ └── api/
├── internal/
│ ├── auth/
│ ├── database/
│ ├── project/
│ ├── server/
│ ├── task/
│ └── user/
├── migrations/
├── tests/
└── docs/

### cmd/api

Application entrypoint.

Contains main.go responsible for starting the HTTP server.

### internal/server

HTTP server configuration and route registration.

### internal/database

Database connection, configuration and persistence logic.

### internal/auth

Authentication, JWT generation and authorization middleware.

### internal/user

User business logic and user-related endpoints.

### internal/project

Project business logic and project-related endpoints.

### internal/task

Task business logic and task-related endpoints.

### migrations

Database schema migrations.

### tests

Integration and automated tests.

### docs

Project documentation and architecture decisions.
