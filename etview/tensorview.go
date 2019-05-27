// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package etview

import (
	"fmt"
	"image"
	"reflect"

	"github.com/chewxy/math32"
	"github.com/emer/etable/etensor"
	"github.com/goki/gi/gi"
	"github.com/goki/gi/giv"
	"github.com/goki/gi/oswin/mimedata"
	"github.com/goki/gi/units"
	"github.com/goki/ki/ints"
	"github.com/goki/ki/ki"
	"github.com/goki/ki/kit"
)

// etview.TensorView provides a GUI interface for etable.Tensor's
// using a tabular rows-and-columns interface
type TensorView struct {
	giv.SliceViewBase
	Tensor etensor.Tensor `desc:"the tensor that we're a view of"`
	TsrLay TensorLayout   `desc:"layout config of the tensor"`
	NCols  int            `inactive:"+" desc:"number of columns in table (as of last update)"`
}

var KiT_TensorView = kit.Types.AddType(&TensorView{}, TensorViewProps)

// AddNewTensorView adds a new tableview to given parent node, with given name.
func AddNewTensorView(parent ki.Ki, name string) *TensorView {
	return parent.AddNewChild(KiT_TensorView, name).(*TensorView)
}

// check for interface impl
var _ giv.SliceViewer = (*TensorView)(nil)

// SetTensor sets the source tensor that we are viewing
func (tv *TensorView) SetTensor(tsr etensor.Tensor, tmpSave giv.ValueView) {
	updt := false
	if tsr == nil {
		return
	}
	if tv.Tensor != tsr {
		if !tv.IsInactive() {
			tv.SelectedIdx = -1
		}
		tv.StartIdx = 0
		tv.Tensor = tsr
		updt = tv.UpdateStart()
		tv.ResetSelectedIdxs()
		tv.SelectMode = false
		tv.SetFullReRender()
	}
	tv.ShowIndex = true
	if sidxp, err := tv.PropTry("index"); err == nil {
		tv.ShowIndex, _ = kit.ToBool(sidxp)
	}
	tv.InactKeyNav = true
	if siknp, err := tv.PropTry("inact-key-nav"); err == nil {
		tv.InactKeyNav, _ = kit.ToBool(siknp)
	}
	tv.TmpSave = tmpSave
	tv.Config()
	tv.UpdateEnd(updt)
}

var TensorViewProps = ki.Props{
	"background-color": &gi.Prefs.Colors.Background,
	"color":            &gi.Prefs.Colors.Font,
	"max-width":        -1,
	"max-height":       -1,
}

// IsConfiged returns true if the widget is fully configured
func (tv *TensorView) IsConfiged() bool {
	if len(tv.Kids) == 0 {
		return false
	}
	sf := tv.SliceFrame()
	if len(sf.Kids) == 0 {
		return false
	}
	return true
}

// Config configures the view
func (tv *TensorView) Config() {
	tv.Lay = gi.LayoutVert
	tv.SetProp("spacing", gi.StdDialogVSpaceUnits)
	config := kit.TypeAndNameList{}
	config.Add(gi.KiT_ToolBar, "toolbar")
	config.Add(gi.KiT_Frame, "frame")
	mods, updt := tv.ConfigChildren(config, true)
	tv.ConfigSliceGrid()
	tv.ConfigToolbar()
	if mods {
		tv.SetFullReRender()
		tv.UpdateEnd(updt)
	}
}

func (tv *TensorView) UpdtSliceSize() int {
	tv.SliceSize, tv.NCols = etensor.Prjn2DShape(tv.Tensor, tv.TsrLay.OddRow)
	return tv.SliceSize
}

// SliceFrame returns the outer frame widget, which contains all the header,
// fields and values
func (tv *TensorView) SliceFrame() *gi.Frame {
	return tv.ChildByName("frame", 0).(*gi.Frame)
}

// GridLayout returns the SliceGrid grid-layout widget, with grid and scrollbar
func (tv *TensorView) GridLayout() *gi.Layout {
	return tv.SliceFrame().ChildByName("grid-lay", 0).(*gi.Layout)
}

