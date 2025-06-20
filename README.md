# 🛍️ Inventory Management API (Go + MySQL)

A lightweight RESTful API for managing product inventory, built with **Go**, **MySQL**, and **Gorilla Mux**. It supports standard CRUD operations and includes integration tests using Go's `net/http/httptest`.

---

## 📦 Features

* Create, retrieve, update, and delete products.
* JSON-based API communication.
* Modular structure (handlers, model, router, tests).
* Integration-tested endpoints.
* MySQL-based persistent storage.

---

## 🧱 Project Structure

```bash
.
├── app.go           # Application core (routing and handler wiring)
├── app_test.go      # Integration tests for API endpoints
├── model.go         # Database model and logic
├── main.go          # Entry point of the app
├── constants.go     # Database configuration constants
└── README.md        # Project documentation
```

---

## 🔧 Requirements

* Go 1.24.2
* MySQL 8.0
* Git

---

## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/inventory-api.git
cd inventory-api
```

### 2. Configure MySQL

Create a MySQL database called `inventory` and user credentials matching `constants.go`:

```sql
CREATE DATABASE inventory;
CREATE USER 'root'@'localhost' IDENTIFIED BY 'root';
GRANT ALL PRIVILEGES ON inventory.* TO 'root'@'localhost';
```

> You can modify `DBUser`, `DBPassword`, and `DBName` in `constants.go` if needed.

### 3. Run the Application

```bash
go run main.go
```

Server will start at `http://localhost:10000`

---

## 🧪 Running Tests

To run tests, ensure you have a `test` database created:

```sql
CREATE DATABASE test;
```

Then run:

```bash
go test
```

---

## 📘 API Endpoints

### 🔍 GET All Products

```http
GET /products
```

### 🔍 GET Single Product

```http
GET /product/{id}
```

### ➕ Create Product

```http
POST /product
Content-Type: application/json

{
  "name": "Laptop",
  "quantity": 10,
  "price": 999.99
}
```

### 🔁 Update Product

```http
PUT /product/{id}
Content-Type: application/json

{
  "name": "Updated Laptop",
  "quantity": 8,
  "price": 899.99
}
```

### ❌ Delete Product

```http
DELETE /product/{id}
```