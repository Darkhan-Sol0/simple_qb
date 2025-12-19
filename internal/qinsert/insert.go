package qinsert

import (
	"fmt"

	"github.com/Darkhan-Sol0/simple_qb/internal/query"
)

type (
	qInsert struct {
		query     query.Query
		returning string
	}

	Insert interface {
		Returning(column string) Insert
		Generate() (string, []any, error)
	}
)

func New(tableName string, data any) Insert {
	return &qInsert{
		query: query.New("INSERT", tableName, data),
	}
}

func (s *qInsert) Returning(column string) Insert {
	if column != "" {
		s.returning = fmt.Sprintf("RETURNING %s", column)
	}
	return s
}

func (s *qInsert) Generate() (string, []any, error) {
	q, args, err := s.query.InsertGenerate()
	if err != nil {
		return "", nil, err
	}

	if s.returning != "" {
		q = fmt.Sprintf("%s %s", q, s.returning)
	}
	return q, args, nil
}
