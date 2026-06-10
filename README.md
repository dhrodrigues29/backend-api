# Backend API

[![CI](https://github.com/dhrodrigues29/backend-api/actions/workflows/ci.yml/badge.svg)](https://github.com/dhrodrigues29/backend-api/actions)

REST API built with Go, PostgreSQL and Docker, designed as a portfolio project to demonstrate backend development, database design, authentication and CI/CD practices.

## Overview

This project aims to showcase modern backend engineering practices using Go and PostgreSQL. It includes authentication, database integration, containerized deployment and automated quality checks through GitHub Actions.

## Tech Stack

- Go
- PostgreSQL
- Docker
- JWT Authentication
- GitHub Actions
- Git
- REST API

## Features

### Current

- REST API structure
- PostgreSQL integration
- Dockerized environment
- Automated CI workflow

### Planned

- User authentication with JWT
- User registration and login
- Role-based authorization
- CRUD operations
- Database migrations
- API documentation
- Unit and integration tests

## Project Structure

```text
backend-api/
├── cmd/
├── internal/
├── pkg/
├── configs/
├── migrations/
├── tests/
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.23+
- Docker
- PostgreSQL

### Clone Repository

```bash
git clone git@github.com:dhrodrigues29/backend-api.git
cd backend-api
```

### Install Dependencies

```bash
go mod download
```

### Run Application

```bash
go run ./cmd/api
```

### Run Tests

```bash
go test ./...
```

### Build

```bash
go build ./...
```

## Continuous Integration

This repository uses GitHub Actions to automatically:

- Build the application
- Run tests
- Validate code quality

Every push and pull request triggers the CI pipeline.

## Roadmap

- [ ] JWT Authentication
- [ ] User CRUD
- [ ] PostgreSQL persistence
- [ ] Docker Compose environment
- [ ] Swagger/OpenAPI documentation
- [ ] Integration tests
- [ ] Deployment pipeline

## Author

Davi Haas Rodrigues

- LinkedIn: linkedin.com/in/davi-haas-rodrigues
- GitHub: github.com/dhrodrigues29