// SliceGrid returns the SliceGrid grid frame widget, which contains all the
// fields and values, within SliceFrame
func (tv *TensorView) SliceGrid() *gi.Frame {
	return tv.GridLayout().ChildByName("grid", 0).(*gi.Frame)
}

// ScrollBar returns the SliceGrid scrollbar
func (tv *TensorView) ScrollBar() *gi.ScrollBar {
	return tv.GridLayout().ChildByName("scrollbar", 1).(*gi.ScrollBar)
}

// SliceHeader returns the Toolbar header for slice grid
func (tv *TensorView) SliceHeader() *gi.ToolBar {
	return tv.SliceFrame().Child(0).(*gi.ToolBar)
}

// ToolBar returns the toolbar widget
func (tv *TensorView) ToolBar() *gi.ToolBar {
	return tv.ChildByName("toolbar", 0).(*gi.ToolBar)
}

// RowWidgetNs returns number of widgets per row and offset for index label
func (tv *TensorView) RowWidgetNs() (nWidgPerRow, idxOff int) {
	nWidgPerRow = 1 + tv.NCols
	if !tv.IsInactive() {
		if !tv.NoAdd {
			nWidgPerRow += 1
		}
		if !tv.NoDelete {
			nWidgPerRow += 1
		}
	}
	idxOff = 1
	if !tv.ShowIndex {
		nWidgPerRow -= 1
		idxOff = 0
	}
	return
}

// ConfigSliceGrid configures the SliceGrid for the current slice
// this is only called by global Config and updates are guarded by that
func (tv *TensorView) ConfigSliceGrid() {
	if tv.Tensor == nil {
		return
	}

	sz := tv.UpdtSliceSize()
	if sz == 0 {
		return
	}

	nWidgPerRow, idxOff := tv.RowWidgetNs()

	sg := tv.SliceFrame()
	updt := sg.UpdateStart()
	defer sg.UpdateEnd(updt)

	sg.Lay = gi.LayoutVert
	sg.SetMinPrefWidth(units.NewEm(10))
	sg.SetStretchMaxHeight() // for this to work, ALL layers above need it too
	sg.SetStretchMaxWidth()  // for this to work, ALL layers above need it too

	sgcfg := kit.TypeAndNameList{}
	sgcfg.Add(gi.KiT_ToolBar, "header")
	sgcfg.Add(gi.KiT_Layout, "grid-lay")
	sg.ConfigChildren(sgcfg, true)

	sgh := tv.SliceHeader()
	sgh.Lay = gi.LayoutHoriz
	sgh.SetProp("overflow", gi.OverflowHidden) // no scrollbars!
	sgh.SetProp("spacing", 0)
	// sgh.SetStretchMaxWidth()

	gl := tv.GridLayout()
	gl.Lay = gi.LayoutHoriz
	gl.SetStretchMaxHeight() // for this to work, ALL layers above need it too
	gl.SetStretchMaxWidth()  // for this to work, ALL layers above need it too
	gconfig := kit.TypeAndNameList{}
	gconfig.Add(gi.KiT_Frame, "grid")
	gconfig.Add(gi.KiT_ScrollBar, "scrollbar")
	gl.ConfigChildren(gconfig, true) // covered by above

	sgf := tv.SliceGrid()
	sgf.Lay = gi.LayoutGrid
	sgf.Stripes = gi.RowStripes
	sgf.SetMinPrefHeight(units.NewEm(10))
	sgf.SetStretchMaxHeight() // for this to work, ALL layers above need it too
	sgf.SetStretchMaxWidth()  // for this to work, ALL layers above need it too
	sgf.SetProp("columns", nWidgPerRow)

	// Configure Header
	hcfg := kit.TypeAndNameList{}
	if tv.ShowIndex {
		hcfg.Add(gi.KiT_Label, "head-idx")
	}
	for fli := 0; fli < tv.NCols; fli++ {
		labnm := fmt.Sprintf("head-%03d", fli)
		hcfg.Add(gi.KiT_Label, labnm)
	}
	if !tv.IsInactive() {
		hcfg.Add(gi.KiT_Label, "head-add")
		hcfg.Add(gi.KiT_Label, "head-del")
	}
	sgh.ConfigChildren(hcfg, false) // headers SHOULD be unique, but with labels..

	// at this point, we make one dummy row to get size of widgets

	sgf.DeleteChildren(true)
	sgf.Kids = make(ki.Slice, nWidgPerRow)

	itxt := fmt.Sprintf("%05d", 0)
	labnm := fmt.Sprintf("index-%v", itxt)

	if tv.ShowIndex {
		lbl := sgh.Child(0).(*gi.Label)
		lbl.Text = "Index"

		idxlab := &gi.Label{}
		sgf.SetChild(idxlab, 0, labnm)
		idxlab.Text = itxt
	}

	for fli := 0; fli < tv.NCols; fli++ {
		colnm := fmt.Sprintf("X:%03d", fli) // todo: deal with embedded dims
		hdr := sgh.Child(idxOff + fli).(*gi.Label)
		hdr.SetText(colnm)
		// no metadata on tensors by themselves
		// if dsc, has := tv.Tensor.MetaData[colnm+":desc"]; has {
		// 	hdr.Tooltip += ": " + dsc
		// }

		fval := 1.0
		vv := giv.ToValueView(&fval, "")
		vv.SetStandaloneValue(reflect.ValueOf(&fval))
		vtyp := vv.WidgetType()
		valnm := fmt.Sprintf("value-%v.%v", fli, itxt)
		cidx := idxOff + fli
		widg := ki.NewOfType(vtyp).(gi.Node2D)
		sgf.SetChild(widg, cidx, valnm)
		vv.ConfigWidget(widg)
	}

	if !tv.IsInactive() {
		cidx := tv.NCols + idxOff
		if !tv.NoAdd {
			lbl := sgh.Child(cidx).(*gi.Label)
			lbl.Text = "+"
			lbl.Tooltip = "insert row"
			addnm := fmt.Sprintf("add-%v", itxt)
			addact := gi.Action{}
			sgf.SetChild(&addact, cidx, addnm)
			addact.SetIcon("plus")
			cidx++
		}
		if !tv.NoDelete {
			lbl := sgh.Child(cidx).(*gi.Label)
			lbl.Text = "-"
			lbl.Tooltip = "delete row"
			delnm := fmt.Sprintf("del-%v", itxt)
			delact := gi.Action{}
			sgf.SetChild(&delact, cidx, delnm)
			delact.SetIcon("minus")
			cidx++
		}
	}

	tv.ConfigScroll()
}

