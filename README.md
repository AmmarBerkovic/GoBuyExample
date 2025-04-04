# GO User API

## Table of Contents
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Installation & Setup](#installation--setup)
- [Configuration](#configuration)
- [Proposed SQL Queries](#proposed-sql-queries)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)

## Features
- User Management (Create, Read, Update, Delete)
- Layered architecture for maintainability
- RESTful API design

## Tech Stack
- **Go** (Programming Language)
- **Chi** (Lightweight router for HTTP handling)
- **GORM** (ORM for MySQL)
- **Viper** (Configuration management)
- **MySQL** (Relational Database)

## Installation & Setup
### Prerequisites
Ensure you have the following installed:
- [Go](https://go.dev/dl/)
- [MySQL](https://dev.mysql.com/downloads/installer/)

### Steps to Run the Project
1. Clone the repository:
   ```sh
   git clone https://github.com/AmmarBerkovic/GoBuyExample.git
   cd GoBuyExample
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Configure your database settings in `config/config.yaml` (see [Configuration](#configuration)).
4. Run the application:
   ```sh
   go run main.go
   ```

## Configuration
The application reads database configurations from `config/config.yaml`:
```yaml
database:
  user: "root"
  password: "password"
  host: "localhost"
  port: 3306
  name: "gobuyexample"
```
Modify these values based on your MySQL setup.

## Proposed SQL Queries
For local testing, you can execute the following SQL commands:
```sql
CREATE DATABASE go_user_db;
CREATE USER 'go_user'@'localhost' IDENTIFIED BY 'go_password';
GRANT ALL PRIVILEGES ON go_user_db.* TO 'go_user'@'localhost';
FLUSH PRIVILEGES;

CREATE TABLE user_dbs (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    Age INT NOT NULL
);

INSERT INTO user_dbs (Name, Email, Age)
VALUES
  ('Alice Johnson', 'alice.johnson@example.com', 30),
  ('Bob Smith', 'bob.smith@example.com', 25),
  ('Charlie Brown', 'charlie.brown@example.com', 35),
  ('Diana Prince', 'diana.prince@example.com', 28),
  ('Frank Castle', 'frank.castle@example.com', 38),
  ('Grace Harper', 'grace.harper@example.com', 27),
  ('Henry Adams', 'henry.adams@example.com', 45),
  ('Isabella Clarke', 'isabella.clarke@example.com', 32),
  ('Jack Dawson', 'jack.dawson@example.com', 29),
  ('Katherine Pierce', 'katherine.pierce@example.com', 34),
  ('Liam OBrien', 'liam.obrien@example.com', 31),
  ('Monica Bell', 'monica.bell@example.com', 26),
  ('Nathan Drake', 'nathan.drake@example.com', 37),
  ('Olivia Turner', 'olivia.turner@example.com', 33);
```

## Generating User with Protobuf
The project supports user creation with **Protocol Buffers (Protobuf)** for efficient serialization. To generate the user model:
1. Install the Protocol Buffers compiler (`protoc`):
   ```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   ```
2. Define the `user.proto` file:
   ```proto
   syntax = "proto3";

   package user;

   message User {
     int32 id = 1;
     string name = 2;
     string email = 3;
     int32 age = 4;
   }
   ```
3. Generate Go code from the `.proto` file:
   ```sh
   protoc --go_out=. --go_opt=paths=source_relative user.proto
   ```
4. Use the generated Go struct in your application to handle user data efficiently.

## API Endpoints
| Method | Endpoint      | Description          |
|--------|-------------|----------------------|
| GET    | /users      | Fetch all users (supports `page` and `page_size` parameters) |
| POST   | /users      | Create a new user    |
| PUT    | /users/{id} | Update a user        |
| DELETE | /users/{id} | Delete a user        |

## Project Structure
```
├───config          # Configuration files
├───db              # Database connection setup
└───internal
    ├───handlers    # HTTP handlers
    ├───models      # Data models (GORM)
    ├───repository  # Database queries
    ├───services    # Business logic
```