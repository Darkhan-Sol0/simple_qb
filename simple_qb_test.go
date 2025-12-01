package simple_qb

import (
	"fmt"
	"testing"
)

type S struct {
	Num  int    `db:"num"`
	Text string `db:"text"`
}

var table = "users"

func TestQueryInsert(t *testing.T) {
	check := fmt.Sprintf("SELECT num, text FROM %s", table)
	s, arg := New(table).Select(S{}).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect(t *testing.T) {
	check := fmt.Sprintf("SELECT num, text FROM %s WHERE num IN ($1, $2, $3)", table)
	s, arg := New(table).Select(S{}).Params(NewParam("num", "in", []string{"1", "sad", "da11"})).Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect1(t *testing.T) {
	check := fmt.Sprintf("SELECT num, text FROM %s WHERE num IN ($1, $2, $3) RETURNING num", table)
	s, arg := New(table).Select(S{}).Params(NewParam("num", "in", []string{"1", "sad", "da11"})).Returning("num").Generate()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}