// LayoutSliceGrid does the proper layout of slice grid depending on allocated size
// returns true if UpdateSliceGrid should be called after this
func (tv *TensorView) LayoutSliceGrid() bool {
	sg := tv.SliceGrid()
	if tv.Tensor == nil {
		sg.DeleteChildren(true)
		return false
	}
	sz := tv.UpdtSliceSize()
	if sz == 0 {
		sg.DeleteChildren(true)
		return false
	}

	sgHt := tv.AvailHeight()
	tv.LayoutHeight = sgHt
	if sgHt == 0 {
		return false
	}

	nWidgPerRow, _ := tv.RowWidgetNs()
	tv.RowHeight = sg.GridData[gi.Row][0].AllocSize + sg.Spacing.Dots
	tv.VisRows = int(math32.Floor(sgHt / tv.RowHeight))
	tv.DispRows = ints.MinInt(tv.SliceSize, tv.VisRows)

	nWidg := nWidgPerRow * tv.DispRows

	updt := sg.UpdateStart()
	defer sg.UpdateEnd(updt)
	if tv.Values == nil || sg.NumChildren() != nWidg {
		sg.DeleteChildren(true)

		tv.Values = make([]giv.ValueView, tv.NCols*tv.DispRows)
		sg.Kids = make(ki.Slice, nWidg)
	}
	tv.ConfigScroll()
	tv.LayoutHeader()
	return true
}

