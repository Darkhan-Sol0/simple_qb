package qupdate

import (
	"fmt"

	"github.com/Darkhan-Sol0/simple_qb/internal/params"
	"github.com/Darkhan-Sol0/simple_qb/internal/query"
)

type (
	qUpdate struct {
		query  query.Query
		params params.Params
	}

	Update interface {
		Params(nodes ...params.Node) Update
		Generate() (string, []any)
	}
)

func New(tableName string, data any) Update {
	return &qUpdate{
		query: query.New("UPDATE", tableName, data),
	}
}

func (s *qUpdate) Params(nodes ...params.Node) Update {
	if nodes != nil {
		s.params = params.New(nodes...)
	}
	return s
}

func (s *qUpdate) Generate() (string, []any) {
	if s.params == nil {
		return "", nil
	}
	q, args := s.query.UpdateGenerate()
	w, arg := s.params.Generate(len(args))
	q = fmt.Sprintf("%s %s", q, w)
	args = append(args, arg...)
	return q, args
}
