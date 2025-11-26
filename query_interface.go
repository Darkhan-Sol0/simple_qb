// Package simple_qb provides a lightweight toolkit for building SQL queries for standard CRUD operations.
// It supports creating INSERT, SELECT, and UPDATE statements by processing structured data and customizable parameters.
// Built-in methods allow automatic generation of SQL queries with support for WHERE clauses and prepared expressions.
package simple_qb

import (
	"errors"
	"fmt"
	"strings"
)

type QBilder interface {
	Select() (query string, args []any)
	Insert() (query string, args []any)
	Update() (query string, args []any, err error)
	Delete() (query string, args []any, err error)
}

func New(table string, data any, params FilterNode) QBilder {
	return &qBilder{
		table:  table,
		data:   data,
		params: params,
	}
}

func NewFillter(node ...*Node) (nodes FilterNode) {
	nodes = append(nodes, node...)
	return nodes
}

func AddFillter(nodes FilterNode, node ...*Node) FilterNode {
	nodes = append(nodes, node...)
	return nodes
}

func NewNode(tag, operator, logic string, value any) *Node {
	return &Node{
		Tag:      tag,
		Operator: operator,
		Value:    value,
		Logic:    logic,
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
