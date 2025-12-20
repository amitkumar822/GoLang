# User Management System - Backend API

A production-ready RESTful API built with Go (Golang) and MongoDB for managing users with authentication, authorization, and CRUD operations.

## 🚀 Quick Start

**Prerequisites:** Go 1.21+, MongoDB running

```bash
# 1. Install dependencies
go mod download

# 2. Create .env file (copy from .env.example)
# Windows PowerShell:
Copy-Item .env.example .env

# Linux/Mac:
# cp .env.example .env

# 3. Start the application
go run cmd/main.go

# OR for auto-reload during development (recommended):
# Install air: go install github.com/cosmtrek/air@latest
# Then run: air
```

**That's it!** The server will start on `http://localhost:8080`

> 📝 **Note:** Make sure MongoDB is running before starting the application. See [Installation & Setup](#installation--setup) for detailed instructions.
> 
> 💡 **Tip:** Use `air` for development - it automatically restarts the server when you make code changes (like `nodemon` in Node.js)!

## 📋 Table of Contents

- [Quick Start](#-quick-start)
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Prerequisites](#-prerequisites)
- [Installation & Setup](#-installation--setup)
- [Configuration](#-configuration)
- [Running the Application](#-running-the-application)
  - [Development Mode (Standard)](#development-mode-standard)
  - [Development Mode (with Auto-Reload)](#development-mode-with-auto-reload-)
  - [Production Build](#production-build)
- [API Documentation](#-api-documentation)
- [Authentication Flow](#-authentication-flow)
- [Error Handling](#-error-handling)
- [Testing Examples](#-testing-examples)

## ✨ Features

- ✅ User Registration & Login
- ✅ JWT-based Authentication
- ✅ Role-based Authorization (User/Admin)
- ✅ Password Hashing with bcrypt
- ✅ CRUD Operations for Users
- ✅ Pagination Support
- ✅ Bulk User Retrieval (All Users) with Concurrent Processing
- ✅ Input Validation
- ✅ CORS Support
- ✅ Request Logging
- ✅ Panic Recovery
- ✅ Clean Architecture

## 🛠 Tech Stack

- **Language**: Go 1.21+
- **Database**: MongoDB
- **Router**: Gorilla Mux
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Config Management**: .env files

## 📁 Project Structure

```
01_First/
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── config.go               # Configuration management
├── database/
│   └── mongo.go                # MongoDB connection
├── models/
│   └── user.go                 # User models and DTOs
├── repositories/
│   └── user_repository.go      # Database operations
├── services/
│   └── user_service.go         # Business logic
├── handlers/
│   ├── auth_handler.go         # Authentication handlers
│   └── user_handler.go         # User CRUD handlers
├── middleware/
│   ├── jwt_middleware.go       # JWT validation
│   ├── auth_middleware.go      # Authorization middleware
│   ├── cors_middleware.go      # CORS handling
│   ├── logging_middleware.go   # Request logging
│   └── recovery_middleware.go  # Panic recovery
├── routes/
│   └── routes.go               # Route configuration
├── utils/
│   ├── jwt.go                  # JWT utilities
│   └── response.go             # Response helpers
├── .env.example                # Environment variables template
├── go.mod                      # Go dependencies
└── README.md                   # This file
```

## 📦 Prerequisites

Before running this application, ensure you have:

1. **Go** (version 1.21 or higher)
   ```bash
   go version
   ```

2. **MongoDB** (running locally or remote instance)
   - Download from [MongoDB Official Website](https://www.mongodb.com/try/download/community)
   - Or use MongoDB Atlas (cloud)

3. **Postman** or **curl** (for testing API endpoints)

4. **Air** (optional but recommended for development)
   - Auto-reload tool for Go applications
   - Install: `go install github.com/cosmtrek/air@latest`
   - See [Development Mode with Auto-Reload](#development-mode-with-auto-reload-) for details

## 🚀 Installation & Setup

### Step 1: Clone/Navigate to Project

```bash
cd 01_First
```

### Step 2: Install Dependencies

```bash
go mod download
```

Or simply run (it will auto-download):

```bash
go mod tidy
```

### Step 3: Configure Environment Variables

Create a `.env` file from the example:

```bash
# Windows PowerShell
Copy-Item .env.example .env

# Linux/Mac
cp .env.example .env
```

Edit `.env` file with your configuration:

```env
APP_PORT=8080
MONGO_URI=mongodb://localhost:27017
MONGO_DB=userdb
JWT_SECRET=your-super-secret-key-change-this-in-production
JWT_EXPIRE_HOURS=24
```

**Important**: Change `JWT_SECRET` to a secure random string in production!

### Step 4: Start MongoDB

Make sure MongoDB is running:

```bash
# Windows (if installed as service, it should auto-start)
# Or start manually:
mongod

# Linux/Mac
sudo systemctl start mongod
# or
mongod
```

## ⚙️ Configuration

The application reads configuration from environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | Server port number | `8080` |
| `MONGO_URI` | MongoDB connection string | `mongodb://localhost:27017` |
| `MONGO_DB` | Database name | `userdb` |
| `JWT_SECRET` | Secret key for JWT signing | `supersecretkey` |
| `JWT_EXPIRE_HOURS` | JWT token expiration (hours) | `24` |

## 🏃 Running the Application

### Development Mode (Standard)

```bash
go run cmd/main.go
```

### Development Mode (with Auto-Reload) ⚡

For automatic server restart on code changes (recommended for development):

**Step 1: Install Air**
```bash
# Install air globally
go install github.com/cosmtrek/air@latest

# Verify installation
air -v
```

**Step 2: Run with Air**
```bash
# From the project root directory
air
```

**How it works:**
- ✅ Watches all `.go` files for changes
- ✅ Automatically rebuilds when files are saved
- ✅ Restarts the server automatically
- ✅ Shows build errors in real-time
- ✅ No need to manually stop/start the server

**Configuration:**
The project includes a `.air.toml` configuration file that's already set up. Air will:
- Watch all Go files in the project
- Rebuild when any `.go` file changes
- Exclude test files and temporary directories
- Display colored output for better visibility
- Create a `tmp/` directory for build artifacts (already in `.gitignore`)

**Note:** The `tmp/` directory and `build-errors.log` file are automatically created by air and are already excluded in `.gitignore`.

**Example workflow:**
```bash
# Terminal 1: Start the server with air
air

# Terminal 2: Make changes to any .go file
# Save the file → Server automatically restarts!
```

**Alternative: Quick Air Command (without config file)**
```bash
air -c . -build.cmd "go build -o ./tmp/main ./cmd/main.go" -build.bin "./tmp/main"
```

### Production Build

```bash
# Build the application
go build -o user-management-api cmd/main.go

# Run the executable
./user-management-api    # Linux/Mac
user-management-api.exe  # Windows
```

### Expected Output

```
✅ MongoDB connected successfully
🚀 Server starting on port 8080
🌐 Root endpoint: http://localhost:8080/
📝 API endpoints available at http://localhost:8080/api
```

The server will start on `http://localhost:8080` (or your configured port).

## 📚 API Documentation

### Base URL

```
http://localhost:8080/api
```

### Authentication

Most endpoints require JWT authentication. Include the token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

---

## 🏠 Root Endpoint

### Welcome Message

Get API information and welcome message.

**Endpoint:** `GET /`

**Authentication:** Not required (Public)

**Success Response (200 OK):**
```json
{
  "success": true,
  "message": "API is running",
  "data": {
    "message": "Welcome to Go Lang User API",
    "version": "1.0.0",
    "status": "running",
    "endpoints": {
      "register": "/api/auth/register",
      "login": "/api/auth/login",
      "users": "/api/users",
      "usersAll": "/api/users/all"
    }
  }
}
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/
```

**Browser:** Simply open `http://localhost:8080/` in your browser!

---

## 🔐 Authentication Endpoints

### 1. Register User

Register a new user account.

**Endpoint:** `POST /api/auth/register`

**Authentication:** Not required (Public)

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Validation Rules:**
- `name`: Required, minimum 2 characters
- `email`: Required, valid email format
- `password`: Required, minimum 6 characters

**Success Response (200 OK):**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "id": "65ab1234567890abcdef1234",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "isActive": true,
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-01-15T10:30:00Z"
  }
}
```

**Error Response (400 Bad Request):**
```json
{
  "success": false,
  "error": "email already registered"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

---

### 2. Login User

Authenticate user and receive JWT token.

**Endpoint:** `POST /api/auth/login`

**Authentication:** Not required (Public)

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Success Response (200 OK):**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "user": {
      "id": "65ab1234567890abcdef1234",
      "name": "John Doe",
      "email": "john@example.com",
      "role": "user",
      "isActive": true,
      "createdAt": "2024-01-15T10:30:00Z",
      "updatedAt": "2024-01-15T10:30:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Error Response (401 Unauthorized):**
```json
{
  "success": false,
  "error": "invalid email or password"
}
```

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Save the token** from the response for authenticated requests!

---

## 👥 User Management Endpoints

All user endpoints require JWT authentication.

### 3. Get All Users (with Pagination)

Retrieve a paginated list of all users.

**Endpoint:** `GET /api/users?page=1&limit=10`

**Authentication:** Required (JWT)

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10, max: 100)

**Success Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": "65ab1234567890abcdef1234",
      "name": "John Doe",
      "email": "john@example.com",
      "role": "user",
      "isActive": true,
      "createdAt": "2024-01-15T10:30:00Z",
      "updatedAt": "2024-01-15T10:30:00Z"
    },
    {
      "id": "65ab1234567890abcdef1235",
      "name": "Jane Smith",
      "email": "jane@example.com",
      "role": "admin",
      "isActive": true,
      "createdAt": "2024-01-15T11:00:00Z",
      "updatedAt": "2024-01-15T11:00:00Z"
    }
  ],
  "page": 1,
  "limit": 10,
  "total": 25,
  "totalPages": 3
}
```

**cURL Example:**
```bash
curl -X GET "http://localhost:8080/api/users?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 4. Get All Users (Without Pagination) ⚡

Retrieve **ALL** users from the database without pagination limits. This endpoint is optimized for large datasets using concurrent processing with goroutines.

**Endpoint:** `GET /api/users/all`

**Authentication:** Required (JWT)

**Note:** This endpoint is designed for bulk operations and large datasets. It uses:
- MongoDB cursor batching (1000 documents per batch)
- Concurrent goroutines for data processing
- Optimized memory management for handling hundreds of thousands of users

**Success Response (200 OK):**
```json
{
  "success": true,
  "message": "All users retrieved successfully",
  "data": {
    "users": [
      {
        "id": "65ab1234567890abcdef1234",
        "name": "John Doe",
        "email": "john@example.com",
        "role": "user",
        "isActive": true,
        "createdAt": "2024-01-15T10:30:00Z",
        "updatedAt": "2024-01-15T10:30:00Z"
      },
      {
        "id": "65ab1234567890abcdef1235",
        "name": "Jane Smith",
        "email": "jane@example.com",
        "role": "admin",
        "isActive": true,
        "createdAt": "2024-01-15T11:00:00Z",
        "updatedAt": "2024-01-15T11:00:00Z"
      }
      // ... all users in database
    ],
    "total": 460000,
    "count": 460000
  }
}
```

**Performance Features:**
- ✅ Concurrent database fetching (users and count in parallel)
- ✅ Worker pool pattern for concurrent data conversion
- ✅ Batch processing for memory efficiency
- ✅ Handles large datasets (100K+ users) efficiently

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/users/all \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**⚠️ Important Notes:**
- This endpoint returns **all users** including duplicates (same email or name)
- For large datasets (100K+ users), the response may take some time to process
- Response size can be very large - ensure your client can handle large JSON responses
- Consider using pagination endpoint (`/api/users`) for regular use cases
- This endpoint is best suited for data export, bulk operations, or admin dashboards

---

### 5. Get User by ID

Retrieve a specific user by their ID.

**Endpoint:** `GET /api/users/:id`

**Authentication:** Required (JWT)

**URL Parameters:**
- `id`: User ID (MongoDB ObjectID)

**Success Response (200 OK):**
```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "id": "65ab1234567890abcdef1234",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "isActive": true,
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-01-15T10:30:00Z"
  }
}
```

**Error Response (404 Not Found):**
```json
{
  "success": false,
  "error": "user not found"
}
```

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/users/65ab1234567890abcdef1234 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 6. Update User

Update user information. Users can update their own account, admins can update any account.

**Endpoint:** `PUT /api/users/:id`

**Authentication:** Required (JWT)

**Authorization:** User can update own account OR Admin can update any account

**URL Parameters:**
- `id`: User ID (MongoDB ObjectID)

**Request Body (all fields optional):**
```json
{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "password": "newpassword123",
  "role": "admin",
  "isActive": true
}
```

**Success Response (200 OK):**
```json
{
  "success": true,
  "message": "User updated successfully",
  "data": {
    "id": "65ab1234567890abcdef1234",
    "name": "John Updated",
    "email": "john.updated@example.com",
    "role": "admin",
    "isActive": true,
    "createdAt": "2024-01-15T10:30:00Z",
    "updatedAt": "2024-01-15T12:00:00Z"
  }
}
```

**Error Response (403 Forbidden):**
```json
{
  "success": false,
  "error": "You can only modify your own account"
}
```

**Error Response (400 Bad Request):**
```json
{
  "success": false,
  "error": "email already in use"
}
```

**cURL Example:**
```bash
curl -X PUT http://localhost:8080/api/users/65ab1234567890abcdef1234 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "email": "john.updated@example.com"
  }'
```

---

### 7. Delete User

Delete a user account. Only admins can delete users.

**Endpoint:** `DELETE /api/users/:id`

**Authentication:** Required (JWT)

**Authorization:** Admin only

**URL Parameters:**
- `id`: User ID (MongoDB ObjectID)

**Success Response (200 OK):**
```json
{
  "success": true,
  "message": "User deleted successfully",
  "data": null
}
```

**Error Response (403 Forbidden):**
```json
{
  "success": false,
  "error": "Insufficient permissions"
}
```

**Error Response (404 Not Found):**
```json
{
  "success": false,
  "error": "user not found"
}
```

**cURL Example:**
```bash
curl -X DELETE http://localhost:8080/api/users/65ab1234567890abcdef1234 \
  -H "Authorization: Bearer YOUR_ADMIN_JWT_TOKEN"
```

---

## 🔄 Authentication Flow

### Complete Flow Example

1. **Register a new user:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

2. **Login to get JWT token:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

3. **Use the token for authenticated requests:**
```bash
# Save token from login response
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

# Use token in subsequent requests
curl -X GET http://localhost:8080/api/users \
  -H "Authorization: Bearer $TOKEN"
```

---

## ⚠️ Error Handling

The API uses consistent error response format:

```json
{
  "success": false,
  "error": "Error message description"
}
```

### Common HTTP Status Codes

| Status Code | Description |
|-------------|-------------|
| `200` | Success |
| `400` | Bad Request (validation errors) |
| `401` | Unauthorized (missing/invalid token) |
| `403` | Forbidden (insufficient permissions) |
| `404` | Not Found |
| `500` | Internal Server Error |

### Error Examples

**Missing Authorization Header:**
```json
{
  "success": false,
  "error": "Authorization header required"
}
```

**Invalid Token:**
```json
{
  "success": false,
  "error": "Invalid or expired token"
}
```

**Validation Error:**
```json
{
  "success": false,
  "error": "password must be at least 6 characters"
}
```

---

## 🧪 Testing Examples

### Using cURL

#### 1. Register User
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","password":"test123"}'
```

#### 2. Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"test123"}'
```

#### 3. Get All Users (with token)
```bash
curl -X GET "http://localhost:8080/api/users?page=1&limit=5" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

#### 3b. Get All Users (Without Pagination)
```bash
curl -X GET http://localhost:8080/api/users/all \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

#### 4. Get User by ID
```bash
curl -X GET http://localhost:8080/api/users/USER_ID_HERE \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

#### 5. Update User
```bash
curl -X PUT http://localhost:8080/api/users/USER_ID_HERE \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Name"}'
```

#### 6. Delete User (Admin only)
```bash
curl -X DELETE http://localhost:8080/api/users/USER_ID_HERE \
  -H "Authorization: Bearer ADMIN_TOKEN_HERE"
```

### Using Postman

1. **Create a Collection** named "User Management API"

2. **Set Environment Variables:**
   - `base_url`: `http://localhost:8080/api`
   - `token`: (will be set after login)

3. **Create Requests:**

   **Register:**
   - Method: `POST`
   - URL: `{{base_url}}/auth/register`
   - Body (raw JSON):
     ```json
     {
       "name": "John Doe",
       "email": "john@example.com",
       "password": "password123"
     }
     ```

   **Login:**
   - Method: `POST`
   - URL: `{{base_url}}/auth/login`
   - Body (raw JSON):
     ```json
     {
       "email": "john@example.com",
       "password": "password123"
     }
     ```
   - Add Test Script to save token:
     ```javascript
     if (pm.response.code === 200) {
         const response = pm.response.json();
         pm.environment.set("token", response.data.token);
     }
     ```

   **Get All Users (Paginated):**
   - Method: `GET`
   - URL: `{{base_url}}/users?page=1&limit=10`
   - Headers:
     - `Authorization`: `Bearer {{token}}`

   **Get All Users (Without Pagination):**
   - Method: `GET`
   - URL: `{{base_url}}/users/all`
   - Headers:
     - `Authorization`: `Bearer {{token}}`
   - ⚠️ Note: Returns all users - use with caution for large datasets

   **Get User by ID:**
   - Method: `GET`
   - URL: `{{base_url}}/users/:id`
   - Headers:
     - `Authorization`: `Bearer {{token}}`

   **Update User:**
   - Method: `PUT`
   - URL: `{{base_url}}/users/:id`
   - Headers:
     - `Authorization`: `Bearer {{token}}`
     - `Content-Type`: `application/json`
   - Body (raw JSON):
     ```json
     {
       "name": "Updated Name"
     }
     ```

   **Delete User:**
   - Method: `DELETE`
   - URL: `{{base_url}}/users/:id`
   - Headers:
     - `Authorization`: `Bearer {{token}}`

---

## 🔒 Security Features

- ✅ Passwords are hashed using bcrypt (never stored in plain text)
- ✅ JWT tokens with expiration
- ✅ Role-based access control
- ✅ Input validation
- ✅ CORS protection
- ✅ Panic recovery to prevent crashes
- ✅ No sensitive data in error messages

---

## 📝 Notes

- **Password Field**: Never returned in API responses (excluded from JSON)
- **Default Role**: New users are assigned `"user"` role by default
- **Admin Role**: Must be manually updated via update endpoint
- **Pagination**: Default page size is 10, maximum is 100
- **Token Expiration**: Configured via `JWT_EXPIRE_HOURS` in `.env`
- **Email Uniqueness**: Email field has unique index in MongoDB

---

## 🐛 Troubleshooting

### MongoDB Connection Error
```
Failed to connect to MongoDB: ...
```
**Solution**: Ensure MongoDB is running and `MONGO_URI` in `.env` is correct.

### Port Already in Use
```
Failed to start server: listen tcp :8080: bind: address already in use
```
**Solution**: Change `APP_PORT` in `.env` or stop the process using port 8080.

### Invalid Token Error
```
Invalid or expired token
```
**Solution**: Login again to get a new token. Tokens expire after the configured hours.

### Permission Denied
```
Insufficient permissions
```
**Solution**: Ensure your user has the required role (admin for delete operations).

---

## 📄 License

This project is for educational purposes.

---

## 👨‍💻 Author

Built with ❤️ using Go and MongoDB

---

**Happy Coding! 🚀**

