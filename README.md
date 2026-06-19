# Restaurant Mie API

RESTful API untuk sistem Point of Sale (POS) restoran berbasis Golang menggunakan Clean Architecture sederhana, Echo Framework, PostgreSQL, GORM, JWT Authentication, dan Railway Deployment.

## Features

### Authentication

* Register User
* Login User
* JWT Authentication
* Role Based Access Control (USER & KASIR)

### Menu Management

* Create Menu
* Get Menu List
* Get Menu Detail
* Update Menu
* Delete Menu

### Table Management

* Create Table
* Get Table List
* Get Table Detail
* Update Table
* Delete Table

### Order Management

* Create Order
* Get Order Detail
* Update Order
* Update Order Status

### Payment Management

* Create Payment
* Get Payment By Order
* Payment MVP Integration

---

## Environment Variables

Create `.env`

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=restaurant_db

JWT_SECRET=your-secret-key
```

---

## Installation

Clone Repository

```bash
git clone https://github.com/Akmalll1-123/restaurant-mie-api.git
```

Go to project

```bash
cd restaurant-mie-api
```

Install dependencies

```bash
go mod tidy
```

Run application

```bash
go run app/api/main.go
```

---

## API Base URL

Local

```text
http://localhost:8080/api/v1
```

Production

```text
https://restaurant-mie-api-production.up.railway.app
```

---

## Authentication Endpoints

### Register

```http
POST /api/v1/auth/register
```

Request

```json
{
  "name": "Akmal",
  "email": "akmal@mail.com",
  "password": "123456",
  "role": "USER"
}
```

### Login

```http
POST /api/v1/auth/login
```

Request

```json
{
  "email": "akmal@mail.com",
  "password": "123456"
}
```

Response

```json
{
  "token": "jwt-token"
}
```

---

## Menu Endpoints

### Create Menu

```http
POST /api/v1/menus
```

### Get Menus

```http
GET /api/v1/menus
```

### Get Menu Detail

```http
GET /api/v1/menus/:id
```

### Update Menu

```http
PUT /api/v1/menus/:id
```

### Delete Menu

```http
DELETE /api/v1/menus/:id
```

---

## Table Endpoints

### Create Table

```http
POST /api/v1/tables
```

### Get Tables

```http
GET /api/v1/tables
```

### Get Table Detail

```http
GET /api/v1/tables/:id
```

### Update Table

```http
PUT /api/v1/tables/:id
```

### Delete Table

```http
DELETE /api/v1/tables/:id
```

---

## Order Endpoints

### Create Order

```http
POST /api/v1/orders
```

### Get Order Detail

```http
GET /api/v1/orders/:id
```

### Update Order

```http
PUT /api/v1/orders/:id
```

### Update Order Status

```http
PATCH /api/v1/orders/:id/status
```

---

## Payment Endpoints

### Create Payment

```http
POST /api/v1/orders/:id/payment
```

### Get Payment By Order

```http
GET /api/v1/orders/:id/payment
```

---

## Database

Main Tables

* users
* menus
* tables
* orders
* order_items
* payments

---

## Roles

### USER

* Create Order
* View Order
* Create Payment

### KASIR

* Manage Menu
* Manage Table
* Update Order Status

---

## Author

Muhammad Akmal

GitHub:
https://github.com/Akmalll1-123
