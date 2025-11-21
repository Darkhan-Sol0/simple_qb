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
	check := fmt.Sprintf(insertTemplate, table, "num, text", "$1, $2")
	r := New(table, S{Num: 123, Text: "qwe"}, nil)
	s, arg := r.Insert()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQuerySelect(t *testing.T) {
	check := fmt.Sprintf(selectTemplate, "num, text", table)
	r := New(table, S{}, nil)
	s, arg := r.Select()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQuerySelectWhere(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(selectTemplate, "num, text", table), fmt.Sprintf(whereTemplate, "num = $1"))
	filter := NewFillter(NewNode("num", "eq", 123))
	r := New(table, S{}, filter)
	s, arg := r.Select()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQuerySelectWhere1(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(selectTemplate, "num, text", table), fmt.Sprintf(whereTemplate, "num = $1"))
	filter := NewFillter(NewNode("num", "eq", 123))
	r := New(table, S{}, filter)
	s, arg := r.Select()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num = $3"))
	filter := NewFillter(NewNode("num", "eq", 123))
	r := New(table, S{Num: 123, Text: "qwe"}, filter)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate2(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num > $3"))
	filter := NewFillter(NewNode("num", "gt", 123))
	r := New(table, S{Num: 123, Text: "qwe"}, filter)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate3(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), "")
	r := New(table, S{Num: 123, Text: "qwe"}, FilterNode{})
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate4(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num IN ($3)"))
	filter := NewFillter(NewNode("num", "in", []int{1, 2, 3}))
	r := New(table, S{Num: 123, Text: "qwe"}, filter)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}

func TestQueryUpdate5(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num IS NULL AND text IS NOT NULL"))
	filter := NewFillter(NewNode("num", "null", 123), NewNode("text", "notnull", 123))
	r := New(table, S{Num: 123, Text: "qwe"}, filter)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error")
	}
}
