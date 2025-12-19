[English Version](README_eng.md)
---

# Simple QB: Легкий генератор SQL-запросов для PostgreSQL

Simple Query Builder for PGX - Простой и мощный инструмент для генерации SQL-запросов при работе с базой данных PostgreSQL. Используйте аннотации полей структуры для автоматической генерации запросов на основе ваших моделей.

## Особенности

- **Простота использования:** Создавайте запросы легко и быстро.
- **Гибкие цепочки условий:** Поддержка сложных условий WHERE с цепочками вызовов.
- **Безопасность:** Параметризованные запросы защищают от SQL-инъекций.
- **(Почти) Полная поддержка PostgreSQL:** Операторы LIKE, IN, BETWEEN, IS NULL и другие.

---

## Примечание

Пока для пакета драйвера pgx для работы в PostgreSQL

## Установка

Добавьте пакет в ваш проект:

```sh
go get github.com/Darkhan-Sol0/simple_qb
```

Импортируйте пакет в ваш код:

```go
import "github.com/Darkhan-Sol0/simple_qb"
```

---

## Начало работы

Перед использованием пакета необходимо описать структуру данных с аннотациями полей. Аннотация `db:"имя_поля"` устанавливает соответствие полей структуры с полями в базе данных.

### Пример структуры:

```go
type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
    Age  int    `db:"age"`
    Email string `db:"email"`
}

```

| Поле          | Назначение                  |
|-------------------|----------------------------|
| `db:"имя_поля"`   | Имя поля в базе данных       |


---

## Генерация запросов

### Конструктор New:

```go
qb := simple_qb.New("users")
```

### Генерация запроса SELECT:

```go
// Простой SELECT
query, args := qb.Select(User{}).Generate()
// SELECT id, name, age, email FROM users

// SELECT с условиями
query, args := qb.Select(User{}).
    Params(simple_qb.NewParam(
        simple_qb.NewNode("age").Gr(18).
        And(simple_qb.NewNode("name").Like("%John%"))
    )).
    Generate()
// SELECT id, name, age, email FROM users WHERE (age > $1) AND (name LIKE $2)

// SELECT с ORDER BY и LIMIT
query, args := qb.Select(User{}).
    OrderBy("age", "DESC").
    Limit(10, 0).
    Generate()
// SELECT id, name, age, email FROM users ORDER BY age DESC LIMIT 10
```

### Генерация запроса INSERT:

```go
user := User{Name: "John", Age: 25, Email: "john@example.com"}
query, args := qb.Insert(user).Generate()
// INSERT INTO users (name, age, email) VALUES ($1, $2, $3)

// INSERT с RETURNING
query, args := qb.Insert(user).
    Returning("id").
    Generate()
// INSERT INTO users (name, age, email) VALUES ($1, $2, $3) RETURNING id
```

### Генерация запроса UPDATE:

```go
user := User{Name: "John Updated", Age: 26}
query, args := qb.Update(user).
    Params(simple_qb.NewParam(
        simple_qb.NewNode("id").Eq(1)
    )).
    Generate()
// UPDATE users SET (name, age) = ROW($1, $2) WHERE (id = $3)
```

### Генерация запроса DELETE:

```go
query, args := qb.Delete(User{}).
    Params(simple_qb.NewParam(
        simple_qb.NewNode("id").Eq(1)
    )).
    Generate()
// DELETE FROM users WHERE (id = $1)
```

### Условия WHERE:

```go
// Равенство
simple_qb.NewNode("age").Eq(25)
// age = $1

// Неравенство
simple_qb.NewNode("age").NotEq(25)
// age <> $1

// Больше/Меньше
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

// NULL проверки
simple_qb.NewNode("email").Null()     // email IS NULL
simple_qb.NewNode("email").NotNull()  // email IS NOT NULL
```

### Логические операторы:

```go
// Цепочки условий внутри одной ноды
simple_qb.NewNode("age").Eq(25).Or().Less(30)
// age = $1 OR age < $2

simple_qb.NewNode("name").Eq("John").And().Like("%Doe%")
// name = $1 AND name LIKE $2

// Комбинирование нод через AND/OR
params := simple_qb.NewParam(
    simple_qb.NewNode("age").Gr(18)
).And(
    simple_qb.NewNode("name").Like("%John%")
).Or(
    simple_qb.NewNode("email").NotNull()
)
// (age > $1) AND (name LIKE $2) OR (email IS NOT NULL)
```

