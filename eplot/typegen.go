// Code generated by "core generate"; DO NOT EDIT.

package eplot

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/eplot.Plot2D", IDName: "plot2-d", Doc: "Plot2D is a Cogent Core Widget that provides a 2D plot of selected columns of etable data", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Methods: []types.Method{{Name: "SaveSVG", Doc: "SaveSVG saves the plot to an svg -- first updates to ensure that plot is current", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname"}}, {Name: "SavePNG", Doc: "SavePNG saves the current plot to a png, capturing current render", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname"}}, {Name: "SaveCSV", Doc: "SaveCSV saves the Table data to a csv (comma-separated values) file with headers (any delim)", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname", "delim"}}, {Name: "SaveAll", Doc: "SaveAll saves the current plot to a png, svg, and the data to a tsv -- full save\nAny extension is removed and appropriate extensions are added", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"fname"}}, {Name: "OpenCSV", Doc: "OpenCSV opens the Table data from a csv (comma-separated values) file (or any delim)", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"filename", "delim"}}, {Name: "SetColsByName", Doc: "SetColsByName turns cols On or Off if their name contains given string", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"nameContains", "on"}}}, Embeds: []types.Field{{Name: "Layout"}}, Fields: []types.Field{{Name: "Table", Doc: "the idxview of the table that we're plotting"}, {Name: "Params", Doc: "the overall plot parameters"}, {Name: "Cols", Doc: "the parameters for each column of the table"}, {Name: "Plot", Doc: "the gonum plot that actually does the plotting -- always save the last one generated"}, {Name: "ConfigPlotFunc", Doc: "ConfigPlotFunc is a function to call to configure [Plot2D.Plot], the gonum plot that\nactually does the plotting. It is called after [Plot] is generated, and properties\nof [Plot] can be modified in it. Properties of [Plot] should not be modified outside\nof this function, as doing so will have no effect."}, {Name: "SVGFile", Doc: "current svg file"}, {Name: "DataFile", Doc: "current csv data file"}, {Name: "InPlot", Doc: "currently doing a plot"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/eplot.PlotParams", IDName: "plot-params", Doc: "PlotParams are parameters for overall plot", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Title", Doc: "optional title at top of plot"}, {Name: "Type", Doc: "type of plot to generate.  For a Bar plot, items are plotted ordinally by row and the XAxis is optional"}, {Name: "Lines", Doc: "whether to plot lines"}, {Name: "Points", Doc: "whether to plot points with symbols"}, {Name: "LineWidth", Doc: "width of lines"}, {Name: "PointSize", Doc: "size of points"}, {Name: "PointShape", Doc: "the shape used to draw points"}, {Name: "BarWidth", Doc: "width of bars for bar plot, as fraction of available space (1 = no gaps)"}, {Name: "NegXDraw", Doc: "draw lines that connect points with a negative X-axis direction -- otherwise these are treated as breaks between repeated series and not drawn"}, {Name: "Scale", Doc: "overall scaling factor -- the larger the number, the larger the fonts are relative to the graph"}, {Name: "XAxisCol", Doc: "what column to use for the common X axis -- if empty or not found, the row number is used.  This optional for Bar plots -- if present and LegendCol is also present, then an extra space will be put between X values."}, {Name: "LegendCol", Doc: "optional column for adding a separate colored / styled line or bar according to this value -- acts just like a separate Y variable, crossed with Y variables"}, {Name: "XAxisRot", Doc: "rotation of the X Axis labels, in degrees"}, {Name: "XAxisLabel", Doc: "optional label to use for XAxis instead of column name"}, {Name: "YAxisLabel", Doc: "optional label to use for YAxis -- if empty, first column name is used"}, {Name: "Plot", Doc: "our plot, for update method"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/eplot.ColParams", IDName: "col-params", Doc: "ColParams are parameters for plotting one column of data", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "On", Doc: "whether to plot this column"}, {Name: "Col", Doc: "name of column we're plotting"}, {Name: "Lines", Doc: "whether to plot lines; uses the overall plot option if unset"}, {Name: "Points", Doc: "whether to plot points with symbols; uses the overall plot option if unset"}, {Name: "LineWidth", Doc: "the width of lines; uses the overall plot option if unset"}, {Name: "PointSize", Doc: "the size of points; uses the overall plot option if unset"}, {Name: "PointShape", Doc: "the shape used to draw points; uses the overall plot option if unset"}, {Name: "Range", Doc: "effective range of data to plot -- either end can be fixed"}, {Name: "FullRange", Doc: "full actual range of data -- only valid if specifically computed"}, {Name: "Color", Doc: "color to use when plotting the line / column"}, {Name: "NTicks", Doc: "desired number of ticks"}, {Name: "Lbl", Doc: "if non-empty, this is an alternative label to use in plotting"}, {Name: "TensorIndex", Doc: "if column has n-dimensional tensor cells in each row, this is the index within each cell to plot -- use -1 to plot *all* indexes as separate lines"}, {Name: "ErrCol", Doc: "specifies a column containing error bars for this column"}, {Name: "IsString", Doc: "if true this is a string column -- plots as labels"}, {Name: "Plot", Doc: "our plot, for update method"}}})
