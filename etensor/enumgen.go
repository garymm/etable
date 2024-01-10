// Code generated by "goki generate ./..."; DO NOT EDIT.

package etensor

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _TypeValues = []Type{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

// TypeN is the highest valid value
// for type Type, plus one.
const TypeN Type = 15

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _TypeNoOp() {
	var x [1]struct{}
	_ = x[NULL-(0)]
	_ = x[BOOL-(1)]
	_ = x[UINT8-(2)]
	_ = x[INT8-(3)]
	_ = x[UINT16-(4)]
	_ = x[INT16-(5)]
	_ = x[UINT32-(6)]
	_ = x[INT32-(7)]
	_ = x[UINT64-(8)]
	_ = x[INT64-(9)]
	_ = x[FLOAT16-(10)]
	_ = x[FLOAT32-(11)]
	_ = x[FLOAT64-(12)]
	_ = x[STRING-(13)]
	_ = x[INT-(14)]
}

var _TypeNameToValueMap = map[string]Type{
	`NULL`:    0,
	`null`:    0,
	`BOOL`:    1,
	`bool`:    1,
	`UINT8`:   2,
	`uint8`:   2,
	`INT8`:    3,
	`int8`:    3,
	`UINT16`:  4,
	`uint16`:  4,
	`INT16`:   5,
	`int16`:   5,
	`UINT32`:  6,
	`uint32`:  6,
	`INT32`:   7,
	`int32`:   7,
	`UINT64`:  8,
	`uint64`:  8,
	`INT64`:   9,
	`int64`:   9,
	`FLOAT16`: 10,
	`float16`: 10,
	`FLOAT32`: 11,
	`float32`: 11,
	`FLOAT64`: 12,
	`float64`: 12,
	`STRING`:  13,
	`string`:  13,
	`INT`:     14,
	`int`:     14,
}

var _TypeDescMap = map[Type]string{
	0:  `Null type having no physical storage`,
	1:  `Bool is a 1 bit, LSB bit-packed ordering`,
	2:  `UINT8 is an Unsigned 8-bit little-endian integer`,
	3:  `INT8 is a Signed 8-bit little-endian integer`,
	4:  `UINT16 is an Unsigned 16-bit little-endian integer`,
	5:  `INT16 is a Signed 16-bit little-endian integer`,
	6:  `UINT32 is an Unsigned 32-bit little-endian integer`,
	7:  `INT32 is a Signed 32-bit little-endian integer`,
	8:  `UINT64 is an Unsigned 64-bit little-endian integer`,
	9:  `INT64 is a Signed 64-bit little-endian integer`,
	10: `FLOAT16 is a 2-byte floating point value`,
	11: `FLOAT32 is a 4-byte floating point value`,
	12: `FLOAT64 is an 8-byte floating point value`,
	13: `STRING is a UTF8 variable-length string`,
	14: `INT is a Signed 64-bit little-endian integer -- should only use on 64bit machines!`,
}

var _TypeMap = map[Type]string{
	0:  `NULL`,
	1:  `BOOL`,
	2:  `UINT8`,
	3:  `INT8`,
	4:  `UINT16`,
	5:  `INT16`,
	6:  `UINT32`,
	7:  `INT32`,
	8:  `UINT64`,
	9:  `INT64`,
	10: `FLOAT16`,
	11: `FLOAT32`,
	12: `FLOAT64`,
	13: `STRING`,
	14: `INT`,
}

// String returns the string representation
// of this Type value.
func (i Type) String() string {
	if str, ok := _TypeMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Type value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Type) SetString(s string) error {
	if val, ok := _TypeNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _TypeNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Type")
}

// Int64 returns the Type value as an int64.
func (i Type) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Type value from an int64.
func (i *Type) SetInt64(in int64) {
	*i = Type(in)
}

// Desc returns the description of the Type value.
func (i Type) Desc() string {
	if str, ok := _TypeDescMap[i]; ok {
		return str
	}
	return i.String()
}

// TypeValues returns all possible values
// for the type Type.
func TypeValues() []Type {
	return _TypeValues
}

// Values returns all possible values
// for the type Type.
func (i Type) Values() []enums.Enum {
	res := make([]enums.Enum, len(_TypeValues))
	for i, d := range _TypeValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Type.
func (i Type) IsValid() bool {
	_, ok := _TypeMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Type) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Type) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}
