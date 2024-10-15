# Tutorial: Accessing a relational database
[Accessing a relational database](https://go.dev/doc/tutorial/database-access)  

## package
[database/sql](https://pkg.go.dev/database/sql)  
[doc on Accessing relational databases](https://go.dev/doc/database/)  
this documentation does introduce gORM, which is my goal.  
this tutorial will be treated lightly.  

## Create a folder for your code
nothing to note here.  

## Set up a database
I used Jetbrains DataGrip.  
had to reinstall MySQL.  
on first install, run :  
```shell
mysql -u root -p
```
then set the password when prompted.  

```shell
mysql> source /path/to/create-tables.sql
```
this command intakes sql file and runs it in MySQL DBMS.  

## Find and import a database driver¶
[list of SQL DBMS drivers](https://go.dev/wiki/SQLDrivers)  
```shell
go get github.com/go-sql-driver/mysql
```

## Get a database handle and connect¶
this instruction guides me to export env variable to run the code.  

## Query for multiple rows¶

## Query for a single row¶

## Add data¶