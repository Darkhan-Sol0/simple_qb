package params

import (
	"fmt"
	"reflect"
	"strings"
)

var whereTemplate = "WHERE %s"

var opMap = map[string]string{
	"eq":      "=",           // равно
	"neq":     "<>",          // неравно
	"lt":      "<",           // меньше
	"lte":     "<=",          // меньше или равно
	"gt":      ">",           // больше
	"gte":     ">=",          // больше или равно
	"like":    "LIKE",        // похоже на (для строковых выражений)
	"in":      "IN",          // входит в перечень (depricatet)
	"null":    "IS NULL",     // пустое значение
	"notnull": "IS NOT NULL", // непустое значение
}

type (
	Node struct {
		column   string
		operator string
		value    any
		logic    string
	}

	params struct {
		nodes []*Node
	}

	Params interface {
		Generate(startIndex int) (query string, args []any)
	}
)

func NewNode(column, operator string, value any) *Node {
	return &Node{
		column:   column,
		operator: operator,
		value:    value,
		logic:    "AND",
	}
}

func NewNodeOr(column, operator string, value any) *Node {
	return &Node{
		column:   column,
		operator: operator,
		value:    value,
		logic:    "OR",
	}
}

func New(nodes ...*Node) Params {
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
	colums := make([]string, 2)
	count := 1
	for _, v := range p.nodes {
		logic := v.logic
		operator := v.operator
		column := v.column
		val := v.value
		if operator != "" && opMap[operator] != "" {
			switch operator {
			case "in":
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
					colums[1] = fmt.Sprintf("%s %s (%s)", column, opMap[operator], strings.Join(placeholders, ", "))
				}
			case "null", "notnull":
				colums[1] = fmt.Sprintf("%s %s", column, opMap[operator])

			default:
				colums[1] = fmt.Sprintf("%s %s $%d", column, opMap[operator], startIndex+count)
				args = append(args, val)
				count++
			}
			if colums[0] != "" {
				text := strings.Join(colums, fmt.Sprintf(" %s ", logic))
				colums[0] = text
			} else {
				colums[0] = colums[1]
			}
			colums[1] = ""
		}
	}
	if len(colums) == 0 {
		return "", nil
	}
	return fmt.Sprintf(whereTemplate, colums[0]), args
}