// LayoutHeader updates the header layout based on field widths
func (tv *TensorView) LayoutHeader() {
	nWidgPerRow, idxOff := tv.RowWidgetNs()
	nfld := tv.NCols + idxOff
	sgh := tv.SliceHeader()
	sgf := tv.SliceGrid()
	spc := sgf.Spacing.Dots
	if len(sgf.Kids) >= nfld {
		sumwd := float32(0)
		for fli := 0; fli < nfld; fli++ {
			lbl := sgh.Child(fli).(gi.Node2D).AsWidget()
			wd := sgf.GridData[gi.Col][fli].AllocSize
			lbl.SetMinPrefWidth(units.NewValue(wd+spc, units.Dot))
			sumwd += wd + spc
		}
		if !tv.IsInactive() {
			for fli := nfld; fli < nWidgPerRow; fli++ {
				lbl := sgh.Child(fli).(gi.Node2D).AsWidget()
				wd := sgf.GridData[gi.Col][fli].AllocSize
				lbl.SetMinPrefWidth(units.NewValue(wd+spc, units.Dot))
				sumwd += wd + spc
			}
		}
		sgh.SetMinPrefWidth(units.NewValue(sumwd, units.Dot))
	}
}

// UpdateSliceGrid updates grid display -- robust to any time calling
func (tv *TensorView) UpdateSliceGrid() {
	if tv.Tensor == nil {
		return
	}
	sz := tv.UpdtSliceSize()
	if sz == 0 {
		return
	}
	sg := tv.SliceGrid()
	tv.DispRows = ints.MinInt(tv.SliceSize, tv.VisRows)

	nWidgPerRow, idxOff := tv.RowWidgetNs()
	nWidg := nWidgPerRow * tv.DispRows

	if tv.Viewport != nil && tv.Viewport.Win != nil {
		wupdt := tv.Viewport.Win.UpdateStart()
		defer tv.Viewport.Win.UpdateEnd(wupdt)
	}

	updt := sg.UpdateStart()
	defer sg.UpdateEnd(updt)

	if tv.Values == nil || sg.NumChildren() != nWidg { // shouldn't happen..
		tv.LayoutSliceGrid()
		nWidg = nWidgPerRow * tv.DispRows
	}

	if sz > tv.DispRows {
		sb := tv.ScrollBar()
		tv.StartIdx = int(sb.Value)
		lastSt := sz - tv.DispRows
		tv.StartIdx = ints.MinInt(lastSt, tv.StartIdx)
		tv.StartIdx = ints.MaxInt(0, tv.StartIdx)
	} else {
		tv.StartIdx = 0
	}

	for i := 0; i < tv.DispRows; i++ {
		ridx := i * nWidgPerRow
		si := tv.StartIdx + i // slice idx
		if !tv.TsrLay.TopZero {
			si = (tv.SliceSize - 1) - si
		}
		issel := tv.IdxIsSelected(si)

		itxt := fmt.Sprintf("%05d", i)
		sitxt := fmt.Sprintf("%05d", si)
		labnm := fmt.Sprintf("index-%v", itxt)
		if tv.ShowIndex {
			var idxlab *gi.Label
			if sg.Kids[ridx] != nil {
				idxlab = sg.Kids[ridx].(*gi.Label)
			} else {
				idxlab = &gi.Label{}
				sg.SetChild(idxlab, ridx, labnm)
				idxlab.SetProp("tv-row", i)
				idxlab.Selectable = true
				idxlab.Redrawable = true
				idxlab.Sty.Template = "View.IndexLabel"
				idxlab.WidgetSig.ConnectOnly(tv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
					if sig == int64(gi.WidgetSelected) {
						wbb := send.(gi.Node2D).AsWidget()
						row := wbb.Prop("tv-row").(int)
						tvv := recv.Embed(KiT_TensorView).(*TensorView)
						tvv.UpdateSelectRow(row, wbb.IsSelected())
					}
				})
			}
			idxlab.CurBgColor = gi.Prefs.Colors.Background
			idxlab.SetText(sitxt)
			idxlab.SetSelectedState(issel)
		}

		for fli := 0; fli < tv.NCols; fli++ {
			fval := etensor.Prjn2DVal(tv.Tensor, tv.TsrLay.OddRow, si, fli)
			vvi := i*tv.NCols + fli
			var vv giv.ValueView
			if tv.Values[vvi] == nil {
				vv = giv.ToValueView(&fval, "")
				vv.SetStandaloneValue(reflect.ValueOf(&fval))
				tv.Values[vvi] = vv
				vv.SetProp("tv-row", i)
				vv.SetProp("tv-col", fli)
			} else {
				vv = tv.Values[vvi]
				vv.SetStandaloneValue(reflect.ValueOf(&fval))
			}

			vtyp := vv.WidgetType()
			valnm := fmt.Sprintf("value-%v.%v", fli, itxt)
			cidx := ridx + idxOff + fli
			var widg gi.Node2D
			if sg.Kids[cidx] != nil {
				widg = sg.Kids[cidx].(gi.Node2D)
				vv.UpdateWidget()
				if tv.IsInactive() {
					widg.AsNode2D().SetInactive()
				}
				widg.AsNode2D().SetSelectedState(issel)
			} else {
				widg = ki.NewOfType(vtyp).(gi.Node2D)
				sg.SetChild(widg, cidx, valnm)
				vv.ConfigWidget(widg)
				wb := widg.AsWidget()
				if wb != nil {
					wb.SetProp("tv-row", i)
					wb.ClearSelected()
					wb.WidgetSig.ConnectOnly(tv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
						if sig == int64(gi.WidgetSelected) || sig == int64(gi.WidgetFocused) {
							wbb := send.(gi.Node2D).AsWidget()
							row := wbb.Prop("tv-row").(int)
							tvv := recv.Embed(KiT_TensorView).(*TensorView)
							if sig != int64(gi.WidgetFocused) || !tvv.InFocusGrab {
								tvv.UpdateSelectRow(row, wbb.IsSelected())
							}
						}
					})
				}
				if tv.IsInactive() {
					widg.AsNode2D().SetInactive()
				} else {
					vvb := vv.AsValueViewBase()
					vvb.ViewSig.ConnectOnly(tv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
						tvv, _ := recv.Embed(KiT_TensorView).(*TensorView)
						tvv.SetChanged()
						vvv := send.(giv.ValueView).AsValueViewBase()
						row := vvv.Prop("tv-row").(int)
						col := vvv.Prop("tv-col").(int)
						npv := kit.NonPtrValue(vvv.Value)
						fv, ok := kit.ToFloat(npv.Interface())
						if ok {
							etensor.Prjn2DSet(tv.Tensor, tv.TsrLay.OddRow, tvv.StartIdx+row, col, fv)
							tvv.ViewSig.Emit(tvv.This(), 0, nil)
						}
					})
				}
			}
		}

		if !tv.IsInactive() {
			cidx := ridx + tv.NCols + idxOff
			if !tv.NoAdd {
				if sg.Kids[cidx] == nil {
					addnm := fmt.Sprintf("add-%v", itxt)
					addact := gi.Action{}
					sg.SetChild(&addact, cidx, addnm)
					addact.SetIcon("plus")
					addact.Tooltip = "insert a new element at this index"
					addact.Data = i
					addact.Sty.Template = "etview.TensorView.AddAction"
					addact.ActionSig.ConnectOnly(tv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
						act := send.(*gi.Action)
						tvv := recv.Embed(KiT_TableView).(*TableView)
						tvv.SliceNewAtRow(act.Data.(int) + 1)
					})
				}
				cidx++
			}
			if !tv.NoDelete {
				if sg.Kids[cidx] == nil {
					delnm := fmt.Sprintf("del-%v", itxt)
					delact := gi.Action{}
					sg.SetChild(&delact, cidx, delnm)
					delact.SetIcon("minus")
					delact.Tooltip = "delete this element"
					delact.Data = i
					delact.Sty.Template = "etview.TensorView.DelAction"
					delact.ActionSig.ConnectOnly(tv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
						act := send.(*gi.Action)
						tvv := recv.Embed(KiT_TableView).(*TableView)
						tvv.SliceDeleteAtRow(act.Data.(int), true)
					})
				}
				cidx++
			}
		}
	}

	if tv.IsInactive() && tv.SelectedIdx >= 0 {
		tv.SelectIdx(tv.SelectedIdx)
	}
	tv.UpdateScroll()
}

