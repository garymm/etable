// Code generated by "core generate"; DO NOT EDIT.

package eplot

import (
	"cogentcore.org/core/enums"
)

var _PlotTypesValues = []PlotTypes{0, 1}

// PlotTypesN is the highest valid value for type PlotTypes, plus one.
const PlotTypesN PlotTypes = 2

var _PlotTypesValueMap = map[string]PlotTypes{`XY`: 0, `Bar`: 1}

var _PlotTypesDescMap = map[PlotTypes]string{0: `XY is a standard line / point plot`, 1: `Bar plots vertical bars`}

var _PlotTypesMap = map[PlotTypes]string{0: `XY`, 1: `Bar`}

// String returns the string representation of this PlotTypes value.
func (i PlotTypes) String() string { return enums.String(i, _PlotTypesMap) }

// SetString sets the PlotTypes value from its string representation,
// and returns an error if the string is invalid.
func (i *PlotTypes) SetString(s string) error {
	return enums.SetString(i, s, _PlotTypesValueMap, "PlotTypes")
}

// Int64 returns the PlotTypes value as an int64.
func (i PlotTypes) Int64() int64 { return int64(i) }

// SetInt64 sets the PlotTypes value from an int64.
func (i *PlotTypes) SetInt64(in int64) { *i = PlotTypes(in) }

// Desc returns the description of the PlotTypes value.
func (i PlotTypes) Desc() string { return enums.Desc(i, _PlotTypesDescMap) }

// PlotTypesValues returns all possible values for the type PlotTypes.
func PlotTypesValues() []PlotTypes { return _PlotTypesValues }

// Values returns all possible values for the type PlotTypes.
func (i PlotTypes) Values() []enums.Enum { return enums.Values(_PlotTypesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i PlotTypes) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *PlotTypes) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "PlotTypes")
}

var _ShapesValues = []Shapes{0, 1, 2, 3, 4, 5, 6, 7}

// ShapesN is the highest valid value for type Shapes, plus one.
const ShapesN Shapes = 8

var _ShapesValueMap = map[string]Shapes{`Ring`: 0, `Circle`: 1, `Square`: 2, `Box`: 3, `Triangle`: 4, `Pyramid`: 5, `Plus`: 6, `Cross`: 7}

var _ShapesDescMap = map[Shapes]string{0: `Ring is the outline of a circle`, 1: `Circle is a solid circle`, 2: `Square is the outline of a square`, 3: `Box is a filled square`, 4: `Triangle is the outline of a triangle`, 5: `Pyramid is a filled triangle`, 6: `Plus is a plus sign`, 7: `Cross is a big X`}

var _ShapesMap = map[Shapes]string{0: `Ring`, 1: `Circle`, 2: `Square`, 3: `Box`, 4: `Triangle`, 5: `Pyramid`, 6: `Plus`, 7: `Cross`}

// String returns the string representation of this Shapes value.
func (i Shapes) String() string { return enums.String(i, _ShapesMap) }

// SetString sets the Shapes value from its string representation,
// and returns an error if the string is invalid.
func (i *Shapes) SetString(s string) error { return enums.SetString(i, s, _ShapesValueMap, "Shapes") }

// Int64 returns the Shapes value as an int64.
func (i Shapes) Int64() int64 { return int64(i) }

// SetInt64 sets the Shapes value from an int64.
func (i *Shapes) SetInt64(in int64) { *i = Shapes(in) }

// Desc returns the description of the Shapes value.
func (i Shapes) Desc() string { return enums.Desc(i, _ShapesDescMap) }

// ShapesValues returns all possible values for the type Shapes.
func ShapesValues() []Shapes { return _ShapesValues }

// Values returns all possible values for the type Shapes.
func (i Shapes) Values() []enums.Enum { return enums.Values(_ShapesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Shapes) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Shapes) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Shapes") }
