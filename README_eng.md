[Russian Version](./readme.md)
---

# Simple QB: Easy SQL Generator for PostgreSQL

Simple Query Builder for PGX is a simple yet powerful tool for generating SQL queries when working with the PostgreSQL database. Use field annotations of your structure to automatically generate queries based on your models.

## Features

- **Ease of use:** Create queries easily and quickly.
- **Automatic query generation:** Queries are generated automatically based on your data model.
- **Low performance impact:** Optimized queries ensure high execution speed. (not guaranteed)

---

## Note

Currently designed for the pgx driver package for working in PostgreSQL.

## Installation

Add the package to your project:

```sh
go get github.com/Darkhan-Sol0/simple_qb
```

Import the package into your code:

```go
import "github.com/Darkhan-Sol0/simple_qb"
```

---

## Getting Started

Before using the package, you need to describe the data structure with field annotations. The annotation `db:"field_name"` establishes correspondence between structure fields and database fields.

### Example Structure:

```go
type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
}
```

| Field            | Purpose               |
|------------------|-----------------------|
| `db:"field_name"`| Database field name   |

```go
type User struct {
    ID   int    `db:"id" op:"eq"`
    Name string `db:"name" op:"gt"`
}
```

| Field              | Purpose                     |
|--------------------|-----------------------------|
| `op:"operator_name"`| Filter operation by value   |

```go
var opMap = map[string]string{
	"eq":      "=",           // equal
	"neq":     "<>",          // not equal
	"lt":      "<",           // less than
	"lte":     "<=",          // less than or equal
	"gt":      ">",           // greater than
	"gte":     ">=",          // greater than or equal
	"like":    "LIKE",        // like (for string expressions) (deprecated)
	"in":      "IN",          // in list (deprecated)
	"null":    "IS NULL",     // null value (deprecated)
	"notnull": "IS NOT NULL", // non-null value (deprecated)
}
```
---

## Generating Queries

### Constructor New:

```go
user := User{Name: "John"}
queryParams := User{ID: 123}
myStruct := simple_qb.New("users", user, queryParams)
```

### Generate INSERT Query:

```go
query, args := myStruct.Insert()
```

### Generate SELECT Query with Condition:

```go
query, args := myStruct.Select()
```

### Generate UPDATE Query:

```go
query, args, err := myStruct.Update()
```

### Generate DELETE Query:

```go
query, args, err := myStruct.Delete()
```

---

## Package API

Main functions provided by the package:

- **New(tableName, data, params)**: Generates an SQL query.
  - `tableName`: Table name.
  - `data`: Data structure for the query.
  - `params`: Parameters for the WHERE condition (can be a map).

- **Insert()**: Generates an Insert SQL query.

- **Select()**: Generates a Select SQL query.

- **Update()**: Generates an Update SQL query.

- **Delete()**: Generates a Delete SQL query.
---

## Query Examples

### Sample Queries:

```sql
INSERT INTO users (num, text) VALUES ($1, $2)
SELECT num, text FROM users
SELECT num, text FROM users WHERE num = $1
UPDATE users SET (num, text) = ($1, $2) WHERE num = $3
UPDATE users SET (num, text) = ($1, $2) WHERE num >= $3
DELETE FROM users WHERE num = $1
```
---

## Supported Go Versions

- Minimum supported Go version: 1.22+

---

## Copyright and License

This package is free and open-source. You can freely use it in your projects.

---

## Feedback

If you have any questions or suggestions for improvement, please contact the author through the specified contacts.

## PS

**Attention:** This package is experimental and actively developed. If you want to participate in development or leave feedback, please contact the author.

## ToDo

- Improve filtering system.
- Add error handling.
- Work with nullable values.
- Other improvements.

---