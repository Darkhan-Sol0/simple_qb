
---

# Simple QB

Simple Query Builder for PGX — маленький и удобный пакет для генерации SQL-запросов при работе с базой данных PostgreSQL. Использует аннотации полей структуры для автоматической генерации запросов на основе моделей.

## Особенности

- Простота использования.
- Автоматическая генерация запросов с помощью рефлексивного подхода.
- Низкое воздействие на производительность.

---

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
}
```

| Поле          | Назначение                  |
|-------------------|----------------------------|
| `db:"имя_поля"`   | Имя поля в базе данных       |

---

## Генерация запросов

### Генерация запроса INSERT:

```go
user := User{Name: "John"}
query, args := simple_qb.QueryGenerate(simple_qb.Insert, "users", user, nil)
```

### Генерация запроса SELECT с условием:

```go
queryParams := User{ID: 123}
query, args := simple_qb.QueryGenerate(simple_qb.Select, "users", User{}, queryParams)
```

### Генерация запроса UPDATE:

```go
userData := User{Name: "Jane"}
queryParams := User{ID: 123}
query, args := simple_qb.QueryGenerate(simple_qb.Update, "users", userData, queryParams)
```

---

## API пакета

Основные функции пакета:

- **QueryGenerate(queryType, tableName, data, params)**: Генерация SQL-запроса.
  - `queryType`: Тип запроса (например, `simple_qb.Insert`, `simple_qb.Select`, `simple_qb.Update`).
  - `tableName`: Имя таблицы.
  - `data`: Структура данных для запроса.
  - `params`: Параметры для условия WHERE (может быть картой).

---

## Примеры запросов

### Пример INSERT-запроса:

```sql
INSERT INTO users (name) VALUES ('John');
```

### Пример SELECT-запроса с условием:

```sql
SELECT * FROM users WHERE id = $1;
```

---

## Поддерживаемые версии Go

- Минимальная поддерживаемая версия Go: 1.22+

---

## Авторские права и лицензия

Данный пакет предоставляется бесплатно и с открытым исходным кодом. Вы можете свободно использовать его в своих проектах.

---

## Обратная связь

Если у вас возникли вопросы или есть предложения по улучшению, свяжитесь с автором через указанные контакты.

## PS

Но пока пакет недоделаный. 

---
