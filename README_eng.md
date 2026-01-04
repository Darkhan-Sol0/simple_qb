# Simple QB: Lightweight SQL Query Builder for PostgreSQL

Simple Query Builder for PGX - A simple yet powerful tool for generating SQL queries when working with PostgreSQL database. Use struct field annotations for automatic query generation based on your models.

## Features

- **Easy to Use:** Create queries quickly and effortlessly.
- **Flexible Condition Chaining:** Support for complex WHERE conditions with method chaining.
- **Security:** Parameterized queries protect against SQL injection.
- **(Almost) Full PostgreSQL Support:** LIKE, IN, BETWEEN, IS NULL operators and more.

---

## Note

Currently only for the pgx driver package to work with PostgreSQL.

## Compatibility

✅ **Works with:** [pgx](https://github.com/jackc/pgx)

## Installation

Add the package to your project:

```sh
go get github.com/jackc/pgx/v5
go get github.com/Darkhan-Sol0/simple_qb
```

Import the package in your code:

```go
import "github.com/Darkhan-Sol0/simple_qb"
```

---

## Getting Started

Before using the package, you need to describe your data structure with field annotations. The `db:"field_name"` annotation establishes correspondence between struct fields and database fields.

### Example Structure:

```go
type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
    Age  int    `db:"age"`
    Email string `db:"email"`
}
```

| Field          | Purpose                    |
|----------------|----------------------------|
| `db:"field_name"` | Database field name       |

---

## Query Generation

### Constructor New:

```go
qb := simple_qb.New("users")
```

### SELECT Query Generation:

```go
// Simple SELECT
query, args := qb.Select(User{}).Generate()
// SELECT id, name, age, email FROM users

// SELECT with conditions
query, args := qb.Select(User{}).
    Params(
        simple_qb.NewParam(
            simple_qb.NewNode("age").Gr(18)).
            And(simple_qb.NewNode("name").Like("%John%")
        )
    ).
    Generate()
// SELECT id, name, age, email FROM users WHERE (age > $1) AND (name LIKE $2)

// SELECT with ORDER BY and LIMIT
query, args := qb.Select(User{}).
    OrderBy("age", "DESC").
    Limit(10, 0).
    Generate()
// SELECT id, name, age, email FROM users ORDER BY age DESC LIMIT 10
```

### INSERT Query Generation:

```go
user := User{Name: "John", Age: 25, Email: "john@example.com"}
query, args := qb.Insert(user).Generate()
// INSERT INTO users (name, age, email) VALUES ($1, $2, $3)

// INSERT with RETURNING
query, args := qb.Insert(user).
    Returning("id").
    Generate()
// INSERT INTO users (name, age, email) VALUES ($1, $2, $3) RETURNING id
```

### UPDATE Query Generation:

```go
user := User{Name: "John Updated", Age: 26}
query, args := qb.Update(user).
    Params(
        simple_qb.NewParam(
            simple_qb.NewNode("id").Eq(1)
        )
    ).
    Generate()
// UPDATE users SET (name, age) = ROW($1, $2) WHERE (id = $3)
```

### DELETE Query Generation:

```go
query, args := qb.Delete(User{}).
    Params(
        simple_qb.NewParam(
            simple_qb.NewNode("id").Eq(1)
        )
    ).
    Generate()
// DELETE FROM users WHERE (id = $1)
```

### WHERE Conditions:

```go
// Equality
simple_qb.NewNode("age").Eq(25)
// age = $1

// Inequality
simple_qb.NewNode("age").NotEq(25)
// age <> $1

// Greater/Less than
simple_qb.NewNode("age").Gr(18)      // age > $1
simple_qb.NewNode("age").GrEq(18)    // age >= $1
simple_qb.NewNode("age").Less(65)    // age < $1
simple_qb.NewNode("age").LessEq(65)  // age <= $1

// LIKE
simple_qb.NewNode("name").Like("%John%")
// name LIKE $1

// IN
simple_qb.NewNode("id").In([]any{1, 2, 3})
// id IN ($1, $2, $3)

// BETWEEN
simple_qb.NewNode("age").Between(18, 65)
// age BETWEEN $1 AND $2

// NULL checks
simple_qb.NewNode("email").Null()     // email IS NULL
simple_qb.NewNode("email").NotNull()  // email IS NOT NULL
```

### Logical Operators:

```go
// Condition chains within a single node
simple_qb.NewNode("age").Eq(25).Or().Less(30)
// age = $1 OR age < $2

simple_qb.NewNode("name").Eq("John").And().Like("%Doe%")
// name = $1 AND name LIKE $2

// Combining nodes with AND/OR
params := simple_qb.NewParam(
    simple_qb.NewNode("age").Gr(18)
).And(
    simple_qb.NewNode("name").Like("%John%")
).Or(
    simple_qb.NewNode("email").NotNull()
)
// (age > $1) AND (name LIKE $2) OR (email IS NOT NULL)
```

### Complex Conditions

```go
// Complex nested conditions
params := simple_qb.NewParam(
    simple_qb.NewNode("age").Between(18, 65).Or().Null()
).
And(
    simple_qb.NewNode("name").Like("J%").And().NotEq("")
)
// (age BETWEEN $1 AND $2 OR age IS NULL) AND (name LIKE $3 AND name <> $4)
```

## Additional Features

### COUNT Queries

```go
// COUNT all records
query, args := qb.Select(nil).Count("").Generate()
// SELECT COUNT(*) FROM users

// COUNT specific field
query, args := qb.Select(nil).Count("id").
    Params(
        simple_qb.NewParam(
            simple_qb.NewNode("active").Eq(true)
        )
    ).
    Generate()
// SELECT COUNT(id) FROM users WHERE (active = $1)
```

### ORDER BY

```go
qb.Select(User{}).OrderBy("age", "DESC")
// ORDER BY age DESC

qb.Select(User{}).OrderBy("name", "ASC")
// ORDER BY name ASC
```

### LIMIT and OFFSET

```go
// Simple LIMIT
qb.Select(User{}).Limit(10, 0)
// LIMIT 10

// LIMIT with OFFSET (pagination)
qb.Select(User{}).Limit(10, 20)
// LIMIT 10 OFFSET 20
```

---

## Package API

Main package functions:
Simple generation of SQL query strings.

---

## Complete Examples

### Example 1: Finding Users

```go
func FindActiveUsers() (string, []any) {
    qb := simple_qb.New("users")
    
    return qb.Select(User{}).
        Params(
            simple_qb.NewParam(
                simple_qb.NewNode("active").Eq(true)
            ).
            And(
                simple_qb.NewNode("age").Between(18, 65)
            )
        ).
        OrderBy("name", "ASC").
        Limit(50, 0).
        Generate()
}
// SELECT id, name, age, email FROM users 
// WHERE (active = $1) AND (age BETWEEN $2 AND $3)
// ORDER BY name ASC LIMIT 50
```

### Example 2: Update with Complex Condition

```go
func UpdateUserEmail(userID int, newEmail string) (string, []any) {
    user := User{Email: newEmail}
    
    return simple_qb.New("users").
        Update(user).
        Params(
            simple_qb.NewParam(
                simple_qb.NewNode("id").Eq(userID)
            ).
            And(
                simple_qb.NewNode("active").Eq(true)
            )
        ).
        Generate()
}
// UPDATE users SET (email) = ROW($1) 
// WHERE (id = $2) AND (active = $3)
```

### Example 3: Deleting Inactive Users

```go
func DeleteInactiveUsers() (string, []any) {
    return simple_qb.New("users").
        Delete(User{}).
        Params(
            simple_qb.NewParam(
                simple_qb.NewNode("active").Eq(false)
            ).
            Or(
                simple_qb.NewNode("last_login").Less(time.Now().AddDate(0, -6, 0))
            )
        ).
        Generate()
}
// DELETE FROM users 
// WHERE (active = $1) OR (last_login < $2)
```

### Example Generated Queries:

```sql
INSERT INTO users (num, text) VALUES ($1, $2)
SELECT num, text FROM users
SELECT num, text FROM users WHERE num = $1
UPDATE users SET (num, text) = ($1, $2) WHERE num = $3
UPDATE users SET (num, text) = ($1, $2) WHERE num >= $3
DELETE FROM users WHERE num = $1
```

---

# Notes

## Supported Go Versions

- Minimum supported Go version: 1.22+

## Limitations
Package only for driver **pgx**
- PostgreSQL syntax only
- Placeholders format `$1, $2, $3...`

---

## Copyright and License

This package is provided free and open-source. You are free to use it in your projects.

---

## Feedback

If you have questions or suggestions for improvement, contact the author through the provided contacts.

## PS

**Note:** This package is experimental and actively developed. If you'd like to participate in development or leave feedback, please contact the author.

## ToDo

- Not sure what else is needed yet.