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
		Params(par params.Params) Update
		Generate() (string, []any, error)
	}
)

func New(tableName string, data any) Update {
	return &qUpdate{
		query: query.New("UPDATE", tableName, data),
	}
}

func (s *qUpdate) Params(par params.Params) Update {
	if par != nil {
		s.params = par
	}
	return s
}

func (s *qUpdate) Generate() (string, []any, error) {
	if s.params == nil {
		return "", nil, fmt.Errorf("empty params where")
	}
	q, args, err := s.query.UpdateGenerate()
	if err != nil {
		return "", nil, err
	}
	w, arg := s.params.Generate(len(args))
	q = fmt.Sprintf("%s %s", q, w)
	args = append(args, arg...)
	return q, args, nil
}
