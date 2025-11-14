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
	Insert = "INSERT INTO %s (%s) VALUES (%s)"
	Select = "SELECT %s FROM %s"
	Where  = "WHERE %s"
	Update = "UPDATE %s SET (%s) = (%s)"
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

// QueryGenerate creates an SQL query from given parameters.
// Arguments:
//
//	qtype: Type of query being constructed (INSERT, SELECT, UPDATE).
//	table: Target database table.
//	data: Structured data for insertion or update.
//	params: Filtering parameters for WHERE clause.
//
// Returns:
//
//	Formatted SQL query and array of arguments for execution.
func QueryGenerate(qtype, table string, data, params any) (query string, args []any) {
	// Logic constructs SQL query by analyzing structure fields and forming appropriate sections.
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

// queryBuild joins base elements of an SQL query (template, table, columns, placeholders).
// Argument qtype determines the query type.
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
	return fmt.Sprintf(Where, strings.Join(colums, " AND "))
}
