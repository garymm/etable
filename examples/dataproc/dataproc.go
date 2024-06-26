// Copyright (c) 2019, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"fmt"
	"math"

	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
	"github.com/emer/etable/v2/agg"
	"github.com/emer/etable/v2/etable"
	"github.com/emer/etable/v2/etview"
	"github.com/emer/etable/v2/split"
)

// Planets is raw data
var Planets *etable.Table

// PlanetsDesc are descriptive stats of all (non-Null) data
var PlanetsDesc *etable.Table

// PlanetsNNDesc are descriptive stats of planets where entire row is non-null
var PlanetsNNDesc *etable.Table

// GpMethodOrbit shows the median of orbital period as a function of method
var GpMethodOrbit *etable.Table

// GpMethodYear shows all stats of year described by orbit
var GpMethodYear *etable.Table

// GpMethodDecade shows number of planets found in each decade by given method
var GpMethodDecade *etable.Table

// GpDecade shows number of planets found in each decade
var GpDecade *etable.Table

//go:embed *.csv
var csv embed.FS

// AnalyzePlanets analyzes planets.csv data following some of the examples
// given here, using pandas:
//
//	https://jakevdp.github.io/PythonDataScienceHandbook/03.08-aggregation-and-grouping.html
func AnalyzePlanets() {
	Planets = etable.NewTable("planets")
	Planets.OpenFS(csv, "planets.csv", etable.Comma)

	PlanetsAll := etable.NewIndexView(Planets) // full original data

	NonNull := etable.NewIndexView(Planets)
	NonNull.Filter(etable.FilterNull) // filter out all rows with Null values

	PlanetsDesc = agg.DescAll(PlanetsAll) // individually excludes Null values in each col, but not row-wise
	PlanetsNNDesc = agg.DescAll(NonNull)  // standard descriptive stats for row-wise non-nulls

	byMethod := split.GroupBy(PlanetsAll, []string{"method"})
	split.Agg(byMethod, "orbital_period", agg.AggMedian)
	GpMethodOrbit = byMethod.AggsToTable(etable.AddAggName)

	byMethod.DeleteAggs()
	split.Desc(byMethod, "year") // full desc stats of year

	byMethod.Filter(func(idx int) bool {
		ag := byMethod.AggByColName("year:Std")
		return ag.Aggs[idx][0] > 0 // exclude results with 0 std
	})

	GpMethodYear = byMethod.AggsToTable(etable.AddAggName)

	byMethodDecade := split.GroupByFunc(PlanetsAll, func(row int) []string {
		meth := Planets.CellString("method", row)
		yr := Planets.CellFloat("year", row)
		decade := math.Floor(yr/10) * 10
		return []string{meth, fmt.Sprintf("%gs", decade)}
	})
	byMethodDecade.SetLevels("method", "decade")

	split.Agg(byMethodDecade, "number", agg.AggSum)

	// uncomment this to switch to decade first, then method
	// byMethodDecade.ReorderLevels([]int{1, 0})
	// byMethodDecade.SortLevels()

	decadeOnly, _ := byMethodDecade.ExtractLevels([]int{1})
	split.Agg(decadeOnly, "number", agg.AggSum)
	GpDecade = decadeOnly.AggsToTable(etable.AddAggName)

	GpMethodDecade = byMethodDecade.AggsToTable(etable.AddAggName) // here to ensure that decadeOnly didn't mess up..

	// todo: need unstack -- should be specific to the splits data because we already have the cols and
	// groups etc -- the ExtractLevels method provides key starting point.

	// todo: pivot table -- neeeds unstack function.

	// todo: could have a generic unstack-like method that takes a column for the data to turn into columns
	// and another that has the data to put in the cells.
}

func main() {
	AnalyzePlanets()

	b := core.NewBody("dataproc")
	tv := core.NewTabs(b)

	nt := tv.NewTab("Planets Data")
	tbv := etview.NewTableView(nt).SetTable(Planets)
	b.AddAppBar(tbv.ConfigToolbar)
	b.AddAppBar(func(tb *core.Toolbar) {
		core.NewButton(tb).SetText("README").SetIcon(icons.FileMarkdown).
			SetTooltip("open README help file").OnClick(func(e events.Event) {
			core.TheApp.OpenURL("https://github.com/emer/etable/blob/master/examples/dataproc/README.md")
		})
	})

	nt = tv.NewTab("Non-Null Rows Desc")
	etview.NewTableView(nt).SetTable(PlanetsNNDesc)
	nt = tv.NewTab("All Desc")
	etview.NewTableView(nt).SetTable(PlanetsDesc)
	nt = tv.NewTab("By Method Orbit")
	etview.NewTableView(nt).SetTable(GpMethodOrbit)
	nt = tv.NewTab("By Method Year")
	etview.NewTableView(nt).SetTable(GpMethodYear)
	nt = tv.NewTab("By Method Decade")
	etview.NewTableView(nt).SetTable(GpMethodDecade)
	nt = tv.NewTab("By Decade")
	etview.NewTableView(nt).SetTable(GpDecade)

	tv.SelectTabIndex(0)

	b.RunMainWindow()
}
