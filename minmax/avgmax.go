// Copyright (c) 2019, The eTable Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package minmax

//gosl: start minmax

const (
	MaxFloat32 float32 = 3.402823466e+38
	MinFloat32 float32 = 1.175494351e-38
)

// AvgMax holds average and max statistics
type AvgMax32 struct {
	Avg    float32
	Max    float32
	Sum    float32 `desc:"sum for computing average"`
	MaxIdx int32   `desc:"index of max item"`
	N      int32   `desc:"number of items in sum"`

	pad, pad1, pad2 int32
}

// Init initializes prior to new updates
func (am *AvgMax32) Init() {
	am.Avg = 0
	am.Sum = 0
	am.N = 0
	am.Max = -MaxFloat32
	am.MaxIdx = -1
}

// UpdateVal updates stats from given value
func (am *AvgMax32) UpdateVal(val float32, idx int) {
	am.Sum += val
	am.N++
	if val > am.Max {
		am.Max = val
		am.MaxIdx = int32(idx)
	}
}

// CalcAvg computes the average given the current Sum and N values
func (am *AvgMax32) CalcAvg() {
	if am.N > 0 {
		am.Avg = am.Sum / float32(am.N)
	} else {
		am.Avg = am.Sum
		am.Max = am.Avg // prevents Max from being -MaxFloat..
	}
}

//gosl: end minmax

// UpdateFrom updates these values from other AvgMax32
func (am *AvgMax32) UpdateFrom(oth *AvgMax32) {
	am.Sum += oth.Sum
	am.N += oth.N
	if oth.Max > am.Max {
		am.Max = oth.Max
		am.MaxIdx = oth.MaxIdx
	}
}

// CopyFrom copies from other AvgMax32
func (am *AvgMax32) CopyFrom(oth *AvgMax32) {
	am.Avg = oth.Avg
	am.Max = oth.Max
	am.MaxIdx = oth.MaxIdx
	am.Sum = oth.Sum
	am.N = oth.N
}

///////////////////////////////////////////////////////////////////////////
//  64

// AvgMax holds average and max statistics
type AvgMax64 struct {
	Avg    float64
	Max    float64
	Sum    float64 `desc:"sum for computing average"`
	MaxIdx int32   `desc:"index of max item"`
	N      int32   `desc:"number of items in sum"`
}

// Init initializes prior to new updates
func (am *AvgMax64) Init() {
	am.Avg = 0
	am.Sum = 0
	am.N = 0
	am.Max = -MaxFloat64
	am.MaxIdx = -1
}

// UpdateVal updates stats from given value
func (am *AvgMax64) UpdateVal(val float64, idx int) {
	am.Sum += val
	am.N++
	if val > am.Max {
		am.Max = val
		am.MaxIdx = int32(idx)
	}
}

// CalcAvg computes the average given the current Sum and N values
func (am *AvgMax64) CalcAvg() {
	if am.N > 0 {
		am.Avg = am.Sum / float64(am.N)
	} else {
		am.Avg = am.Sum
		am.Max = am.Avg // prevents Max from being -MaxFloat..
	}
}

// UpdateFrom updates these values from other AvgMax64
func (am *AvgMax64) UpdateFrom(oth *AvgMax64) {
	am.Sum += oth.Sum
	am.N += oth.N
	if oth.Max > am.Max {
		am.Max = oth.Max
		am.MaxIdx = oth.MaxIdx
	}
}

// CopyFrom copies from other AvgMax64
func (am *AvgMax64) CopyFrom(oth *AvgMax64) {
	am.Avg = oth.Avg
	am.Max = oth.Max
	am.MaxIdx = oth.MaxIdx
	am.Sum = oth.Sum
	am.N = oth.N
}
