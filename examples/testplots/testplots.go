// Copyright (c) 2019, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cogentcore.org/core/core"
	"github.com/emer/etable/v2/eplot"
	"github.com/emer/etable/v2/etable"
	"github.com/emer/etable/v2/etensor"
)

func main() {
	b := core.NewBody("Test Plots")

	tv := core.NewTabs(b)

	PlotColorSpread(tv)

	b.RunMainWindow()
}

func PlotColorSpread(tv *core.Tabs) {
	label := "Color Spread"
	dt := etable.NewTable(label)
	dt.SetMetaData("name", label)
	dt.SetMetaData("read-only", "true")

	sch := etable.Schema{
		{"Index", etensor.INT, nil, nil},
		{"Collapse", etensor.INT, nil, nil},
		{"Val", etensor.FLOAT64, nil, nil},
	}
	dt.SetFromSchema(sch, 0)

	mx := 100
	dt.SetNumRows(mx)

	for i := 0; i < mx; i++ {
		val := i                                   // colors.BinarySpacedNumber(i)
		dt.SetCellFloat("Index", i, float64(i))    // select this to see the timecourse
		dt.SetCellFloat("Collapse", i, float64(0)) // select this to collapse all points on top
		dt.SetCellFloat("Val", i, float64(val))
	}

	pl := eplot.NewSubPlot(tv.NewTab(label))
	pl.SetTable(dt)
	pl.Params.XAxisCol = "Index"
	pl.Params.Lines = false
	pl.Params.Points = true
	pl.ColParams("Val").On = true
}
