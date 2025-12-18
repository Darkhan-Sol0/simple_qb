package qinsert

import (
	"fmt"

	"github.com/Darkhan-Sol0/simple_qb/internal/params"
	"github.com/Darkhan-Sol0/simple_qb/internal/query"
)

type (
	qInsert struct {
		query     query.Query
		params    params.Params
		returning string
	}

	Insert interface {
		Params(nodes ...params.Node) Insert
		Returning(column string) Insert
		Generate() (string, []any)
	}
)

func New(tableName string, data any) Insert {
	return &qInsert{
		query: query.New("INSERT", tableName, data),
	}
}

func (s *qInsert) Params(nodes ...params.Node) Insert {
	if nodes != nil {
		s.params = params.New(nodes...)
	}
	return s
}

func (s *qInsert) Returning(column string) Insert {
	if column != "" {
		s.returning = fmt.Sprintf("RETURNING %s", column)
	}
	return s
}

func (s *qInsert) Generate() (string, []any) {
	q, args := s.query.InsertGenerate()
	if s.params != nil {
		w, arg := s.params.Generate(len(args))

		q = fmt.Sprintf("%s %s", q, w)
		args = append(args, arg...)
	}
	if s.returning != "" {
		q = fmt.Sprintf("%s %s", q, s.returning)
	}
	return q, args
}
