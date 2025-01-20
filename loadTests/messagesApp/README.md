
# Messages App

This is a simple application that allows users to send and retrieve messages. It uses a RESTful API to handle requests and stores messages in a database.

## Features

- **Message Insertion**: Users can insert messages using a RESTful API.
- **Message Retrieval**: Users can retrieve messages by UUID or by a part of the UUID.
- **Database Support**: The application supports databases (by default it uses MS SQL Server).
- **Worker Pool**: The application uses a worker pool to handle message insertions in batches.
  - There is also a second version of the worker pool that inserts messages one by one, but it is less efficient than the batch version.
- **Middleware**: The application includes a simple implementation of a rate limiter middleware.
- **Request Statistics**: The application provides a statistics page that shows the number of requests made to the application.

## Important Dependency

When working with MS SQL Server, the following driver is required:
```go
_ "github.com/denisenkom/go-mssqldb" // MS SQL Server Driver
```

### Notes

Without the driver, the application:
1. Builds without any problem.
2. But runs with the following error message:
   ```
   Error creating connection pool: sql: unknown driver "sqlserver" (forgotten import?)
   exit status 1
   ```

## Application Request Stats

You can view the request statistics by opening the following URL in a web browser:
[http://localhost:8080/stats](http://localhost:8080/stats)
