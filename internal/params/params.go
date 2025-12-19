package params

import (
	"fmt"
)

var whereTemplate = "WHERE %s"

type (
	params struct {
		nodes []Node
		query string
	}

	Params interface {
		Generate(startIndex int) (query string, args []any)
		Or(node Node) Params
		And(node Node) Params
	}
)

// --------------------------------

func New(node Node) Params {
	return &params{
		nodes: []Node{node},
		query: whereTemplate,
	}
}

func (p *params) Or(node Node) Params {
	if node != nil {
		p.nodes = append(p.nodes, node)
		p.query = fmt.Sprint(p.query, " OR %s")
	}
	return p
}

func (p *params) And(node Node) Params {
	if node != nil {
		p.nodes = append(p.nodes, node)
		p.query = fmt.Sprint(p.query, " AND %s")
	}
	return p
}

func (p *params) Generate(startIndex int) (query string, args []any) {
	if len(p.nodes) == 0 {
		return "", nil
	}
	if len(p.nodes) == 1 {
		s := fmt.Sprintf(whereTemplate, p.nodes[0].query(startIndex+len(args)))
		args = append(args, p.nodes[0].args()...)
		return s, args
	}
	var s []any
	if len(p.nodes) > 1 {
		for _, i := range p.nodes {
			s = append(s, fmt.Sprintf("(%s)", i.query(startIndex+len(args))))
			args = append(args, i.args()...)
		}
	}
	return fmt.Sprintf(p.query, s...), args
}
