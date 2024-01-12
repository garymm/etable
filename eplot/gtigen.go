// Code generated by "goki generate"; DO NOT EDIT.

package eplot

import (
	"goki.dev/gi"
	"goki.dev/gti"
	"goki.dev/ki/v2"
	"goki.dev/ordmap"
)

// Plot2DType is the [gti.Type] for [Plot2D]
var Plot2DType = gti.AddType(&gti.Type{
	Name:      "github.com/goki/etable/v2/eplot.Plot2D",
	ShortName: "eplot.Plot2D",
	IDName:    "plot-2-d",
	Doc:       "Plot2D is a GoGi Widget that provides a 2D plot of selected columns of etable data",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Table", &gti.Field{Name: "Table", Type: "*github.com/goki/etable/v2/etable.IdxView", LocalType: "*etable.IdxView", Doc: "the idxview of the table that we're plotting", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"Params", &gti.Field{Name: "Params", Type: "github.com/goki/etable/v2/eplot.PlotParams", LocalType: "PlotParams", Doc: "the overall plot parameters", Directives: gti.Directives{}, Tag: ""}},
		{"Cols", &gti.Field{Name: "Cols", Type: "[]*github.com/goki/etable/v2/eplot.ColParams", LocalType: "[]*ColParams", Doc: "the parameters for each column of the table", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"Plot", &gti.Field{Name: "Plot", Type: "*gonum.org/v1/plot.Plot", LocalType: "*plot.Plot", Doc: "the gonum plot that actually does the plotting -- always save the last one generated", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
		{"ConfigPlotFunc", &gti.Field{Name: "ConfigPlotFunc", Type: "func()", LocalType: "func()", Doc: "ConfigPlotFunc is a function to call to configure [Plot2D.Plot], the gonum plot that\nactually does the plotting. It is called after [Plot] is generated, and properties\nof [Plot] can be modified in it. Properties of [Plot] should not be modified outside\nof this function, as doing so will have no effect.", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\""}},
		{"SVGFile", &gti.Field{Name: "SVGFile", Type: "goki.dev/gi.FileName", LocalType: "gi.FileName", Doc: "current svg file", Directives: gti.Directives{}, Tag: ""}},
		{"DataFile", &gti.Field{Name: "DataFile", Type: "goki.dev/gi.FileName", LocalType: "gi.FileName", Doc: "current csv data file", Directives: gti.Directives{}, Tag: ""}},
		{"InPlot", &gti.Field{Name: "InPlot", Type: "bool", LocalType: "bool", Doc: "currently doing a plot", Directives: gti.Directives{}, Tag: "set:\"-\" edit:\"-\" json:\"-\" xml:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{
		{"SaveSVG", &gti.Method{Name: "SaveSVG", Doc: "SaveSVG saves the plot to an svg -- first updates to ensure that plot is current", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"fname", &gti.Field{Name: "fname", Type: "goki.dev/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"SavePNG", &gti.Method{Name: "SavePNG", Doc: "SavePNG saves the current plot to a png, capturing current render", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"fname", &gti.Field{Name: "fname", Type: "goki.dev/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"SaveCSV", &gti.Method{Name: "SaveCSV", Doc: "SaveCSV saves the Table data to a csv (comma-separated values) file with headers (any delim)", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"fname", &gti.Field{Name: "fname", Type: "goki.dev/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
			{"delim", &gti.Field{Name: "delim", Type: "github.com/goki/etable/v2/etable.Delims", LocalType: "etable.Delims", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"SaveAll", &gti.Method{Name: "SaveAll", Doc: "SaveAll saves the current plot to a png, svg, and the data to a tsv -- full save\nAny extension is removed and appropriate extensions are added", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"fname", &gti.Field{Name: "fname", Type: "goki.dev/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"OpenCSV", &gti.Method{Name: "OpenCSV", Doc: "OpenCSV opens the Table data from a csv (comma-separated values) file (or any delim)", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"fname", &gti.Field{Name: "fname", Type: "goki.dev/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
			{"delim", &gti.Field{Name: "delim", Type: "github.com/goki/etable/v2/etable.Delims", LocalType: "etable.Delims", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"SetColsByName", &gti.Method{Name: "SetColsByName", Doc: "SetColsByName turns cols On or Off if their name contains given string", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"nameContains", &gti.Field{Name: "nameContains", Type: "string", LocalType: "string", Doc: "", Directives: gti.Directives{}, Tag: ""}},
			{"on", &gti.Field{Name: "on", Type: "bool", LocalType: "bool", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
	}),
	Instance: &Plot2D{},
})

// NewPlot2D adds a new [Plot2D] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewPlot2D(par ki.Ki, name ...string) *Plot2D {
	return par.NewChild(Plot2DType, name...).(*Plot2D)
}

// KiType returns the [*gti.Type] of [Plot2D]
func (t *Plot2D) KiType() *gti.Type {
	return Plot2DType
}

// New returns a new [*Plot2D] value
func (t *Plot2D) New() ki.Ki {
	return &Plot2D{}
}

// SetParams sets the [Plot2D.Params]:
// the overall plot parameters
func (t *Plot2D) SetParams(v PlotParams) *Plot2D {
	t.Params = v
	return t
}

// SetConfigPlotFunc sets the [Plot2D.ConfigPlotFunc]:
// ConfigPlotFunc is a function to call to configure [Plot2D.Plot], the gonum plot that
// actually does the plotting. It is called after [Plot] is generated, and properties
// of [Plot] can be modified in it. Properties of [Plot] should not be modified outside
// of this function, as doing so will have no effect.
func (t *Plot2D) SetConfigPlotFunc(v func()) *Plot2D {
	t.ConfigPlotFunc = v
	return t
}

// SetSvgfile sets the [Plot2D.SVGFile]:
// current svg file
func (t *Plot2D) SetSvgfile(v gi.FileName) *Plot2D {
	t.SVGFile = v
	return t
}

// SetDataFile sets the [Plot2D.DataFile]:
// current csv data file
func (t *Plot2D) SetDataFile(v gi.FileName) *Plot2D {
	t.DataFile = v
	return t
}

// SetTooltip sets the [Plot2D.Tooltip]
func (t *Plot2D) SetTooltip(v string) *Plot2D {
	t.Tooltip = v
	return t
}

// SetStackTop sets the [Plot2D.StackTop]
func (t *Plot2D) SetStackTop(v int) *Plot2D {
	t.StackTop = v
	return t
}

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/goki/etable/v2/eplot.PlotParams",
	ShortName: "eplot.PlotParams",
	IDName:    "plot-params",
	Doc:       "PlotParams are parameters for overall plot",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Title", &gti.Field{Name: "Title", Type: "string", LocalType: "string", Doc: "optional title at top of plot", Directives: gti.Directives{}, Tag: ""}},
		{"Type", &gti.Field{Name: "Type", Type: "github.com/goki/etable/v2/eplot.PlotTypes", LocalType: "PlotTypes", Doc: "type of plot to generate.  For a Bar plot, items are plotted ordinally by row and the XAxis is optional", Directives: gti.Directives{}, Tag: ""}},
		{"Lines", &gti.Field{Name: "Lines", Type: "bool", LocalType: "bool", Doc: "whether to plot lines", Directives: gti.Directives{}, Tag: ""}},
		{"Points", &gti.Field{Name: "Points", Type: "bool", LocalType: "bool", Doc: "whether to plot points with symbols", Directives: gti.Directives{}, Tag: ""}},
		{"LineWidth", &gti.Field{Name: "LineWidth", Type: "float64", LocalType: "float64", Doc: "width of lines", Directives: gti.Directives{}, Tag: ""}},
		{"PointSize", &gti.Field{Name: "PointSize", Type: "float64", LocalType: "float64", Doc: "size of points", Directives: gti.Directives{}, Tag: ""}},
		{"PointShape", &gti.Field{Name: "PointShape", Type: "github.com/goki/etable/v2/eplot.Shapes", LocalType: "Shapes", Doc: "the shape used to draw points", Directives: gti.Directives{}, Tag: ""}},
		{"BarWidth", &gti.Field{Name: "BarWidth", Type: "float64", LocalType: "float64", Doc: "width of bars for bar plot, as fraction of available space -- 1 = no gaps, .8 default", Directives: gti.Directives{}, Tag: "min:\"0.01\" max:\"1\""}},
		{"NegXDraw", &gti.Field{Name: "NegXDraw", Type: "bool", LocalType: "bool", Doc: "draw lines that connect points with a negative X-axis direction -- otherwise these are treated as breaks between repeated series and not drawn", Directives: gti.Directives{}, Tag: ""}},
		{"Scale", &gti.Field{Name: "Scale", Type: "float64", LocalType: "float64", Doc: "overall scaling factor -- the larger the number, the larger the fonts are relative to the graph", Directives: gti.Directives{}, Tag: "def:\"2\""}},
		{"XAxisCol", &gti.Field{Name: "XAxisCol", Type: "string", LocalType: "string", Doc: "what column to use for the common X axis -- if empty or not found, the row number is used.  This optional for Bar plots -- if present and LegendCol is also present, then an extra space will be put between X values.", Directives: gti.Directives{}, Tag: ""}},
		{"LegendCol", &gti.Field{Name: "LegendCol", Type: "string", LocalType: "string", Doc: "optional column for adding a separate colored / styled line or bar according to this value -- acts just like a separate Y variable, crossed with Y variables", Directives: gti.Directives{}, Tag: ""}},
		{"XAxisRot", &gti.Field{Name: "XAxisRot", Type: "float64", LocalType: "float64", Doc: "rotation of the X Axis labels, in degrees", Directives: gti.Directives{}, Tag: ""}},
		{"XAxisLabel", &gti.Field{Name: "XAxisLabel", Type: "string", LocalType: "string", Doc: "optional label to use for XAxis instead of column name", Directives: gti.Directives{}, Tag: ""}},
		{"YAxisLabel", &gti.Field{Name: "YAxisLabel", Type: "string", LocalType: "string", Doc: "optional label to use for YAxis -- if empty, first column name is used", Directives: gti.Directives{}, Tag: ""}},
		{"Plot", &gti.Field{Name: "Plot", Type: "*github.com/goki/etable/v2/eplot.Plot2D", LocalType: "*Plot2D", Doc: "our plot, for update method", Directives: gti.Directives{}, Tag: "copy:\"-\" json:\"-\" xml:\"-\" view:\"-\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/goki/etable/v2/eplot.ColParams",
	ShortName: "eplot.ColParams",
	IDName:    "col-params",
	Doc:       "ColParams are parameters for plotting one column of data",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"On", &gti.Field{Name: "On", Type: "bool", LocalType: "bool", Doc: "whether to plot this column", Directives: gti.Directives{}, Tag: ""}},
		{"Col", &gti.Field{Name: "Col", Type: "string", LocalType: "string", Doc: "name of column we're plotting", Directives: gti.Directives{}, Tag: "label:\"Column\""}},
		{"Lines", &gti.Field{Name: "Lines", Type: "goki.dev/glop/option.Option", LocalType: "option.Option[bool]", Doc: "whether to plot lines; uses the overall plot option if unset", Directives: gti.Directives{}, Tag: ""}},
		{"Points", &gti.Field{Name: "Points", Type: "goki.dev/glop/option.Option", LocalType: "option.Option[bool]", Doc: "whether to plot points with symbols; uses the overall plot option if unset", Directives: gti.Directives{}, Tag: ""}},
		{"LineWidth", &gti.Field{Name: "LineWidth", Type: "goki.dev/glop/option.Option", LocalType: "option.Option[float64]", Doc: "the width of lines; uses the overall plot option if unset", Directives: gti.Directives{}, Tag: ""}},
		{"PointSize", &gti.Field{Name: "PointSize", Type: "goki.dev/glop/option.Option", LocalType: "option.Option[float64]", Doc: "the size of points; uses the overall plot option if unset", Directives: gti.Directives{}, Tag: ""}},
		{"PointShape", &gti.Field{Name: "PointShape", Type: "goki.dev/glop/option.Option", LocalType: "option.Option[Shapes]", Doc: "the shape used to draw points; uses the overall plot option if unset", Directives: gti.Directives{}, Tag: ""}},
		{"Range", &gti.Field{Name: "Range", Type: "github.com/goki/etable/v2/minmax.Range64", LocalType: "minmax.Range64", Doc: "effective range of data to plot -- either end can be fixed", Directives: gti.Directives{}, Tag: ""}},
		{"FullRange", &gti.Field{Name: "FullRange", Type: "github.com/goki/etable/v2/minmax.F64", LocalType: "minmax.F64", Doc: "full actual range of data -- only valid if specifically computed", Directives: gti.Directives{}, Tag: ""}},
		{"Color", &gti.Field{Name: "Color", Type: "image/color.Color", LocalType: "color.Color", Doc: "color to use when plotting the line / column", Directives: gti.Directives{}, Tag: ""}},
		{"NTicks", &gti.Field{Name: "NTicks", Type: "int", LocalType: "int", Doc: "desired number of ticks", Directives: gti.Directives{}, Tag: ""}},
		{"Lbl", &gti.Field{Name: "Lbl", Type: "string", LocalType: "string", Doc: "if non-empty, this is an alternative label to use in plotting", Directives: gti.Directives{}, Tag: "label:\"Label\""}},
		{"TensorIdx", &gti.Field{Name: "TensorIdx", Type: "int", LocalType: "int", Doc: "if column has n-dimensional tensor cells in each row, this is the index within each cell to plot -- use -1 to plot *all* indexes as separate lines", Directives: gti.Directives{}, Tag: ""}},
		{"ErrCol", &gti.Field{Name: "ErrCol", Type: "string", LocalType: "string", Doc: "specifies a column containing error bars for this column", Directives: gti.Directives{}, Tag: ""}},
		{"IsString", &gti.Field{Name: "IsString", Type: "bool", LocalType: "bool", Doc: "if true this is a string column -- plots as labels", Directives: gti.Directives{}, Tag: "edit:\"-\""}},
		{"Plot", &gti.Field{Name: "Plot", Type: "*github.com/goki/etable/v2/eplot.Plot2D", LocalType: "*Plot2D", Doc: "our plot, for update method", Directives: gti.Directives{}, Tag: "copy:\"-\" json:\"-\" xml:\"-\" view:\"-\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
