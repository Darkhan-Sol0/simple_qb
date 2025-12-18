package qdelete

import (
	"fmt"

	"github.com/Darkhan-Sol0/simple_qb/internal/params"
	"github.com/Darkhan-Sol0/simple_qb/internal/query"
)

type (
	qDelete struct {
		query  query.Query
		params params.Params
	}

	Delete interface {
		Params(par params.Params) Delete
		Generate() (string, []any)
	}
)

func New(tableName string, data any) Delete {
	return &qDelete{
		query: query.New("DELETE", tableName, data),
	}
}

func (s *qDelete) Params(par params.Params) Delete {
	if par != nil {
		s.params = par
	}
	return s
}

func (s *qDelete) Generate() (string, []any) {
	if s.params == nil {
		return "", nil
	}
	q := s.query.DeleteGenerate()
	w, args := s.params.Generate(0)
	q = fmt.Sprintf("%s %s", q, w)
	return q, args
}
