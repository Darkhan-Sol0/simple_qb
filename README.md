[English Version](README_eng.md)
---

# Simple QB: Легкий генератор SQL-запросов для PostgreSQL

Simple Query Builder for PGX - Простой и мощный инструмент для генерации SQL-запросов при работе с базой данных PostgreSQL. Используйте аннотации полей структуры для автоматической генерации запросов на основе ваших моделей.

## Особенности

- **Простота использования:** Создавайте запросы легко и быстро.
- **Автоматическая генерация запросов:** Запросы создаются автоматически на основе вашей модели данных.
- **Низкое влияние на производительность:** Оптимизация запросов обеспечивает высокую скорость выполнения. (но это не точно)

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
}
```

| Поле          | Назначение                  |
|-------------------|----------------------------|
| `db:"имя_поля"`   | Имя поля в базе данных       |


```go
type Node struct {
	Tag 	 string // Name colum from table database
	Operator string // Operator from opMap
	Value    any	// Values
	Logic 	 string // AND or OR
}

type FilterNode = []*Node
```

```go
var opMap = map[string]string{
	"eq":      "=",           // равно
	"neq":     "<>",          // неравно
	"lt":      "<",           // меньше
	"lte":     "<=",          // меньше или равно
	"gt":      ">",           // больше
	"gte":     ">=",          // больше или равно
	"like":    "LIKE",        // похоже на (для строковых выражений) 
	"in":      "IN",          // входит в перечень 
	"null":    "IS NULL",     // пустое значение
	"notnull": "IS NOT NULL", // непустое значение 
}
```
---

## Генерация запросов

### Конструктор New:

```go
user := User{Name: "John"}
queryParams := NewFilter(NewNode("id", "eq", 123), NewNodeOr("text", "like", 123)...)
myStruct := simple_qb.New("users", user, queryParams)
```

### Генерация запроса INSERT:

```go
query, args := myStruct.Insert()
```

### Генерация запроса SELECT с условием:

```go
query, args := myStruct.Select()
```

### Генерация запроса UPDATE:

```go
query, args, err := myStruct.Update()
```

### Генерация запроса Delete:

```go
query, args, err := myStruct.Delete()
```

---

## API пакета

Основные функции пакета:

- **New(tableName, data, params)**: Генерация SQL-запроса.
  - `tableName`: Имя таблицы.
  - `data`: Структура данных для запроса.
  - `params`: Параметры для условия WHERE (может быть картой).

- **Insert()**: Генерация SQL-запроса Inserr.

- **Select()**: Генерация SQL-запроса Select.

- **Update()**: Генерация SQL-запроса Update.

- **Delete()**: Генерация SQL-запроса Delete.

- **NewNode(tag, operator string, value any)**: Генерация ноды для параметра м логикой AND
 
- **NewNodeOr(tag, operator string, value any)**: Генерация ноды для параметра м логикой OR

- **NewParam(tagName, operator string, value any)**: Генерация мапы для условия фильтрации.

- **AddParam(tagName, operator string, value any)**: Добавляет условия фильтрации в мапу.

- **Limit(query string, limit, offset int)**: Вспомогательная функция для добавления LIMIT и OFFSET.

- **OrderBy(query string, column, order string)**: Вспомогательная функция для добавления ORDER BY и ASC/DESC. Default ASC
---

## Примеры запросов

### Пример запросов:

```sql
INSERT INTO users (num, text) VALUES ($1, $2)
SELECT num, text FROM users
SELECT num, text FROM users WHERE num = $1
UPDATE users SET (num, text) = ($1, $2) WHERE num = $3
UPDATE users SET (num, text) = ($1, $2) WHERE num >= $3
DELETE FROM users WHERE num = $1
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

**Внимание:** Данный пакет является экспериментальным и активно развивается. Если вы хотите принять участие в разработке или оставить отзыв, пожалуйста, свяжитесь с автором.

## ToDo

- Улучшение системы фильтрации.
- Добавление обработки ошибок.
- Работа с нулевыми значениями.
- Другие улучшения.

---
