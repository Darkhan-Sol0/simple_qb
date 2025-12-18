package params

import (
	"fmt"
	"reflect"
	"strings"
)

var whereTemplate = "WHERE %s"

type (
	node struct {
		ncolumn   string
		noperator string
		nvalue    any
		nlogic    string
	}

	Node interface {
		column() string
		operator() string
		value() any
		logic() string

		Or() Node
		Eq(value any) Node
		NotEq(value any) Node
		Less(value any) Node
		LessEq(value any) Node
		Gr(value any) Node
		GrEq(value any) Node
		Like(value any) Node
		In(value any) Node
		Null() Node
		NotNull() Node
	}

	params struct {
		nodes []Node
	}

	Params interface {
		Generate(startIndex int) (query string, args []any)
	}
)

func NewNode(column string) Node {
	return &node{
		ncolumn:   column,
		noperator: "=",
		nvalue:    nil,
		nlogic:    "AND",
	}
}

func (n *node) Eq(value any) Node {
	n.noperator = "="
	n.nvalue = value
	return n
}

func (n *node) NotEq(value any) Node {
	n.noperator = "<>"
	n.nvalue = value
	return n
}

func (n *node) Less(value any) Node {
	n.noperator = "<"
	n.nvalue = value
	return n
}

func (n *node) LessEq(value any) Node {
	n.noperator = "<="
	n.nvalue = value
	return n
}

func (n *node) Gr(value any) Node {
	n.noperator = ">"
	n.nvalue = value
	return n
}

func (n *node) GrEq(value any) Node {
	n.noperator = ">="
	n.nvalue = value
	return n
}

func (n *node) Like(value any) Node {
	n.noperator = "LIKE"
	n.nvalue = value
	return n
}

func (n *node) In(value any) Node {
	n.noperator = "IN"
	n.nvalue = value
	return n
}

func (n *node) Null() Node {
	n.noperator = "IS NULL"
	return n
}

func (n *node) NotNull() Node {
	n.noperator = "IS NOT NULL"
	return n
}

func (n *node) Or() Node {
	n.nlogic = "OR"
	return n
}

func (n *node) column() string {
	return n.ncolumn
}

func (n *node) operator() string {
	return n.noperator
}

func (n *node) value() any {
	return n.nvalue
}

func (n *node) logic() string {
	return n.nlogic
}

// --------------------------------

func New(nodes ...Node) Params {
	return &params{
		nodes: nodes,
	}
}

func (p *params) Generate(startIndex int) (query string, args []any) {
	return p.getWhere(startIndex)
}

func (p *params) getWhere(startIndex int) (query string, args []any) {
	if len(p.nodes) == 0 {
		return "", nil
	}
	columns := make([]string, 2)
	count := 1
	for _, v := range p.nodes {
		logic := v.logic()
		operator := v.operator()
		column := v.column()
		val := v.value()
		if operator != "" {
			switch operator {
			case "IN":
				v := reflect.ValueOf(val)
				if v.Kind() == reflect.Slice {
					l := v.Len()
					placeholders := make([]string, l)
					for i := 0; i < l; i++ {
						elem := v.Index(i)
						placeholders[i] = fmt.Sprintf("$%d", startIndex+count)
						args = append(args, elem.Interface())
						count++
					}
					columns[1] = fmt.Sprintf("%s %s (%s)", column, operator, strings.Join(placeholders, ", "))
				}
			case "IS NULL", "IS NOT NULL":
				columns[1] = fmt.Sprintf("%s %s", column, operator)

			default:
				columns[1] = fmt.Sprintf("%s %s $%d", column, operator, startIndex+count)
				args = append(args, val)
				count++
			}
			if columns[0] != "" {
				text := strings.Join(columns, fmt.Sprintf(" %s ", logic))
				columns[0] = text
			} else {
				columns[0] = columns[1]
			}
			columns[1] = ""
		}
	}
	if len(columns) == 0 {
		return "", nil
	}
	return fmt.Sprintf(whereTemplate, columns[0]), args
}
