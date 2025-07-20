# 📝 Todo App in Golang

Welcome to the **Todo App in Golang**! This project is a simple, modern, and efficient implementation of a Todo application built using Go. 🚀

## 📂 Project Structure

Here's an overview of the project structure:

```
├── .gitignore # Git ignore rules 
├── go.mod # Go module dependencies 
├── go.sum # Go module checksums 
├── readme.md # Project documentation 
├── cmd/ # Application entry point 
│ └── todo-app-in-golang/ 
│  └── main.go # Main application logic 
├── config/ # Configuration files 
│ └── local.yaml # Local environment configuration 
├── internal/ # Internal application logic 
│ └── config/ 
│ └── config.go # Configuration loader 
└── storage/ # Storage directory
```
## ⚙️ Configuration

The application uses a YAML configuration file located at `config/local.yaml`. Here's an example of the configuration:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8000"
```

## Environment Variables
You can override the configuration using environment variables:

- CONFIG_PATH: Path to the configuration file.
- ENV: Application environment (e.g., dev, production).
- STORAGE_PATH: Path to the storage database.
- ADDR: HTTP server address.

## 🚀 Getting Started
Prerequisites
- Go 1.20 or higher
- SQLite (for database storage)
## Installation
1- Clone the repository:
```bash
git clone https://github.com/suryanshvermaa/todo-app-in-golang.git
cd todo-app-in-golang
```
2- Install dependencies:
```bash
go mod tidy
```
3- Run the application:
```bash
go run cmd/todo-app-in-golang/main.go
```

## 🛠️ Features
- Configuration Management: Easily manage configurations using YAML and environment variables.
- HTTP Server: A lightweight HTTP server to handle requests.
- Storage: SQLite database for persistent storage.

## 📜 License
This project is licensed under the MIT License. See the LICENSE file for details.

## 🤝 Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

Made with ❤️ Suryansh Verma