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
		SelectGenerate() (string, error)
		InsertGenerate() (string, []any, error)
		UpdateGenerate() (string, []any, error)
		DeleteGenerate() (string, error)
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

func (q *query) SelectGenerate() (string, error) {
	res, err := q.selectBuild()
	if err != nil {
		return "", err
	}
	return res, nil
}

func (q *query) InsertGenerate() (string, []any, error) {
	res, err := q.insertBuild()
	if err != nil {
		return "", nil, err
	}
	return res, getArguments(q.data), nil
}

func (q *query) UpdateGenerate() (string, []any, error) {
	res, err := q.updateBuild()
	if err != nil {
		return "", nil, err
	}
	return res, getArguments(q.data), nil
}

func (q *query) DeleteGenerate() (string, error) {
	res, err := q.deleteBuild()
	if err != nil {
		return "", err
	}
	return res, nil
}

// ------------
func (q *query) selectBuild() (string, error) {
	if q.data == nil {
		return "", fmt.Errorf("data struct is empty")
	}
	if q.table == "" {
		return "", fmt.Errorf("table name is empty")
	}
	if t, ok := q.data.(string); ok {
		if t == "" {
			t = "*"
		}
		return fmt.Sprintf(q.template, fmt.Sprintf("COUNT(%s)", t), q.table), nil
	}
	colums := getColumns(q.data)
	return fmt.Sprintf(q.template, strings.Join(colums, ", "), q.table), nil
}

func (q *query) insertBuild() (string, error) {
	if q.data == nil {
		return "", fmt.Errorf("data struct is empty")
	}
	if q.table == "" {
		return "", fmt.Errorf("table name is empty")
	}
	colums := getColumns(q.data)
	placeholders := getPlaceholders(len(colums))
	return fmt.Sprintf(q.template, q.table, strings.Join(colums, ", "), strings.Join(placeholders, ", ")), nil
}

func (q *query) updateBuild() (string, error) {
	if q.data == nil {
		return "", fmt.Errorf("data struct is empty")
	}
	if q.table == "" {
		return "", fmt.Errorf("table name is empty")
	}
	colums := getColumns(q.data)
	placeholders := getPlaceholders(len(colums))
	return fmt.Sprintf(q.template, q.table, strings.Join(colums, ", "), strings.Join(placeholders, ", ")), nil
}

func (q *query) deleteBuild() (string, error) {
	if q.table == "" {
		return "", fmt.Errorf("table name is empty")
	}
	return fmt.Sprintf(q.template, q.table), nil
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
		if dbTag != "" && dbTag != "-" && !v.Field(i).IsZero() {
			columns = append(columns, dbTag)
		}
	}
	return columns
}

// if v.Kind() == reflect.Slice {
// 		l := v.Len()
// 		var st []string
// 		for i := 0; i < l; i++ {
// 			st = append(st, "$%d")
// 			elem := v.Index(i)
// 			n.nargs = append(n.nargs, elem.Interface())
// 		}
// 		n.nquery = append(n.nquery, fmt.Sprintf("%s IN (%s)", n.ncolumn, strings.Join(st, ", ")))
// 	}

func getPlaceholders(count int) (placeholders []string) {
	for i := 1; i <= count; i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
	}
	return placeholders
}
