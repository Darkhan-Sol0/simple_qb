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
	check := fmt.Sprintf(insertTemplate, table, "num, text", "$1, $2")
	r := New(table, S{Num: 123, Text: "qwe"}, nil, nil)
	s, arg := r.Insert()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQuerySelect(t *testing.T) {
	check := fmt.Sprintf(selectTemplate, "num, text", table)
	r := New(table, S{}, nil, nil)
	s, arg := r.Select()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQuerySelectWhere(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(selectTemplate, "num, text", table), fmt.Sprintf(whereTemplate, "num = $1"))
	r := New(table, S{}, F{Num: 123}, nil)
	s, arg := r.Select()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num = $3"))
	r := New(table, S{Num: 123, Text: "qwe"}, F{Num: 123}, nil)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate2(t *testing.T) {
	type F struct {
		Num int `db:"num" op:"gt"`
	}
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num > $3"))
	r := New(table, S{Num: 123, Text: "qwe"}, F{Num: 123}, nil)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate3(t *testing.T) {
	type F struct {
	}
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), "")
	r := New(table, S{Num: 123, Text: "qwe"}, F{}, nil)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate4(t *testing.T) {
	type F struct {
		Num []int `db:"num" op:"in"`
	}
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num IN ($3)"))
	r := New(table, S{Num: 123, Text: "qwe"}, F{Num: []int{1, 2, 3}}, nil)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate5(t *testing.T) {
	type F struct {
		Num  int    `db:"num" op:"null"`
		Text string `db:"text" op:"notnull"`
	}
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num IS NULL AND text IS NOT NULL"))
	r := New(table, S{Num: 123, Text: "qwe"}, F{}, nil)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQuerySelectLatest(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(selectTemplate, "num, text", table), fmt.Sprintf(whereTemplate, "num = $1"))
	r := New(table, S{}, nil, FilterNode{"num": Node{Operator: "eq", Value: 123}})
	s, arg := r.SelectLatest()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}
