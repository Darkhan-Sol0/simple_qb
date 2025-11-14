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
	Update = "UPDATE %s SET (%s) = (%s)"
)

const tag = "db"

func QueryGenerate(qtype, table string, data, params any) (query string, args []any) {
	var colums []string
	var placeholders []string
	switch qtype {
	case Insert:
		colums = getColums(data)
		placeholders = getPlaceholders(len(colums))
		args = getArguments(data)
		query = queryBuild(qtype, table, colums, placeholders)
	case Select:
		colums = getColums(data)
		query = queryBuild(qtype, table, colums, placeholders)
		if params != nil {
			where := getWhere(params, 0)
			args = append(args, getArguments(params)...)
			query = fmt.Sprintf("%s %s", query, where)
		}
	case Update:
		colums = getColums(data)
		placeholders = getPlaceholders(len(colums))
		args = getArguments(data)
		query = queryBuild(qtype, table, colums, placeholders)
		if params != nil {
			where := getWhere(params, len(colums))
			args = append(args, getArguments(params)...)
			query = fmt.Sprintf("%s %s", query, where)
		}
	default:
	}
	return query, args
}

func queryBuild(qtype, table string, colums, placeholders []string) (query string) {
	switch qtype {
	case Insert:
		query = fmt.Sprintf(Insert, table, strings.Join(colums, ", "), strings.Join(placeholders, ", "))
	case Select:
		query = fmt.Sprintf(Select, strings.Join(colums, ", "), table)
	case Update:
		query = fmt.Sprintf(Update, table, strings.Join(colums, ", "), strings.Join(placeholders, ", "))
	}
	return
}

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
	return
}

func getWhere(data any, startIndex int) (query string) {
	var colums []string
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get("db")
		if dbTag != "" && dbTag != "-" {
			colums = append(colums, fmt.Sprintf("%s = $%d", dbTag, startIndex+i+1))
		}
	}
	return fmt.Sprintf(Where, strings.Join(colums, " AND "))
}
