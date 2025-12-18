package simple_qb

import (
	"github.com/Darkhan-Sol0/simple_qb/internal/params"
	"github.com/Darkhan-Sol0/simple_qb/internal/qdelete"
	"github.com/Darkhan-Sol0/simple_qb/internal/qinsert"
	"github.com/Darkhan-Sol0/simple_qb/internal/qselect"
	"github.com/Darkhan-Sol0/simple_qb/internal/qupdate"
)

type (
	qBuilder struct {
		tableName string
	}

	QBuilder interface {
		Select(data any) qselect.Select
		Insert(data any) qinsert.Insert
		Update(data any) qupdate.Update
		Delete(data any) qdelete.Delete
	}
)

func NewParam(column string) params.Node {
	return params.NewNode(column)
}

func New(tableName string) QBuilder {
	return &qBuilder{
		tableName: tableName,
	}
}
func (q *qBuilder) Select(data any) qselect.Select {
	return qselect.New(q.tableName, data)
}

func (q *qBuilder) Insert(data any) qinsert.Insert {
	return qinsert.New(q.tableName, data)
}

func (q *qBuilder) Update(data any) qupdate.Update {
	return qupdate.New(q.tableName, data)
}

func (q *qBuilder) Delete(data any) qdelete.Delete {
	return qdelete.New(q.tableName, data)
}
