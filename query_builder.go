// Package simple_qb provides a lightweight toolkit for building SQL queries for standard CRUD operations.
// It supports creating INSERT, SELECT, and UPDATE statements by processing structured data and customizable parameters.
// Built-in methods allow automatic generation of SQL queries with support for WHERE clauses and prepared expressions.
package simple_qb

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Insert is a template for constructing INSERT INTO commands.
// Select is a template for constructing SELECT commands.
// Update is a template for constructing UPDATE commands.
// Where is a template for constructing WHERE clauses.
var (
	insertTemplate = "INSERT INTO %s (%s) VALUES (%s)"
	selectTemplate = "SELECT %s FROM %s"
	updateTemplate = "UPDATE %s SET (%s) = ROW(%s)"
	deleteTemplate = "DELETE FROM %s"
	whereTemplate  = "WHERE %s"
)

// Tag identifies column names within structures.
// Op specifies the operation applied to data in WHERE conditions.
const (
	tag = "db"
)

type Node struct {
	Operator string
	Value    any
}

type FilterNode = map[string]*Node

// пока такие теги, может посже изменить

// Operation mappings convert internal representations into corresponding SQL operators.
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

type (
	qBilder struct {
		table  string
		data   any
		params FilterNode
	}

	QBilder interface {
		Select() (query string, args []any)
		Insert() (query string, args []any)
		Update() (query string, args []any, err error)
		Delete() (query string, args []any, err error)

		AddFilter(tagName, operator string, value any)
	}
)

func New(table string, data any, params FilterNode) QBilder {
	return &qBilder{
		table:  table,
		data:   data,
		params: params,
	}
}

func (q *qBilder) AddFilter(tagName, operator string, value any) {
	q.params[tagName] = newNode(operator, value)
}

func NewFillter(tagName, operator string, value any) FilterNode {
	return map[string]*Node{
		tagName: newNode(operator, value),
	}
}

func newNode(operator string, value any) *Node {
	return &Node{
		Operator: operator,
		Value:    value,
	}
}

func (q *qBilder) Select() (query string, args []any) {
	colums := getColums(q.data)
	query = fmt.Sprintf(selectTemplate, strings.Join(colums, ", "), q.table)
	if q.params != nil {
		where, ar := getWhere(q.params, 0)
		args = append(args, ar...)
		query = fmt.Sprintf("%s %s", query, where)
	}
	return query, args
}

func (q *qBilder) Insert() (query string, args []any) {
	colums := getColums(q.data)
	placeholders := getPlaceholders(len(colums))
	args = getArguments(q.data)
	query = fmt.Sprintf(insertTemplate, q.table, strings.Join(colums, ", "), strings.Join(placeholders, ", "))
	return query, args
}

func (q *qBilder) Update() (query string, args []any, err error) {
	if q.params == nil {
		return "", nil, errors.New("cannot perform UPDATE without WHERE condition")
	}
	colums := getColums(q.data)
	placeholders := getPlaceholders(len(colums))
	args = getArguments(q.data)
	query = fmt.Sprintf(updateTemplate, q.table, strings.Join(colums, ", "), strings.Join(placeholders, ", "))
	if q.params != nil {
		where, ar := getWhere(q.params, len(colums))
		args = append(args, ar...)
		query = fmt.Sprintf("%s %s", query, where)
	}
	return query, args, nil
}

func (q *qBilder) Delete() (query string, args []any, err error) {
	if q.params == nil {
		return "", nil, errors.New("cannot perform DELETE without WHERE condition")
	}
	query = fmt.Sprintf(deleteTemplate, q.table)
	if q.params != nil {
		where, ar := getWhere(q.params, 0)
		args = append(args, ar...)
		query = fmt.Sprintf("%s %s", query, where)
	}
	return query, args, nil
}

// getArguments extracts arguments from structured data, considering special tags (db).
// Returns slice of interfaces containing actual values for insertion or updating.
func getArguments(data any) (args []any) {
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get(tag)
		val := v.Field(i)
		if dbTag != "" && dbTag != "-" && val != reflect.Zero(val.Type()) {
			args = append(args, val.Interface())
		}
	}
	return args
}

// getColums retrieves column names from structured data, utilizing the special db tag.
// Returns slice of strings representing column names.
func getColums(data any) (colums []string) {
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get(tag)
		if dbTag != "" && dbTag != "-" {
			colums = append(colums, dbTag)
		}
	}
	return colums
}

// getPlaceholders generates placeholder sequence ($n) for SQL queries.
// Number of placeholders is determined by parameter count.
func getPlaceholders(count int) (placeholders []string) {
	for i := 1; i <= count; i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
	}
	return
}

// Нужно както впиндюрить оператор OR

// getWhere builds WHERE condition string based on structured data and special tags.
// Returns formatted WHERE clause combining columns and operators via AND.

func getWhere(data FilterNode, startIndex int) (query string, args []any) {
	var colums []string
	count := 1
	for i, v := range data {
		tag := v.Operator
		val := v.Value
		if tag != "" && opMap[tag] != "" {
			switch tag {
			case "in":
				colums = append(colums, fmt.Sprintf("%s %s ($%d)", i, opMap[tag], startIndex+count))
				args = append(args, val)
			case "null", "notnull":
				colums = append(colums, fmt.Sprintf("%s %s", i, opMap[tag]))
			default:
				colums = append(colums, fmt.Sprintf("%s %s $%d", i, opMap[tag], startIndex+count))
				args = append(args, val)
			}
			count++
		}
	}
	if len(colums) == 0 {
		return "", nil
	}
	return fmt.Sprintf(whereTemplate, strings.Join(colums, " AND ")), args
}