func (tv *TensorView) StyleRow(svnp reflect.Value, widg gi.Node2D, idx, fidx int, vv giv.ValueView) {
}

// SliceNewAt inserts a new blank element at given index in the slice -- -1
// means the end
func (tv *TensorView) SliceNewAt(idx int) {
	wupdt := tv.Viewport.Win.UpdateStart()
	defer tv.Viewport.Win.UpdateEnd(wupdt)

	updt := tv.UpdateStart()
	defer tv.UpdateEnd(updt)

	// todo: insert row -- do we even have this??  no!
	// kit.SliceNewAt(tv.Slice, idx)

	if tv.TmpSave != nil {
		tv.TmpSave.SaveTmp()
	}
	tv.SetChanged()
	tv.This().(giv.SliceViewer).LayoutSliceGrid()
	tv.This().(giv.SliceViewer).UpdateSliceGrid()
	tv.ViewSig.Emit(tv.This(), 0, nil)
}

// SliceDeleteAt deletes element at given index from slice -- doupdt means
// call UpdateSliceGrid to update display
func (tv *TensorView) SliceDeleteAt(idx int, doupdt bool) {
	if idx < 0 {
		return
	}
	wupdt := tv.Viewport.Win.UpdateStart()
	defer tv.Viewport.Win.UpdateEnd(wupdt)

	updt := tv.UpdateStart()
	defer tv.UpdateEnd(updt)

	// kit.SliceDeleteAt(tv.Slice, idx)

	if tv.TmpSave != nil {
		tv.TmpSave.SaveTmp()
	}
	tv.SetChanged()
	if doupdt {
		tv.This().(giv.SliceViewer).LayoutSliceGrid()
		tv.This().(giv.SliceViewer).UpdateSliceGrid()
	}
	tv.ViewSig.Emit(tv.This(), 0, nil)
}

