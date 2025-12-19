package params

import (
	"fmt"
	"reflect"
	"strings"
)

type (
	node struct {
		ncolumn string
		nquery  []string
		nargs   []any
	}

	Node interface {
		query(startIndex int) string
		args() []any

		Or() Node
		And() Node

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
		Between(left, right any) Node
		NotBetween(left, right any) Node
	}
)

func (n *node) query(startIndex int) string {
	if len(n.nquery) == 0 {
		return ""
	}
	if n.nquery[len(n.nquery)-1] == "OR" || n.nquery[len(n.nquery)-1] == "AND" {
		t := []string{}
		for i := 0; i < len(n.nquery)-1; i++ {
			t = append(t, n.nquery[i])
		}
		n.nquery = t
	}
	q := strings.Join(n.nquery, " ")
	var a []any
	for i := range n.nargs {
		a = append(a, i+1+startIndex)
	}
	return fmt.Sprintf(q, a...)
}
func (n *node) args() []any {
	return n.nargs
}

func NewNode(column string) Node {
	return &node{
		ncolumn: column,
		nquery:  nil,
	}
}

func (n *node) Or() Node {
	n.nquery = append(n.nquery, "OR")
	return n
}

func (n *node) And() Node {
	n.nquery = append(n.nquery, "AND")
	return n
}

func (n *node) Eq(value any) Node {
	n.nargs = append(n.nargs, value)
	n.nquery = append(n.nquery, fmt.Sprintf("%s = %s", n.ncolumn, "$%d"))
	return n
}

func (n *node) NotEq(value any) Node {
	n.nargs = append(n.nargs, value)
	n.nquery = append(n.nquery, fmt.Sprintf("%s <> %s", n.ncolumn, "$%d"))
	return n
}

func (n *node) Less(value any) Node {
	n.nargs = append(n.nargs, value)
	n.nquery = append(n.nquery, fmt.Sprintf("%s < %s", n.ncolumn, "$%d"))
	return n
}

func (n *node) LessEq(value any) Node {
	n.nargs = append(n.nargs, value)
	n.nquery = append(n.nquery, fmt.Sprintf("%s <= %s", n.ncolumn, "$%d"))
	return n
}

func (n *node) Gr(value any) Node {
	n.nargs = append(n.nargs, value)
	n.nquery = append(n.nquery, fmt.Sprintf("%s > %s", n.ncolumn, "$%d"))
	return n
}

func (n *node) GrEq(value any) Node {
	n.nargs = append(n.nargs, value)
	n.nquery = append(n.nquery, fmt.Sprintf("%s >= %s", n.ncolumn, "$%d"))
	return n
}

func (n *node) Like(value any) Node {
	n.nargs = append(n.nargs, value)
	n.nquery = append(n.nquery, fmt.Sprintf("%s LIKE %s", n.ncolumn, "$%d"))
	return n
}

func (n *node) In(value any) Node {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Slice {
		l := v.Len()
		var st []string
		for i := 0; i < l; i++ {
			st = append(st, "$%d")
			elem := v.Index(i)
			n.nargs = append(n.nargs, elem.Interface())
		}
		n.nquery = append(n.nquery, fmt.Sprintf("%s IN (%s)", n.ncolumn, strings.Join(st, ", ")))
	}
	return n
}

func (n *node) Null() Node {
	n.nquery = append(n.nquery, fmt.Sprintf("%s IS NULL", n.ncolumn))
	return n
}

func (n *node) NotNull() Node {
	n.nquery = append(n.nquery, fmt.Sprintf("%s IS NOT NULL", n.ncolumn))
	return n
}

func (n *node) Between(left, right any) Node {
	n.nargs = append(n.nargs, left, right)
	n.nquery = append(n.nquery, fmt.Sprintf("%s BETWEEN %s AND %s", n.ncolumn, "$%d", "$%d"))
	return n
}
func (n *node) NotBetween(left, right any) Node {
	n.nargs = append(n.nargs, left, right)
	n.nquery = append(n.nquery, fmt.Sprintf("%s NOT BETWEEN %s AND %s", n.ncolumn, "$%d", "$%d"))
	return n
}
