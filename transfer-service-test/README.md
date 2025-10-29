# Transfer Service Test

# 🚚 Transfer Service Test (Full Stack Technical Assessment)

This project implements a simplified **Transfer Service System** as part of a technical assessment. It showcases **clean architecture**, **backend reliability** (idempotency + outbox pattern), and a **React + TypeScript frontend** with live polling.

## 🧱 Project Structure


## ⚙️ Tech Stack

| Layer | Technology |
|-------|-------------|
| Backend | Go 1.21+, chi router, PostgreSQL |
| Frontend | React + TypeScript + Vite |
| Database | PostgreSQL |
| Reliability | Outbox pattern + Idempotency |
| Deployment | Docker Compose |

## 🚀 How to Run

### 1️⃣ Prerequisites
- Docker & Docker Compose installed
- Node.js 18+ (optional for local frontend dev)

### 2️⃣ Start Everything
Run from the root folder:
```bash
docker-compose up --build

curl -X POST http://localhost:8080/transfers \
  -H "Content-Type: application/json" \
  -H "Idempotency-Key: abc123" \
  -d '{
    "pallet_id": "PALLET-01",
    "from_location": "LOC-A",
    "to_location": "LOC-B",
    "note": "Urgent move"
  }'


services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: transferdb
    ports:
      - "5432:5432"
    volumes:
      - ./data:/var/lib/postgresql/data

  backend:
    build: ./backend/transfer-service
    environment:
      DATABASE_URL: postgres://postgres:postgres@postgres:5432/transferdb?sslmode=disable
      ENABLE_CAPACITY_CHECK: "true"
    ports:
      - "8080:8080"
    depends_on:
      - postgres

  frontend:
    build: ./frontend
    ports:
      - "5173:5173"


cd backend/transfer-service
go run main.go


cd frontend
npm install
npm run dev

Gustaf (Technical Test Submission)
Built with using Go, React, and PostgreSQL.