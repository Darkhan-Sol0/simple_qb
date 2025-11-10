package simple_qb

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	Insert = "INSERT INTO %s (%s) VALUES (%s)"
	Select = "SELECT %s FROM %s"
	Where  = "WHERE %s"
	Update = "UPDATE %s SET (%s)"
)

func QueryGenerate(query, table string, data, dataWhere any) (string, []any) {
	switch query {
	case Insert:
		col, args := getColums(data)
		placeholders := getPlaceholders(len(col))
		return fmt.Sprintf(Insert,
			table,
			strings.Join(col, ", "),
			strings.Join(placeholders, ", ")), args
	case Select:
		col, args := getColums(data)
		where := ""
		query := fmt.Sprintf(Select,
			strings.Join(col, ", "),
			table)
		if dataWhere != nil {
			where, args = getWhere(dataWhere, 0)
			query = fmt.Sprintf("%s %s", query, where)
		}
		return query, args
	case Update:
		col, args := getColumsUpdate(data)
		query := fmt.Sprintf(Update,
			table,
			strings.Join(col, ", "),
		)
		if dataWhere != nil {
			where, args2 := getWhere(dataWhere, len(args))
			query = fmt.Sprintf("%s %s", query, where)
			args = append(args, args2...)
		}
		return query, args
	}
	return "", nil
}

func getColums(data any) (colums []string, args []any) {
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get("db")
		val := v.Field(i)
		if dbTag != "" && dbTag != "-" {
			colums = append(colums, dbTag)
			args = append(args, val.Interface())
		}
	}
	return colums, args
}

func getPlaceholders(count int) []string {
	placeholders := make([]string, count)
	for i := range count {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	return placeholders
}

// Пока условие на равно значению поля. Нужно будет сделать более гибкую функцию
func getWhere(data any, startIndex int) (query string, args []any) {
	var colums []string
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get("db")
		val := v.Field(i)
		if dbTag != "" && dbTag != "-" && val.IsNil() {
			colums = append(colums, fmt.Sprintf("%s = $%d", dbTag, startIndex+i+1))
			args = append(args, val.Interface())
		}
	}
	return fmt.Sprintf(Where, strings.Join(colums, " AND ")), args
}

func getColumsUpdate(data any) (colums []string, args []any) {
	v := reflect.ValueOf(data)
	t := v.Type()

	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get("db")
		val := v.Field(i)
		if dbTag != "" && dbTag != "-" && val.IsNil() {
			colums = append(colums, fmt.Sprintf("%s = $%d", dbTag, i+1))
			args = append(args, val.Interface())
		}
	}
	return colums, args
}
