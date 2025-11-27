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
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect(t *testing.T) {
	check := fmt.Sprintf(selectTemplate, "num, text", table)
	r := New(table, S{}, nil)
	s, arg := r.Select()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelectWhere(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(selectTemplate, "num, text", table), fmt.Sprintf(whereTemplate, "num = $1"))
	filter := NewParam(NewNode("num", "eq", 123))
	r := New(table, S{}, filter)
	s, arg := r.Select()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelectWhere1(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(selectTemplate, "num, text", table), fmt.Sprintf(whereTemplate, "num = $1"))
	filter := NewParam(NewNode("num", "eq", 123))
	r := New(table, S{}, filter)
	s, arg := r.Select()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQueryUpdate(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num = $3"))
	filter := NewParam(NewNode("num", "eq", 123))
	r := New(table, S{Num: 123, Text: "qwe"}, filter)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQueryUpdate2(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num > $3"))
	filter := NewParam(NewNode("num", "gt", 123))
	r := New(table, S{Num: 123, Text: "qwe"}, filter)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQueryUpdate3(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), "")
	r := New(table, S{Num: 123, Text: "qwe"}, ParamNode{})
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQueryUpdate4(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num IN ($3)"))
	filter := NewParam(NewNode("num", "in", []int{1, 2, 3}))
	r := New(table, S{Num: 123, Text: "qwe"}, filter)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQueryUpdate5(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num IS NULL OR text IS NOT NULL"))
	filter := NewParam(NewNode("num", "null", 123), NewNodeOr("text", "notnull", 123))
	r := New(table, S{Num: 123, Text: "qwe"}, filter)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQueryUpdate6(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num IS NULL AND text IS NOT NULL"))
	filter := NewParam(NewNode("num", "null", 123), NewNode("text", "notnull", 123))
	r := New(table, S{Num: 123, Text: "qwe"}, filter)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQueryUpdate7(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(updateTemplate, table, "num, text", "$1, $2"), fmt.Sprintf(whereTemplate, "num IS NULL AND text IS NOT NULL OR text >= $3"))
	filter := NewParam(NewNode("num", "null", 123), NewNode("text", "notnull", 123), NewNodeOr("text", "gte", 123))
	r := New(table, S{Num: 123, Text: "qwe"}, filter)
	s, arg, _ := r.Update()
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect1(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(selectTemplate, "num, text", table), "ORDER BY num ASC")
	r := New(table, S{}, nil)
	s, arg := r.Select()
	s = OrderBy(s, "num", "ASC")
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect2(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(selectTemplate, "num, text", table), "LIMIT 10")
	r := New(table, S{}, nil)
	s, arg := r.Select()
	s = Limit(s, 10, 0)
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect3(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(selectTemplate, "num, text", table), "LIMIT 10 OFFSET 10")
	r := New(table, S{}, nil)
	s, arg := r.Select()
	s = Limit(s, 10, 10)
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect4(t *testing.T) {
	check := fmt.Sprintf("%s %s", fmt.Sprintf(selectTemplate, "num, text", table), "ORDER BY num DESC LIMIT 10 OFFSET 10")
	r := New(table, S{}, nil)
	s, arg := r.Select()
	s = Limit(OrderBy(s, "num", "DESC"), 10, 10)
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}

func TestQuerySelect5(t *testing.T) {
	check := fmt.Sprintf("%s %s %s", fmt.Sprintf(selectTemplate, "num, text", table), fmt.Sprintf(whereTemplate, "num = $1"), "ORDER BY num DESC LIMIT 10 OFFSET 10")
	filter := NewParam(NewNode("num", "eq", 123))
	r := New(table, S{}, filter)
	s, arg := r.Select()
	s = Limit(OrderBy(s, "num", "DESC"), 10, 10)
	fmt.Println(s, arg)
	if s != check {
		t.Errorf("error: %s || %s", s, check)
	}
}
