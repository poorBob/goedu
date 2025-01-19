Following driver is important when working with MS SQL Server:
_ "github.com/denisenkom/go-mssqldb" // MS SQL Server Driver

Withou it application:
1. builds without any problem,
2. run with following message: Error creating connection pool: sql: unknown driver "sqlserver" (forgotten import?) exit status 1


Application requests stats:
http://localhost:8080/stats - open in webbrowser


Messages App
This is a simple application that allows users to send and retrieve messages. It uses a RESTful API to handle requests and stores messages in a database.

Features
Message insertion: Users can insert messages using a RESTful API.
Message retrieval: Users can retrieve messages by UUID or by a part of the UUID.
Database support: The application supports databases (by default it uses MS SQL Server).
Worker pool: The application uses a worker pool to handle message insertions in batches. 
             There is also second version of worker pool: inserts messages one by one, but it less efficient then the batch version.
Middleware: The Application has simple implementation of rate limitter widdleware
Requests statistics: The application provides a statistics page that shows the number of requests made to the application.