# SimpleBank

A production-style backend API built with Go for learning and demonstrating core backend engineering concepts such as authentication, database access, account workflows, Dockerized development, and clean project structure.

This project is designed to showcase backend fundamentals with a Go-first architecture. It includes API handlers, database access layers, token-based authentication, utility packages, Docker setup, SQLC configuration, and deployment-oriented project organization.

## Why This Project Matters

SimpleBank demonstrates the kind of backend skills required for real product systems:

- designing REST APIs
- structuring a Go backend project
- managing database queries
- handling authentication
- using environment-based configuration
- running services with Docker
- organizing backend logic into clear packages
- preparing a project for CI/CD and deployment workflows

For LinkedIn and recruiter review, this project is strong proof of hands-on Golang backend learning and production-style API structure.

## Tech Stack

- Language: Go
- HTTP Framework: Gin
- Database: PostgreSQL
- Query Generation: SQLC
- Authentication: JWT and PASETO
- Configuration: Viper
- Containerization: Docker and Docker Compose
- Testing: Go testing tools and Testify
- CI/CD: GitHub Actions workflow folder included
- Deployment Structure: EKS folder included

## Core Features

### Account Management

- Create bank accounts
- Fetch account details
- List accounts with pagination-style backend patterns
- Update account data
- Delete account records where applicable

### Transfers

- Create money transfer records
- Track source and destination accounts
- Handle database-backed transfer workflows
- Support transactional backend thinking

### Authentication

- User login flow
- Token-based authentication
- JWT support
- PASETO support
- Secure password hashing using crypto packages

### Database Layer

- SQLC-based query generation
- Typed database access
- PostgreSQL driver integration
- Organized database package structure

### Backend Architecture

- Clean separation between API, database, token, and utility packages
- Environment-based configuration
- Dockerized local development
- Makefile commands for developer productivity

## Project Structure

```text
simplebank/
├── .github/workflows/
├── api/
├── db/
├── doc/
├── eks/
├── token/
├── util/
├── Dockerfile
├── Makefile
├── app.env
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
├── sqlc.yaml
├── start.sh
└── wait-for.sh
```

## Main Concepts Demonstrated

- REST API development in Go
- Gin router and HTTP handlers
- PostgreSQL-backed persistence
- SQLC-generated type-safe queries
- Authentication using JWT and PASETO
- Docker and Docker Compose setup
- Environment configuration with Viper
- Modular backend package design
- Backend testing foundations
- Deployment-ready project organization

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/MustafaKheda/simplebank.git
cd simplebank
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Configure Environment Variables

Update the `app.env` file with your local database and server settings.

Example:

```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
ACCESS_TOKEN_DURATION=15m
```

### 4. Start Services with Docker Compose

```bash
docker-compose up
```

### 5. Run the Application

```bash
go run main.go
```

## Useful Commands

Run tests:

```bash
go test ./...
```

Build Docker image:

```bash
docker build -t simplebank .
```

Start using Makefile commands:

```bash
make
```

## LinkedIn Project Summary

Built a production-style Go backend API with PostgreSQL, Gin, SQLC, Docker, token-based authentication, and modular backend architecture. This project demonstrates my growing Golang backend skills, API design fundamentals, database workflow understanding, and ability to structure backend services for maintainability.

## Featured Section Title

**SimpleBank | Golang Backend API with Auth, PostgreSQL & Docker**

## Featured Section Description

Built a Go backend API with Gin, PostgreSQL, SQLC, Docker, JWT/PASETO authentication, and modular service structure. Demonstrates backend API design, database access, token-based auth, and production-style project organization.

## Skills Demonstrated

- Golang
- Backend API Development
- REST APIs
- Gin
- PostgreSQL
- SQLC
- JWT
- PASETO
- Docker
- Docker Compose
- Authentication
- Database Design
- Backend Architecture
- Clean Project Structure
- Developer Productivity

## Future Improvements

- Add complete API documentation
- Add Swagger/OpenAPI docs
- Add request and response examples
- Add architecture diagram
- Add database schema diagram
- Add CI/CD status badge
- Add deployment guide
- Add screenshots of API testing through Postman

## Author

Mustafa Kheda

Backend-focused full stack developer building scalable APIs, backend systems, and performance-first product infrastructure.
