// package simple_qb

// import (
// 	"fmt"
// 	"reflect"
// 	"strings"
// )

// func getArguments(data any) (args []any) {
// 	v := reflect.ValueOf(data)
// 	t := v.Type()
// 	for i := range t.NumField() {
// 		dbTag := t.Field(i).Tag.Get(tag)
// 		val := v.Field(i)
// 		if dbTag != "" && dbTag != "-" && !val.IsZero() {
// 			args = append(args, val.Interface())
// 		}
// 	}
// 	return args
// }

// func getColums(data any) (colums []string) {
// 	v := reflect.ValueOf(data)
// 	t := v.Type()
// 	for i := range t.NumField() {
// 		dbTag := t.Field(i).Tag.Get(tag)
// 		if dbTag != "" && dbTag != "-" {
// 			colums = append(colums, dbTag)
// 		}
// 	}
// 	return colums
// }

// func getPlaceholders(count int) (placeholders []string) {
// 	for i := 1; i <= count; i++ {
// 		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
// 	}
// 	return placeholders
// }

// func getWhere(data ParamNode, startIndex int) (query string, args []any) {
// 	if len(data) == 0 {
// 		return "", nil
// 	}
// 	colums := make([]string, 2)
// 	count := 1
// 	for _, v := range data {
// 		logic := v.Logic
// 		operator := v.Operator
// 		tag := v.Tag
// 		val := v.Value
// 		if operator != "" && opMap[operator] != "" {
// 			switch operator {
// 			case "in":
// 				colums[1] = fmt.Sprintf("%s %s ($%d)", tag, opMap[operator], startIndex+count)
// 				args = append(args, val)
// 				count++
// 			case "null", "notnull":
// 				colums[1] = fmt.Sprintf("%s %s", tag, opMap[operator])

// 			default:
// 				colums[1] = fmt.Sprintf("%s %s $%d", tag, opMap[operator], startIndex+count)
// 				args = append(args, val)
// 				count++
// 			}
// 			if colums[0] != "" {
// 				text := strings.Join(colums, fmt.Sprintf(" %s ", logic))
// 				colums[0] = text
// 			} else {
// 				colums[0] = colums[1]
// 			}
// 			colums[1] = ""
// 		}
// 	}
// 	if len(colums) == 0 {
// 		return "", nil
// 	}
// 	return fmt.Sprintf(whereTemplate, colums[0]), args
// }
