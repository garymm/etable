// Code generated by "goki generate ./..."; DO NOT EDIT.

package agg

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _AggsValues = []Aggs{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

// AggsN is the highest valid value
// for type Aggs, plus one.
const AggsN Aggs = 16

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _AggsNoOp() {
	var x [1]struct{}
	_ = x[AggCount-(0)]
	_ = x[AggSum-(1)]
	_ = x[AggProd-(2)]
	_ = x[AggMin-(3)]
	_ = x[AggMax-(4)]
	_ = x[AggMean-(5)]
	_ = x[AggVar-(6)]
	_ = x[AggStd-(7)]
	_ = x[AggSem-(8)]
	_ = x[AggVarPop-(9)]
	_ = x[AggStdPop-(10)]
	_ = x[AggSemPop-(11)]
	_ = x[AggMedian-(12)]
	_ = x[AggQ1-(13)]
	_ = x[AggQ3-(14)]
	_ = x[AggSumSq-(15)]
}

var _AggsNameToValueMap = map[string]Aggs{
	`AggCount`:  0,
	`aggcount`:  0,
	`AggSum`:    1,
	`aggsum`:    1,
	`AggProd`:   2,
	`aggprod`:   2,
	`AggMin`:    3,
	`aggmin`:    3,
	`AggMax`:    4,
	`aggmax`:    4,
	`AggMean`:   5,
	`aggmean`:   5,
	`AggVar`:    6,
	`aggvar`:    6,
	`AggStd`:    7,
	`aggstd`:    7,
	`AggSem`:    8,
	`aggsem`:    8,
	`AggVarPop`: 9,
	`aggvarpop`: 9,
	`AggStdPop`: 10,
	`aggstdpop`: 10,
	`AggSemPop`: 11,
	`aggsempop`: 11,
	`AggMedian`: 12,
	`aggmedian`: 12,
	`AggQ1`:     13,
	`aggq1`:     13,
	`AggQ3`:     14,
	`aggq3`:     14,
	`AggSumSq`:  15,
	`aggsumsq`:  15,
}

var _AggsDescMap = map[Aggs]string{
	0:  `Count of number of elements`,
	1:  `Sum of elements`,
	2:  `Product of elements`,
	3:  `Min minimum value`,
	4:  `Max maximum value`,
	5:  `Mean mean value`,
	6:  `Var sample variance (squared diffs from mean, divided by n-1)`,
	7:  `Std sample standard deviation (sqrt of Var)`,
	8:  `Sem sample standard error of the mean (Std divided by sqrt(n))`,
	9:  `VarPop population variance (squared diffs from mean, divided by n)`,
	10: `StdPop population standard deviation (sqrt of VarPop)`,
	11: `SemPop population standard error of the mean (StdPop divided by sqrt(n))`,
	12: `Median middle value in sorted ordering`,
	13: `Q1 first quartile = 25%ile value = .25 quantile value`,
	14: `Q3 third quartile = 75%ile value = .75 quantile value`,
	15: `SumSq sum of squares`,
}

var _AggsMap = map[Aggs]string{
	0:  `AggCount`,
	1:  `AggSum`,
	2:  `AggProd`,
	3:  `AggMin`,
	4:  `AggMax`,
	5:  `AggMean`,
	6:  `AggVar`,
	7:  `AggStd`,
	8:  `AggSem`,
	9:  `AggVarPop`,
	10: `AggStdPop`,
	11: `AggSemPop`,
	12: `AggMedian`,
	13: `AggQ1`,
	14: `AggQ3`,
	15: `AggSumSq`,
}

// String returns the string representation
// of this Aggs value.
func (i Aggs) String() string {
	if str, ok := _AggsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Aggs value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Aggs) SetString(s string) error {
	if val, ok := _AggsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _AggsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Aggs")
}

// Int64 returns the Aggs value as an int64.
func (i Aggs) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Aggs value from an int64.
func (i *Aggs) SetInt64(in int64) {
	*i = Aggs(in)
}

// Desc returns the description of the Aggs value.
func (i Aggs) Desc() string {
	if str, ok := _AggsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// AggsValues returns all possible values
// for the type Aggs.
func AggsValues() []Aggs {
	return _AggsValues
}

// Values returns all possible values
// for the type Aggs.
func (i Aggs) Values() []enums.Enum {
	res := make([]enums.Enum, len(_AggsValues))
	for i, d := range _AggsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Aggs.
func (i Aggs) IsValid() bool {
	_, ok := _AggsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Aggs) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Aggs) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}
