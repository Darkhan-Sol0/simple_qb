package simple_qb

import (
	"fmt"
	"reflect"
	"strings"
)

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
	return placeholders
}

// Нужно както впиндюрить оператор OR

// getWhere builds WHERE condition string based on structured data and special tags.
// Returns formatted WHERE clause combining columns and operators via AND.

func getWhere(data FilterNode, startIndex int) (query string, args []any) {
	var colums []string
	count := 1
	for _, v := range data {
		tag := v.Operator
		val := v.Value
		if tag != "" && opMap[tag] != "" {
			switch tag {
			case "in":
				colums = append(colums, fmt.Sprintf("%s %s ($%d)", v.Tag, opMap[tag], startIndex+count))
				args = append(args, val)
			case "null", "notnull":
				colums = append(colums, fmt.Sprintf("%s %s", v.Tag, opMap[tag]))
			default:
				colums = append(colums, fmt.Sprintf("%s %s $%d", v.Tag, opMap[tag], startIndex+count))
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
