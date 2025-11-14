package simple_qb

import (
	"fmt"
	"testing"
)

type S struct {
	Num  int    `db:"num"`
	Text string `db:"text"`
}

type F struct {
	Num int `db:"num" op:"eq"`
}

var table = "users"

func TestQueryInsert(t *testing.T) {
	check := fmt.Sprintf(Insert, table, "num, text", "$1, $2")
	s, arg := QueryGenerate(Insert, table, S{Num: 123, Text: "qwe"}, nil)
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQuerySelect(t *testing.T) {
	check := fmt.Sprintf(Select, "num, text", table)
	s, arg := QueryGenerate(Select, table, S{}, nil)
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQuerySelectWhere(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(Select, "num, text", table), fmt.Sprintf(Where, "num = $1"))
	s, arg := QueryGenerate(Select, table, S{}, F{Num: 123})
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(Update, table, "num, text", "$1, $2"), fmt.Sprintf(Where, "num = $3"))
	s, arg := QueryGenerate(Update, table, S{Num: 123, Text: "qwe"}, F{Num: 123})
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate2(t *testing.T) {
	type F struct {
		Num int `db:"num" op:"gt"`
	}
	check := fmt.Sprintf("%s %s", fmt.Sprintf(Update, table, "num, text", "$1, $2"), fmt.Sprintf(Where, "num > $3"))
	s, arg := QueryGenerate(Update, table, S{Num: 123, Text: "qwe"}, F{Num: 123})
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}
