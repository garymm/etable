// Code generated by "core generate -add-types"; DO NOT EDIT.

package etview

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/etview.SimMatGrid", IDName: "sim-mat-grid", Doc: "SimMatGrid is a widget that displays a similarity / distance matrix\nwith tensor values as a grid of colored squares, and labels for rows, cols", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Embeds: []types.Field{{Name: "TensorGrid"}}, Fields: []types.Field{{Name: "SimMat", Doc: "the similarity / distance matrix"}, {Name: "rowMaxSz"}, {Name: "rowMinBlank"}, {Name: "rowNGps"}, {Name: "colMaxSz"}, {Name: "colMinBlank"}, {Name: "colNGps"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/etview.TableView", IDName: "table-view", Doc: "etview.TableView provides a GUI interface for etable.Table's", Embeds: []types.Field{{Name: "SliceViewBase"}}, Fields: []types.Field{{Name: "Table", Doc: "the idx view of the table that we're a view of"}, {Name: "TsrDisp", Doc: "overall display options for tensor display"}, {Name: "ColTsrDisp", Doc: "per column tensor display params"}, {Name: "ColTsrBlank", Doc: "per column blank tensor values"}, {Name: "NCols", Doc: "number of columns in table (as of last update)"}, {Name: "SortIndex", Doc: "current sort index"}, {Name: "SortDesc", Doc: "whether current sort order is descending"}, {Name: "HeaderWidths", Doc: "HeaderWidths has number of characters in each header, per visfields"}, {Name: "ColMaxWidths", Doc: "ColMaxWidths records maximum width in chars of string type fields"}, {Name: "BlankString", Doc: "\tblank values for out-of-range rows"}, {Name: "BlankFloat"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/etview.TensorLayout", IDName: "tensor-layout", Doc: "TensorLayout are layout options for displaying tensors", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "OddRow", Doc: "even-numbered dimensions are displayed as Y*X rectangles -- this determines along which dimension to display any remaining odd dimension: OddRow = true = organize vertically along row dimension, false = organize horizontally across column dimension"}, {Name: "TopZero", Doc: "if true, then the Y=0 coordinate is displayed from the top-down; otherwise the Y=0 coordinate is displayed from the bottom up, which is typical for emergent network patterns."}, {Name: "Image", Doc: "display the data as a bitmap image.  if a 2D tensor, then it will be a greyscale image.  if a 3D tensor with size of either the first or last dim = either 3 or 4, then it is a RGB(A) color image"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/etview.TensorDisp", IDName: "tensor-disp", Doc: "TensorDisp are options for displaying tensors", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Embeds: []types.Field{{Name: "TensorLayout"}}, Fields: []types.Field{{Name: "Range", Doc: "range to plot"}, {Name: "MinMax", Doc: "if not using fixed range, this is the actual range of data"}, {Name: "ColorMap", Doc: "the name of the color map to use in translating values to colors"}, {Name: "GridFill", Doc: "what proportion of grid square should be filled by color block -- 1 = all, .5 = half, etc"}, {Name: "DimExtra", Doc: "amount of extra space to add at dimension boundaries, as a proportion of total grid size"}, {Name: "GridMinSize", Doc: "minimum size for grid squares -- they will never be smaller than this"}, {Name: "GridMaxSize", Doc: "maximum size for grid squares -- they will never be larger than this"}, {Name: "TotPrefSize", Doc: "total preferred display size along largest dimension.\ngrid squares will be sized to fit within this size,\nsubject to harder GridMin / Max size constraints"}, {Name: "FontSize", Doc: "font size in standard point units for labels (e.g., SimMat)"}, {Name: "GridView", Doc: "our gridview, for update method"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/etview.TensorGrid", IDName: "tensor-grid", Doc: "TensorGrid is a widget that displays tensor values as a grid of colored squares.", Methods: []types.Method{{Name: "EditSettings", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}}, Embeds: []types.Field{{Name: "WidgetBase"}}, Fields: []types.Field{{Name: "Tensor", Doc: "the tensor that we view"}, {Name: "Disp", Doc: "display options"}, {Name: "ColorMap", Doc: "the actual colormap"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/etview.TensorGridValue", IDName: "tensor-grid-value", Doc: "TensorGridValue manages a TensorGrid view of an etensor.Tensor", Embeds: []types.Field{{Name: "ValueBase"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/etview.TensorValue", IDName: "tensor-value", Doc: "TensorValue presents a button that pulls up the TensorView viewer for an etensor.Tensor", Embeds: []types.Field{{Name: "ValueBase"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/etview.TableValue", IDName: "table-value", Doc: "TableValue presents a button that pulls up the TableView viewer for an etable.Table", Embeds: []types.Field{{Name: "ValueBase"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/etable/v2/etview.SimMatValue", IDName: "sim-mat-value", Doc: "SimMatValue presents a button that pulls up the SimMatGridView viewer for an etable.Table", Embeds: []types.Field{{Name: "ValueBase"}}})
