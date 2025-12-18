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

type count struct {
	column string
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
		SelectGenerate() string
		InsertGenerate() (string, []any)
		UpdateGenerate() (string, []any)
		DeleteGenerate() string
	}
)

func Count(data string) count {
	return count{
		column: data,
	}
}

func New(method, table string, data any) Query {
	return &query{
		method:   method,
		template: methodTemplate[method],
		table:    table,
		data:     data,
	}
}

func (q *query) SelectGenerate() string {
	return q.selectBuild()
}

func (q *query) InsertGenerate() (string, []any) {
	return q.insertBuild(), getArguments(q.data)
}

func (q *query) UpdateGenerate() (string, []any) {
	return q.updateBuild(), getArguments(q.data)
}

func (q *query) DeleteGenerate() string {
	return q.deleteBuild()
}

// ------------

func (q *query) selectBuild() string {
	if data, ok := q.data.(count); ok {
		if data.column == "" {
			return fmt.Sprintf(q.template, "COUNT(*)", q.table)
		}
		return fmt.Sprintf(q.template, fmt.Sprintf("COUNT(%s)", data.column), q.table)
	}
	colums := getColumns(q.data)
	return fmt.Sprintf(q.template, strings.Join(colums, ", "), q.table)
}

func (q *query) insertBuild() string {
	colums := getColumns(q.data)
	placeholders := getPlaceholders(len(colums))
	return fmt.Sprintf(q.template, q.table, strings.Join(colums, ", "), strings.Join(placeholders, ", "))
}

func (q *query) updateBuild() string {
	colums := getColumns(q.data)
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

func getColumns(data any) (columns []string) {
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get(tag)
		if dbTag != "" && dbTag != "-" {
			columns = append(columns, dbTag)
		}
	}
	return columns
}

func getPlaceholders(count int) (placeholders []string) {
	for i := 1; i <= count; i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
	}
	return placeholders
}
