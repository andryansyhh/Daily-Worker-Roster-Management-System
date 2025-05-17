# 🧑‍💼 Daily Worker Roster Management System

## 📌 Project Overview

This is a RESTful API system for managing daily shift rosters for workers. It supports worker shift signup, admin approval/rejection, and assignment viewing. Built with **Golang** and **SQLite**, and packaged using **Docker Compose** for easy deployment.

---

## 🚀 Setup Instructions (Using Docker Compose)

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/daily-worker-roster.git
cd daily-worker-roster
```

### 2. Build and Run

```bash
docker compose up --build
```

The server will run at: [http://localhost:8089](http://localhost:8089)

> Make sure port 8089 is available.

---

## 🧠 Documented Assumptions

* Workers cannot request shifts already assigned to someone else.
* No overlapping shift requests allowed per worker.
* A worker can only have **1 shift per day** and **max 5 shifts per week**.
* Admins can override or reassign approved shifts.
* Conflict checks happen both at **request** and **approval** stages.
* Shift time is stored and compared in **UTC**.

---

## 📡 API Usage Guide

### 🧍‍♂️ Worker Endpoints

#### 🔹 Register Worker

```http
POST /worker/register
```

**Body:**

```json
{
  "name": "Andi",
  "email": "andi@example.com"
}
```

#### 🔹 Login (Admin or Worker)

```http
POST /auth/login
```

**Body:**

```json
{
  "user_id": 1,
  "user_type": "worker" | "admin"
}
```

**Returns:** JWT token

#### 🔹 View Available Shifts (Worker)

```http
GET /worker/shifts/available
```

#### 🔹 Request a Shift

```http
POST /worker/shifts/:id/request
```

#### 🔹 View My Requests

```http
GET /worker/requests
```

#### 🔹 View My Assignments

```http
GET /worker/assignments
```

---

### 🧑‍💼 Admin Endpoints

(Require JWT token + user\_type = admin)

#### 🔹 Create Shift

```http
POST /admin/shifts
```

#### 🔹 View All Shifts

```http
GET /admin/shifts
```

#### 🔹 Get Shift By ID

```http
GET /admin/shifts/:id
```

#### 🔹 Update Shift

```http
PUT /admin/shifts/:id
```

#### 🔹 Delete Shift

```http
DELETE /admin/shifts/:id
```

#### 🔹 Approve Request

```http
PUT /admin/requests/:id/approve
```

#### 🔹 Reject Request

```http
PUT /admin/requests/:id/reject
```

#### 🔹 View Assignments by Date

```http
GET /admin/assignments
```

#### 🔹 View Assignments by Worker

```http
GET /admin/assignments/worker/:id
```

#### 🔹 Reassign Shift

```http
PUT /admin/assignments/:shift_id/reassign
```

---

## 🧪 Postman Collection

Use Postman to test all endpoints:

1. Register: `POST /worker/register`
2. Login: `POST /auth/login`
3. Set token in headers: `Authorization: Bearer <token>`
4. Test `/worker/...` or `/admin/...` endpoints based on your role

---

## 🔐 .env.example

```env
PORT=8089
JWT_SECRET=supersecret
```

---

## ⚙️ GitHub Actions (CI/CD)

Create file: `.github/workflows/build.yml`

```yaml
name: Go Build & Test

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout repo
      uses: actions/checkout@v3

    - name: Setup Go
      uses: actions/setup-go@v4
      with:
        go-version: 1.21

    - name: Build
      run: go build -v ./...

    - name: Run Tests
      run: go test ./...
```

---

## 🗂 Folder Structure

```
.
├── cmd/               # Entry point
├── internal/          # Modules: handler, usecase, repository, model
├── migration/         # SQL schema (optional)
├── main.go            # Starts app via cmd.Run()
├── Dockerfile
├── docker-compose.yml
├── roster.db          # SQLite DB (generated)
```
