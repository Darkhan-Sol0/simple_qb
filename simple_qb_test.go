package simple_qb

type S struct {
	Num  int    `db:"num"`
	Text string `db:"text"`
}

var table = "users"

// func TestQueryInsert(t *testing.T) {
// 	check := "INSERT INTO users (num, text) VALUES ($1, $2)"
// 	s, arg := New(table).Insert(S{Num: 1, Text: "hui"}).Generate()
// 	fmt.Println(s, arg)
// 	if s != check {
// 		t.Errorf("error: %s || %s", s, check)
// 	}
// }

// func TestQueryInsert1(t *testing.T) {
// 	check := "INSERT INTO users (num, text) VALUES ($1, $2)"
// 	s, arg := New(table).Insert(S{Num: 1}).Generate()
// 	fmt.Println(s, arg)
// 	if s != check {
// 		t.Errorf("error: %s || %s", s, check)
// 	}
// }

// func TestQuerySelect(t *testing.T) {
// 	check := fmt.Sprintf("SELECT num, text FROM %s", table)
// 	s, arg := New(table).Select(S{}).Generate()
// 	fmt.Println(s, arg)
// 	if s != check {
// 		t.Errorf("error: %s || %s", s, check)
// 	}
// }

// func TestQuerySelect1(t *testing.T) {
// 	check := fmt.Sprintf("SELECT num, text FROM %s WHERE num IN ($1, $2, $3)", table)
// 	s, arg := New(table).Select(S{}).Params(NewParam("num").In([]string{"1", "sad", "da11"})).Generate()
// 	fmt.Println(s, arg)
// 	if s != check {
// 		t.Errorf("error: %s || %s", s, check)
// 	}
// }

// func TestQuerySelect2(t *testing.T) {
// 	check := fmt.Sprintf("SELECT num, text FROM %s WHERE num IN ($1, $2, $3)", table)
// 	s, arg := New(table).Select(S{}).Params(NewParam("num").In([]string{"1", "sad", "da11"})).Generate()
// 	fmt.Println(s, arg)
// 	if s != check {
// 		t.Errorf("error: %s || %s", s, check)
// 	}
// }

// func TestQuerySelect3(t *testing.T) {
// 	check := fmt.Sprintf("SELECT num, text FROM %s WHERE num = $1 OR num < $2", table)
// 	s, arg := New(table).Select(S{}).Params(NewParam("num").Eq(1), NewParam("num").Less(5).Or()).Generate()
// 	fmt.Println(s, arg)
// 	if s != check {
// 		t.Errorf("error: %s || %s", s, check)
// 	}
// }

// func TestQuerySelect4(t *testing.T) {
// 	check := fmt.Sprintf("SELECT COUNT(num) FROM %s WHERE num = $1 OR num < $2", table)
// 	s, arg := New(table).Select(query.Count("num")).Params(NewParam("num").Eq(1), NewParam("num").Less(5).Or()).Generate()
// 	fmt.Println(s, arg)
// 	if s != check {
// 		t.Errorf("error: %s || %s", s, check)
// 	}
// }

// func TestQuerySelect5(t *testing.T) {
// 	check := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
// 	s, arg := New(table).Select(query.Count("")).Generate()
// 	fmt.Println(s, arg)
// 	if s != check {
// 		t.Errorf("error: %s || %s", s, check)
// 	}
// }
