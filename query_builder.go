// Package simple_qb provides a lightweight toolkit for building SQL queries for standard CRUD operations.
// It supports creating INSERT, SELECT, and UPDATE statements by processing structured data and customizable parameters.
// Built-in methods allow automatic generation of SQL queries with support for WHERE clauses and prepared expressions.
package simple_qb

import (
	"fmt"
	"reflect"
	"strings"
)

// Insert is a template for constructing INSERT INTO commands.
// Select is a template for constructing SELECT commands.
// Where is a template for constructing WHERE clauses.
// Update is a template for constructing UPDATE commands.
var (
	insertTemplate = "INSERT INTO %s (%s) VALUES (%s)"
	selectTemplate = "SELECT %s FROM %s"
	updateTemplate = "UPDATE %s SET (%s) = (%s)"
	whereTemplate  = "WHERE %s"
)

// Tag identifies column names within structures.
// Op specifies the operation applied to data in WHERE conditions.
const (
	tag = "db"
	op  = "op"
)

// пока такие теги, может посже изменить

// Operation mappings convert internal representations into corresponding SQL operators.
var opMap = map[string]string{
	"eq":  "=",  // равно
	"neq": "<>", // неравно
	"lt":  "<",  // меньше
	"lte": "<=", // меньше или равно
	"gt":  ">",  // больше
	"gte": ">=", // больше или равно
	// "like":    "LIKE",        // похоже на (для строковых выражений)
	// "in":      "IN",          // входит в перечень
	// "null":    "IS NULL",     // пустое значение
	// "notnull": "IS NOT NULL", // непустое значение
}

type (
	qBilder struct {
		table  string
		data   any
		params any
	}

	QBilder interface {
		Select() (query string, args []any)
		Insert() (query string, args []any)
		Update() (query string, args []any)
	}
)

func New(table string, data, params any) QBilder {
	return &qBilder{
		table:  table,
		data:   data,
		params: params,
	}
}

func (q *qBilder) Select() (query string, args []any) {
	colums := getColums(q.data)
	query = fmt.Sprintf(selectTemplate, strings.Join(colums, ", "), q.table)
	if q.params != nil {
		where := getWhere(q.params, 0)
		args = append(args, getArguments(q.params)...)
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

func (q *qBilder) Update() (query string, args []any) {
	colums := getColums(q.data)
	placeholders := getPlaceholders(len(colums))
	args = getArguments(q.data)
	query = fmt.Sprintf(updateTemplate, q.table, strings.Join(colums, ", "), strings.Join(placeholders, ", "))
	if q.params != nil {
		where := getWhere(q.params, len(colums))
		args = append(args, getArguments(q.params)...)
		query = fmt.Sprintf("%s %s", query, where)
	}
	return query, args
}

// getArguments extracts arguments from structured data, considering special tags (db).
// Returns slice of interfaces containing actual values for insertion or updating.
func getArguments(data any) (args []any) {
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get(tag)
		val := v.Field(i)
		if dbTag != "" && dbTag != "-" {
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
func getWhere(data any, startIndex int) (query string) {
	var colums []string
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get(tag)
		opTag := t.Field(i).Tag.Get(op)
		if dbTag != "" && dbTag != "-" && opTag != "" && opMap[opTag] != "" {
			colums = append(colums, fmt.Sprintf("%s %s $%d", dbTag, opMap[opTag], startIndex+i+1))
		}
	}
	return fmt.Sprintf(whereTemplate, strings.Join(colums, " AND "))
}
