# go-restapi-cleancode-temp
Template RESTful API using Golang with 3G technologies (Golang, Gin, GORM) and using Logrus, Viper for configuration, and PostgreSQL database

## Overview
This project is a template for building a RESTful API using Golang. It leverages the following technologies:
- **Golang**: The programming language used for development.
- **Gin**: A web framework for building high-performance APIs.
- **GORM**: An ORM library for Golang, used for interacting with the PostgreSQL database.
- **Logrus**: A structured logger for Go.
- **Viper**: A configuration management library.

## Features
- RESTful API structure
- Configuration management with Viper
- Logging with Logrus
- Database interaction using GORM with PostgreSQL
- Clean code architecture

## Getting Started
### Prerequisites
- Go 1.23.3 or later
- PostgreSQL database

### Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/dwiprastyoisworo/go-restapi-cleancode-temp.git
    cd go-restapi-cleancode-temp
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up your PostgreSQL database and update the configuration file `user.config.json` in the `file/configs` directory.

### Configuration
The configuration is managed using Viper. Ensure you have a `user.config.json` file in the `file/configs` directory with the following structure:
```json
{
  "app": {
    "port": "8080"
  },
  "postgres": {
    "host": "localhost",
    "port": 5432,
    "user": "yourusername",
    "password": "yourpassword",
    "database": "yourdatabase",
    "ssl": "disable",
    "max_idle_time": 30,
    "max_life_time": 300,
    "max_open_conns": 10,
    "max_idle_conns": 5
  }
}