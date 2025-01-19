Building app:
go build -o ./bin/climbingStuff

Create doc with:
swag init --parseDependency --parseInternal

on Windows set 'development' env variable with:
$env:APP_ENV = "development"

get APP_ENC value:
$env:APP_ENV

open Swagger in web browser:
http://localhost:8080/swagger/index.html

Run tests:
go test .\tests\ -v


Climbing Stuff App

A web application for managing climbing gyms and shoes. The app provides a RESTful API for creating, reading, updating, and deleting climbing gyms and shoes.

Features

Manage climbing gyms: create, read, update, and delete gyms. Managing DB is done with ORM approach (gorm).
Manage climbing shoes: create, read, update, and delete shoes. Managing DB in this case is done with 'raw SQL' approach.
API documentation available at /swagger/index.html
Getting Started