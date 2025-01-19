Following driver is important when working with MS SQL Server:
_ "github.com/denisenkom/go-mssqldb" // MS SQL Server Driver

Without it application:
1. builds without any problem,
2. but runs with following message: Error creating connection pool: sql: unknown driver "sqlserver" (forgotten import?) exit status 1


This application reads transaction data from CSV files, imports it into a database, and provides summary statistics. 
It supports batch insertion and deletion of transactions, and can be used to analyze donation data.
