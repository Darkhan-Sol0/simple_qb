package simple_qb

import (
	"fmt"

	"github.com/Darkhan-Sol0/simple_qb/internal/params"
	"github.com/Darkhan-Sol0/simple_qb/internal/query"
)

type (
	qBuilder struct {
		dbTableName string
		method      string
		query       query.Query
		params      params.Params

		limit string
		order string
	}

	QBuilder interface {
		Select(data any) QBuilder
		Insert(data any) QBuilder
		Update(data any) QBuilder
		Delete() QBuilder

		Params(nodes ...*params.Node) QBuilder
		Limit(limit, offset int) QBuilder
		OrderBy(column, order string) QBuilder

		Generate() (string, []any)
	}
)

func New(dbTable string) QBuilder {
	return &qBuilder{
		dbTableName: dbTable,
	}
}

func NewParam(column, operator string, value any) *params.Node {
	return params.NewNode(column, operator, value)
}

func NewOrParam(column, operator string, value any) *params.Node {
	return params.NewNodeOr(column, operator, value)
}

func (b *qBuilder) Select(data any) QBuilder {
	b.method = "SELECT"
	b.query = query.New(b.method, b.dbTableName, data)
	return b
}

func (b *qBuilder) Insert(data any) QBuilder {
	b.method = "INSERT"
	b.query = query.New(b.method, b.dbTableName, data)
	return b
}

func (b *qBuilder) Update(data any) QBuilder {
	b.method = "UPDATE"
	b.query = query.New(b.method, b.dbTableName, data)
	return b
}

func (b *qBuilder) Delete() QBuilder {
	b.method = "DELETE"
	b.query = query.New(b.method, b.dbTableName, nil)
	return b
}

func (b *qBuilder) Params(nodes ...*params.Node) QBuilder {
	b.params = params.New(nodes...)
	return b
}

func (b *qBuilder) Limit(limit, offset int) QBuilder {
	if limit > 0 {
		b.limit = fmt.Sprintf("LIMIT %d", limit)
		if offset > 0 {
			b.limit = fmt.Sprintf("%s OFFSET %d", b.limit, offset)
		}
	}
	return b
}

func (b *qBuilder) OrderBy(column, order string) QBuilder {
	if column != "" {
		if order != "ASC" && order != "DESC" {
			order = "ASC"
		}
		b.order = fmt.Sprintf("ORDER BY %s %s", column, order)
	}
	return b
}

func (b *qBuilder) Generate() (string, []any) {
	if (b.method == "UPDATE" || b.method == "DELETE") && b.params == nil {
		return "", nil
	}

	q, args := b.query.Generate()
	if b.params != nil {
		w, arg := b.params.Generate(len(args))

		q = fmt.Sprintf("%s %s", q, w)
		args = append(args, arg...)
	}

	if b.order != "" {
		q = fmt.Sprintf("%s %s", q, b.order)
	}

	if b.limit != "" {
		q = fmt.Sprintf("%s %s", q, b.limit)
	}

	return q, args
}

// func main() {
// 	type S struct {
// 		N int    `db:"num"`
// 		T string `db:"text"`
// 	}

// 	a, b := New("hui").Select(S{N: 123, T: "asd"}).Params(NewParam("num", "eq", 23), NewOrParam("text", "gt", "qwe")).OrderBy("num", "ASC").Limit(10, 10).Generate()

// 	fmt.Println(a, b)
// }