// ConfigToolbar configures the toolbar actions
func (tv *TensorView) ConfigToolbar() {
	if tv.Tensor == nil {
		return
	}
	if tv.ToolbarSlice == tv.Tensor {
		return
	}
	tb := tv.ToolBar()
	if len(*tb.Children()) == 0 {
		tb.SetStretchMaxWidth()
		tb.AddAction(gi.ActOpts{Label: "Updt", Icon: "update"},
			tv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
				tvv := recv.Embed(KiT_TensorView).(*TensorView)
				tvv.Update()
			})
		tb.AddAction(gi.ActOpts{Label: "Config", Icon: "gear"},
			tv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
				tvv := recv.Embed(KiT_TensorView).(*TensorView)
				giv.StructViewDialog(tv.Viewport, &tvv.TsrLay, giv.DlgOpts{Title: "TensorView Display Options", Ok: true, Cancel: true},
					tv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
						tvvv := recv.Embed(KiT_TensorView).(*TensorView)
						tvvv.UpdateSliceGrid()
					})
			})
	}
	nCustom := 2
	sz := len(*tb.Children())
	if sz > nCustom {
		for i := sz - 1; i >= nCustom; i-- {
			tb.DeleteChildAtIndex(i, true)
		}
	}
	if giv.HasToolBarView(tv.Slice) {
		giv.ToolBarView(tv.Slice, tv.Viewport, tb)
	}
	tv.ToolbarSlice = tv.Tensor
}

func (tv *TensorView) Layout2D(parBBox image.Rectangle, iter int) bool {
	redo := tv.Frame.Layout2D(parBBox, iter)
	if !tv.IsConfiged() {
		return redo
	}
	tv.LayoutHeader()
	tv.SliceHeader().Layout2D(parBBox, iter)
	return redo
}

