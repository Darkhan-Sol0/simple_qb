package simple_qb

import (
	"fmt"
	"testing"

	"github.com/Darkhan-Sol0/simple_qb/internal/query"
)

type S struct {
	Num  int    `db:"num"`
	Text string `db:"text"`
}

var table = "users"

func TestQueryInsert(t *testing.T) {
	check := "INSERT INTO users (num, text) VALUES ($1, $2)"
	s, arg := New(table).Insert(S{Num: 1, Text: "hui"}).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQueryInsert1(t *testing.T) {
	check := "INSERT INTO users (num, text) VALUES ($1, $2)"
	s, arg := New(table).Insert(S{Num: 1}).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect(t *testing.T) {
	check := fmt.Sprintf("SELECT num, text FROM %s", table)
	s, arg := New(table).Select(S{}).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect1(t *testing.T) {
	check := fmt.Sprintf("SELECT num, text FROM %s WHERE num IN ($1, $2, $3)", table)
	s, arg := New(table).Select(S{}).Params(NewParam(NewNode("num").In([]any{"1", "sad", "da11"}))).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect2(t *testing.T) {
	check := fmt.Sprintf("SELECT num, text FROM %s WHERE num IN ($1, $2, $3) OR num IS NULL", table)
	s, arg := New(table).Select(S{}).Params(NewParam(NewNode("num").In([]any{"1", "sad", "da11"}).Or().Null())).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect3(t *testing.T) {
	check := fmt.Sprintf("SELECT num, text FROM %s WHERE num = $1", table)
	s, arg := New(table).Select(S{}).Params(NewParam(NewNode("num").Eq(1))).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect4(t *testing.T) {
	check := fmt.Sprintf("SELECT COUNT(num) FROM %s WHERE num = $1 OR num < $2", table)
	s, arg := New(table).Select(query.Count("num")).Params(NewParam(NewNode("num").Eq(1).Or().Less(5))).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect5(t *testing.T) {
	check := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
	s, arg := New(table).Select(query.Count("")).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect6(t *testing.T) {
	check := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE num BETWEEN $1 AND $2", table)
	s, arg := New(table).Select(query.Count("")).Params(NewParam(NewNode("num").Between(123, 134))).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQueryUpdate(t *testing.T) {
	check := fmt.Sprintf("UPDATE %s SET (num, text) = ROW($1, $2) WHERE (num = $3 OR num < $4) AND (text = $5 AND text LIKE $6)", table)
	s, arg := New(table).Update(S{Num: 5, Text: "qwe"}).Params(NewParam(NewNode("num").Eq(123).Or().Less(5)).And(NewNode("text").Eq("asd").And().Like("asd"))).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}
