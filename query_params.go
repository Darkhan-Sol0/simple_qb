package simple_qb

// Insert is a template for constructing INSERT INTO commands.
// Select is a template for constructing SELECT commands.
// Update is a template for constructing UPDATE commands.
// Where is a template for constructing WHERE clauses.
var (
	insertTemplate = "INSERT INTO %s (%s) VALUES (%s)"
	selectTemplate = "SELECT %s FROM %s"
	updateTemplate = "UPDATE %s SET (%s) = ROW(%s)"
	deleteTemplate = "DELETE FROM %s"
	whereTemplate  = "WHERE %s"
)

// Tag identifies column names within structures.
// Op specifies the operation applied to data in WHERE conditions.
const (
	tag = "db"
)

type qBilder struct {
	table  string
	data   any
	params ParamNode
}

type Node struct {
	Tag      string
	Operator string
	Value    any
	Logic    string
}

type ParamNode = []*Node

// пока такие теги, может посже изменить

// Operation mappings convert internal representations into corresponding SQL operators.
var opMap = map[string]string{
	"eq":      "=",           // равно
	"neq":     "<>",          // неравно
	"lt":      "<",           // меньше
	"lte":     "<=",          // меньше или равно
	"gt":      ">",           // больше
	"gte":     ">=",          // больше или равно
	"like":    "LIKE",        // похоже на (для строковых выражений)
	"in":      "IN",          // входит в перечень (depricatet)
	"null":    "IS NULL",     // пустое значение
	"notnull": "IS NOT NULL", // непустое значение
}
