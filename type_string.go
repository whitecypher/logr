// generated by stringer -type=Type; DO NOT EDIT

package logr

import "fmt"

const (
	_Type_name_0 = "None"
	_Type_name_1 = "P"
	_Type_name_2 = "E"
	_Type_name_3 = "W"
	_Type_name_4 = "I"
	_Type_name_5 = "D"
	_Type_name_6 = "S"
)

var (
	_Type_index_0 = [...]uint8{0, 4}
	_Type_index_1 = [...]uint8{0, 1}
	_Type_index_2 = [...]uint8{0, 1}
	_Type_index_3 = [...]uint8{0, 1}
	_Type_index_4 = [...]uint8{0, 1}
	_Type_index_5 = [...]uint8{0, 1}
	_Type_index_6 = [...]uint8{0, 1}
)

func (i Type) String() string {
	switch {
	case i == 0:
		return _Type_name_0
	case i == 2:
		return _Type_name_1
	case i == 4:
		return _Type_name_2
	case i == 8:
		return _Type_name_3
	case i == 16:
		return _Type_name_4
	case i == 32:
		return _Type_name_5
	case i == 64:
		return _Type_name_6
	default:
		return fmt.Sprintf("Type(%d)", i)
	}
}