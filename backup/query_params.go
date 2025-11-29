// package simple_qb

// var (
// 	insertTemplate = "INSERT INTO %s (%s) VALUES (%s)"
// 	selectTemplate = "SELECT %s FROM %s"
// 	updateTemplate = "UPDATE %s SET (%s) = ROW(%s)"
// 	deleteTemplate = "DELETE FROM %s"
// 	whereTemplate  = "WHERE %s"
// )

// const (
// 	tag = "db"
// )

// type qBilder struct {
// 	table  string
// 	data   any
// 	params ParamNode
// }

// type Node struct {
// 	Tag      string
// 	Operator string
// 	Value    any
// 	Logic    string
// }

// type ParamNode = []*Node

// var opMap = map[string]string{
// 	"eq":      "=",           // равно
// 	"neq":     "<>",          // неравно
// 	"lt":      "<",           // меньше
// 	"lte":     "<=",          // меньше или равно
// 	"gt":      ">",           // больше
// 	"gte":     ">=",          // больше или равно
// 	"like":    "LIKE",        // похоже на (для строковых выражений)
// 	"in":      "IN",          // входит в перечень (depricatet)
// 	"null":    "IS NULL",     // пустое значение
// 	"notnull": "IS NOT NULL", // непустое значение
// }
