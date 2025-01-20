
# Climbing Stuff App

A web application for managing climbing gyms and shoes. The app provides a RESTful API for creating, reading, updating, and deleting climbing gyms and shoes.

## Features

- **Manage climbing gyms**: create, read, update, and delete gyms. Managing the database is done with an ORM approach (GORM).
- **Manage climbing shoes**: create, read, update, and delete shoes. Managing the database in this case is done with a 'raw SQL' approach.
- **API documentation**: available at `/swagger/index.html`.

## Getting Started

### Building the App
```bash
go build -o ./bin/climbingStuff
```

### Create Documentation
```bash
swag init --parseDependency --parseInternal
```

### Set Environment Variable (Windows)
```powershell
$env:APP_ENV = "development"
```

### Get Environment Variable Value
```powershell
$env:APP_ENV
```

### Open Swagger in Web Browser
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Run Tests
```bash
go test .\tests\ -v
```