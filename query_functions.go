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
		if dbTag != "" && dbTag != "-" && !val.IsZero() {
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
	if len(data) == 0 {
		return "", nil
	}
	var colums []string
	count := 1
	for i, v := range data {
		logic := v.Logic
		tag := v.Operator
		val := v.Value
		if tag != "" && opMap[tag] != "" {
			switch tag {
			case "in":
				text := fmt.Sprintf("%s %s ($%d)", v.Tag, opMap[tag], startIndex+count)
				if i > 0 {
					if logic == "" {
						text = fmt.Sprintf("%s %s", "AND", text)
					} else {
						text = fmt.Sprintf("%s %s", logic, text)
					}
				}
				colums = append(colums, text)
				args = append(args, val)
				count++
			case "null", "notnull":
				text := fmt.Sprintf("%s %s", v.Tag, opMap[tag])
				if i > 0 {
					if logic == "" {
						text = fmt.Sprintf("%s %s", "AND", text)
					} else {
						text = fmt.Sprintf("%s %s", logic, text)
					}
				}
				colums = append(colums, text)
			default:
				text := fmt.Sprintf("%s %s $%d", v.Tag, opMap[tag], startIndex+count)
				if i > 0 {
					if logic == "" {
						text = fmt.Sprintf("%s %s", "AND", text)
					} else {
						text = fmt.Sprintf("%s %s", logic, text)
					}
				}
				colums = append(colums, text)
				args = append(args, val)
				count++
			}
		}
	}
	if len(colums) == 0 {
		return "", nil
	}

	return fmt.Sprintf(whereTemplate, strings.Join(colums, " ")), args
}
