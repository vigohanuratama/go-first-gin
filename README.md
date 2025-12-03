# 🔥 go-first-gin

*First REST API built with Go + Gin --- part of my learning from Udemy:
"Go --- The Complete Guide (Golang)"*

## 📘 About The Project

**go-first-gin** is a simple REST API built using the **Gin Gonic**
framework.\
This project is part of my learning journey into Go backend development,
focusing on building APIs, structuring Go projects, handling JSON, and
understanding routing.

## 🧰 Built With

-   Go (Golang)
-   Gin Gonic
-   Go Modules

## 📂 Project Structure

    /go-first-gin
      ├── main.go
      ├── go.mod
      ├── controllers/
      ├── routes/
      ├── models/
      ├── services/
      ├── README.md
      └── .gitignore

## 🚀 Getting Started

### Prerequisite

    go version

### Run Locally

    git clone https://github.com/vigohanuratama/go-first-gin.git
    cd go-first-gin
    go mod tidy
    go run main.go

## 🔗 Endpoints (Example)

  Method   Route    Description
  -------- -------- -----------------
  GET      /items   Get all items
  POST     /items   Create new item

## ✨ Features

-   REST API using Gin\
-   JSON binding & validation\
-   Modular architecture

## 🔮 Future Improvements

-   Database integration
-   JWT authentication
-   Swagger docs
-   CI/CD

## 📄 License

MIT License.
