# 🐛 PatchPal - Bug Logging REST API

PatchPal is a simple and extensible bug logging REST API built using **Go**, **Gin**, and **GORM**. It allows you to log, retrieve, update, and delete software bugs. Ideal for teams looking for a minimal internal issue tracker or developers practicing Go-based backend development.

---

## 🚀 Features

- Log new bugs with title and description
- List all reported bugs
- Retrieve a specific bug by ID
- Update an existing bug
- Delete a bug
- Powered by SQLite (easy to swap with PostgreSQL/MySQL)

---

## 🧱 Tech Stack

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) — HTTP Web Framework
- [GORM](https://gorm.io/) — ORM for Go
- [SQLite](https://www.sqlite.org/index.html) — Lightweight DB

---

## ⚙️ Getting Started

### 1. Clone the Repo

```bash
git clone https://github.com/Prakhar-Shankar/patch-pal.git
cd patch-pal
```
Install Dependencies
`go mod tidy`

Run the server
`go run main.go`

---

Create a bug
`curl -X POST http://localhost:8080/bugs \
  -H "Content-Type: application/json" \
  -d '{"title": "Login error", "description": "Button not working", "status": "open"}'
`


Get all bugs
`curl http://localhost:8080/bugs
`



Get bug by ID
`curl http://localhost:8080/bugs/1
`

Update a bug 
`curl -X PATCH http://localhost:8080/bugs/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Login issue updated", "status": "resolved"}'
`



Delete a bug 
`curl -X DELETE http://localhost:8080/bugs/1
`
