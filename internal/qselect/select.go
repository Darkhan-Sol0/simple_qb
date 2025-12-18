package qselect

import (
	"fmt"

	"github.com/Darkhan-Sol0/simple_qb/internal/params"
	"github.com/Darkhan-Sol0/simple_qb/internal/query"
)

type (
	qSelect struct {
		query  query.Query
		params params.Params
		limit  string
		order  string
	}

	Select interface {
		Params(nodes ...params.Node) Select
		Limit(limit, offset int) Select
		OrderBy(column, order string) Select
		Generate() (string, []any)
	}
)

func New(tableName string, data any) Select {
	return &qSelect{
		query: query.New("SELECT", tableName, data),
	}
}

func (s *qSelect) Params(nodes ...params.Node) Select {
	if nodes != nil {
		s.params = params.New(nodes...)
	}
	return s
}

func (s *qSelect) Limit(limit, offset int) Select {
	if limit > 0 {
		s.limit = fmt.Sprintf("LIMIT %d", limit)
		if offset > 0 {
			s.limit = fmt.Sprintf("%s OFFSET %d", s.limit, offset)
		}
	}
	return s
}

func (s *qSelect) OrderBy(column, order string) Select {
	if column != "" {
		if order != "ASC" && order != "DESC" {
			order = "ASC"
		}
		s.order = fmt.Sprintf("ORDER BY %s %s", column, order)
	}
	return s
}

func (s *qSelect) Generate() (q string, args []any) {
	q = s.query.SelectGenerate()
	if s.params != nil {
		w, arg := s.params.Generate(0)
		q = fmt.Sprintf("%s %s", q, w)
		args = append(args, arg...)
	}

	if s.order != "" {
		q = fmt.Sprintf("%s %s", q, s.order)
	}

	if s.limit != "" {
		q = fmt.Sprintf("%s %s", q, s.limit)
	}

	return q, args
}