### Сложные условия

```go
// Сложные вложенные условия
params := simple_qb.NewParam(
    simple_qb.NewNode("age").Between(18, 65).Or().Null()
).And(
    simple_qb.NewNode("name").Like("J%").And().NotEq("")
)
// (age BETWEEN $1 AND $2 OR age IS NULL) AND (name LIKE $3 AND name <> $4)

```
## Дополнительные возможности
### COUNT запросы
``` go

// COUNT всех записей
query, args := qb.Select(nil).Count("").Generate()
// SELECT COUNT(*) FROM users

// COUNT конкретного поля
query, args := qb.Select(nil).Count("id").
    Params(simple_qb.NewParam(
        simple_qb.NewNode("active").Eq(true)
    )).
    Generate()
// SELECT COUNT(id) FROM users WHERE (active = $1)
```
### ORDER BY
``` go

qb.Select(User{}).OrderBy("age", "DESC")
// ORDER BY age DESC

qb.Select(User{}).OrderBy("name", "ASC")
// ORDER BY name ASC
```
### LIMIT и OFFSET
```go

// Просто LIMIT
qb.Select(User{}).Limit(10, 0)
// LIMIT 10

// LIMIT с OFFSET (пагинация)
qb.Select(User{}).Limit(10, 20)
// LIMIT 10 OFFSET 20
```
---

## API пакета

Основные функции пакета:
Простая генерация строки SQL запроса.

---

## Полные примеры
### Пример 1: Поиск пользователей
```go

func FindActiveUsers() (string, []any) {
    qb := simple_qb.New("users")
    
    return qb.Select(User{}).
        Params(simple_qb.NewParam(
            simple_qb.NewNode("active").Eq(true).
            And(simple_qb.NewNode("age").Between(18, 65))
        )).
        OrderBy("name", "ASC").
        Limit(50, 0).
        Generate()
}
// SELECT id, name, age, email FROM users 
// WHERE (active = $1) AND (age BETWEEN $2 AND $3)
// ORDER BY name ASC LIMIT 50
```
### Пример 2: Обновление с комплексным условием
```go

func UpdateUserEmail(userID int, newEmail string) (string, []any) {
    user := User{Email: newEmail}
    
    return simple_qb.New("users").
        Update(user).
        Params(simple_qb.NewParam(
            simple_qb.NewNode("id").Eq(userID).
            And(simple_qb.NewNode("active").Eq(true))
        )).
        Generate()
}
// UPDATE users SET (email) = ROW($1) 
// WHERE (id = $2) AND (active = $3)
```
### Пример 3: Удаление неактивных пользователей
```go

func DeleteInactiveUsers() (string, []any) {
    return simple_qb.New("users").
        Delete(User{}).
        Params(simple_qb.NewParam(
            simple_qb.NewNode("active").Eq(false).
            Or(simple_qb.NewNode("last_login").Less(time.Now().AddDate(0, -6, 0)))
        )).
        Generate()
}
// DELETE FROM users 
// WHERE (active = $1) OR (last_login < $2)
```

### Пример получаемых запросов:

```sql
INSERT INTO users (num, text) VALUES ($1, $2)
SELECT num, text FROM users
SELECT num, text FROM users WHERE num = $1
UPDATE users SET (num, text) = ($1, $2) WHERE num = $3
UPDATE users SET (num, text) = ($1, $2) WHERE num >= $3
DELETE FROM users WHERE num = $1
```
---

# Примечания
## Поддерживаемые версии Go

- Минимальная поддерживаемая версия Go: 1.22+

## Ограничения

    Только PostgreSQL синтаксис

    Нет поддержки JOIN (в разработке)

    Нет поддержки GROUP BY и HAVING (в разработке)
---

## Авторские права и лицензия

Данный пакет предоставляется бесплатно и с открытым исходным кодом. Вы можете свободно использовать его в своих проектах.

---

## Обратная связь

Если у вас возникли вопросы или есть предложения по улучшению, свяжитесь с автором через указанные контакты.

## PS

**Внимание:** Данный пакет является экспериментальным и активно развивается. Если вы хотите принять участие в разработке или оставить отзыв, пожалуйста, свяжитесь с автором.

## ToDo

	Поддержка JOIN операций

    Поддержка GROUP BY и HAVING

    Поддержка подзапросов (subqueries)

    Добавление большего количества SQL операторов

    Поддержка транзакций

    Кэширование подготовленных запросов

---
