# Go TODO API

A simple RESTful TODO API built in Go.

## Features

- Create, read, and delete TODO items
- Stores data in-memory
- Dockerized for easy setup
- Swagger (OpenAPI) documentation included
- Unit tests and CI pipeline via GitHub Actions

## Getting Started

### Prerequisites

- Docker installed

### Run with Docker

```bash
docker build -t go-todo-api .
docker run -p 8080:8080 go-todo-api