// RowFirstVisWidget returns the first visible widget for given row (could be
// index or not) -- false if out of range
func (tv *TensorView) RowFirstVisWidget(row int) (*gi.WidgetBase, bool) {
	if !tv.IsRowInBounds(row) {
		return nil, false
	}
	nWidgPerRow, idxOff := tv.RowWidgetNs()
	sg := tv.SliceGrid()
	widg := sg.Kids[row*nWidgPerRow].(gi.Node2D).AsWidget()
	if widg.VpBBox != image.ZR {
		return widg, true
	}
	ridx := nWidgPerRow * row
	for fli := 0; fli < tv.NCols; fli++ {
		widg := sg.Child(ridx + idxOff + fli).(gi.Node2D).AsWidget()
		if widg.VpBBox != image.ZR {
			return widg, true
		}
	}
	return nil, false
}

// RowGrabFocus grabs the focus for the first focusable widget in given row --
// returns that element or nil if not successful -- note: grid must have
// already rendered for focus to be grabbed!
func (tv *TensorView) RowGrabFocus(row int) *gi.WidgetBase {
	if !tv.IsRowInBounds(row) || tv.InFocusGrab { // range check
		return nil
	}
	nWidgPerRow, idxOff := tv.RowWidgetNs()
	ridx := nWidgPerRow * row
	sg := tv.SliceGrid()
	// first check if we already have focus
	for fli := 0; fli < tv.NCols; fli++ {
		widg := sg.Child(ridx + idxOff + fli).(gi.Node2D).AsWidget()
		if widg.HasFocus() || widg.ContainsFocus() {
			return widg
		}
	}
	tv.InFocusGrab = true
	defer func() { tv.InFocusGrab = false }()
	for fli := 0; fli < tv.NCols; fli++ {
		widg := sg.Child(ridx + idxOff + fli).(gi.Node2D).AsWidget()
		if widg.CanFocus() {
			widg.GrabFocus()
			return widg
		}
	}
	return nil
}

// SelectRowWidgets sets the selection state of given row of widgets
func (tv *TensorView) SelectRowWidgets(row int, sel bool) {
	if row < 0 {
		return
	}
	wupdt := tv.Viewport.Win.UpdateStart()
	defer tv.Viewport.Win.UpdateEnd(wupdt)

	sg := tv.SliceGrid()
	nWidgPerRow, idxOff := tv.RowWidgetNs()
	ridx := row * nWidgPerRow
	for fli := 0; fli < tv.NCols; fli++ {
		seldx := ridx + idxOff + fli
		if sg.Kids.IsValidIndex(seldx) == nil {
			widg := sg.Child(seldx).(gi.Node2D).AsNode2D()
			widg.SetSelectedState(sel)
			widg.UpdateSig()
		}
	}
	if tv.ShowIndex {
		if sg.Kids.IsValidIndex(ridx) == nil {
			widg := sg.Child(ridx).(gi.Node2D).AsNode2D()
			widg.SetSelectedState(sel)
			widg.UpdateSig()
		}
	}
}

// CopySelToMime copies selected rows to mime data
func (tv *TensorView) CopySelToMime() mimedata.Mimes {
	return nil
}

// PasteAssign assigns mime data (only the first one!) to this idx
func (tv *TensorView) PasteAssign(md mimedata.Mimes, idx int) {
	// todo
}

// PasteAtIdx inserts object(s) from mime data at (before) given slice index
func (tv *TensorView) PasteAtIdx(md mimedata.Mimes, idx int) {
	// todo
}

func (tv *TensorView) ItemCtxtMenu(idx int) {
}

// // SelectFieldVal sets SelField and SelVal and attempts to find corresponding
// // row, setting SelectedIdx and selecting row if found -- returns true if
// // found, false otherwise
// func (tv *TensorView) SelectFieldVal(fld, val string) bool {
// 	tv.SelField = fld
// 	tv.SelVal = val
// 	if tv.SelField != "" && tv.SelVal != nil {
// 		idx, _ := StructSliceIdxByValue(tv.Slice, tv.SelField, tv.SelVal)
// 		if idx >= 0 {
// 			tv.ScrollToIdx(idx)
// 			tv.UpdateSelectIdx(idx, true)
// 			return true
// 		}
// 	}
// 	return false
// }