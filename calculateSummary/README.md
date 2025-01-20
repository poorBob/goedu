
# Transaction Data Processor

This application reads transaction data from CSV files, imports it into a database, and provides summary statistics. 
It supports batch insertion and deletion of transactions, and can be used to analyze donation data.

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
