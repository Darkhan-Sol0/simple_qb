package query

import (
	"fmt"
	"reflect"
	"strings"
)

var methodTemplate = map[string]string{
	"SELECT": "SELECT %s FROM %s",
	"INSERT": "INSERT INTO %s (%s) VALUES (%s)",
	"UPDATE": "UPDATE %s SET (%s) = ROW(%s)",
	"DELETE": "DELETE FROM %s",
}

const tag = "db"

type (
	query struct {
		method   string
		template string
		table    string
		data     any
	}

	Query interface {
		Generate() (string, []any)
	}
)

func New(method, table string, data any) Query {
	return &query{
		method:   method,
		template: methodTemplate[method],
		table:    table,
		data:     data,
	}
}

func (q *query) Generate() (string, []any) {
	switch q.method {
	case "SELECT":
		return q.selectBuild(), nil
	case "INSERT":
		return q.insertBuild(), getArguments(q.data)
	case "UPDATE":
		return q.updateBuild(), getArguments(q.data)
	case "DELETE":
		return q.deleteBuild(), nil
	}
	return "", nil
}

// ------------

func (q *query) selectBuild() string {
	colums := getColums(q.data)
	return fmt.Sprintf(q.template, strings.Join(colums, ", "), q.table)
}

func (q *query) insertBuild() string {
	colums := getColums(q.data)
	placeholders := getPlaceholders(len(colums))
	return fmt.Sprintf(q.template, q.table, strings.Join(colums, ", "), strings.Join(placeholders, ", "))
}

func (q *query) updateBuild() string {
	colums := getColums(q.data)
	placeholders := getPlaceholders(len(colums))
	return fmt.Sprintf(q.template, q.table, strings.Join(colums, ", "), strings.Join(placeholders, ", "))
}

func (q *query) deleteBuild() string {
	return fmt.Sprintf(q.template, q.table)
}

// ---------------

func getArguments(data any) (args []any) {
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get(tag)
		val := v.Field(i)
		if dbTag != "" && dbTag != "-" && !val.IsZero() {
			args = append(args, val.Interface())
		}
	}
	return args
}

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

func getPlaceholders(count int) (placeholders []string) {
	for i := 1; i <= count; i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
	}
	return placeholders
}
