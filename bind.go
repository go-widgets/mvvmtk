// Copyright (c) 2026 the go-widgets/mvvmtk authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package mvvmtk

import (
	"github.com/go-widgets/mvvm"
	"github.com/go-widgets/toolkit"
)

// ── Two-way scalar bindings (widget value field ⇄ Observable) ───────────────
//
// Each wraps mvvm.BindField with the widget's real value field and callback
// slot, so a ViewModel change flows to the field (and repaints via invalidate)
// and a user edit flows through the widget callback back to the observable.
// invalidate may be nil. The returned unbind restores the prior callback and
// detaches the subscription.

// BindText two-way-binds a SearchEntry's Text to a string observable
// (fields SearchEntry.Text / SearchEntry.OnChange).
func BindText(e *toolkit.SearchEntry, obs *mvvm.Observable[string], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &e.Text, &e.OnChange, invalidate)
}

// BindEntryText two-way-binds an Entry's Text to a string observable
// (fields Entry.Text / Entry.OnChange).
func BindEntryText(e *toolkit.Entry, obs *mvvm.Observable[string], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &e.Text, &e.OnChange, invalidate)
}

// BindChecked two-way-binds a CheckButton's Checked to a bool observable
// (fields CheckButton.Checked / CheckButton.OnToggle).
func BindChecked(cb *toolkit.CheckButton, obs *mvvm.Observable[bool], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &cb.Checked, &cb.OnToggle, invalidate)
}

// BindSelectedIndex two-way-binds a DropDown's Selected index to an int
// observable (fields DropDown.Selected / DropDown.OnSelect).
func BindSelectedIndex(d *toolkit.DropDown, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &d.Selected, &d.OnSelect, invalidate)
}

// BindSpin two-way-binds a SpinButton's Value to an int observable
// (fields SpinButton.Value / SpinButton.OnChange).
func BindSpin(s *toolkit.SpinButton, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &s.Value, &s.OnChange, invalidate)
}

// BindListSelection two-way-binds a ListBox's Selected row to an int observable
// (fields ListBox.Selected / ListBox.OnActivate). Note OnActivate fires on row
// activation (its only value/index callback slot), so the view→VM edge updates
// the observable when a row is activated; the VM→view edge sets Selected on any
// observable change.
func BindListSelection(lb *toolkit.ListBox, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &lb.Selected, &lb.OnActivate, invalidate)
}

// BindViewSwitcher two-way-binds a ViewSwitcher's Current segment to an int
// observable (fields ViewSwitcher.Current / ViewSwitcher.OnChange).
func BindViewSwitcher(v *toolkit.ViewSwitcher, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &v.Current, &v.OnChange, invalidate)
}

// ── One-way sinks (Observable → widget value field) ─────────────────────────
//
// For view-only widgets that carry no user-edit callback. Wraps mvvm.OneWay.

// BindLabel one-way-binds a string observable to a Label's Text
// (field Label.Text; a Label has no edit callback).
func BindLabel(l *toolkit.Label, obs *mvvm.Observable[string], invalidate func()) (unbind func()) {
	return mvvm.OneWay(obs, &l.Text, invalidate)
}

// BindProgress one-way-binds a float64 observable to a ProgressBar's Fraction
// (field ProgressBar.Fraction; a ProgressBar has no edit callback). Callers
// keep the value in the widget's [0,1] domain.
func BindProgress(p *toolkit.ProgressBar, obs *mvvm.Observable[float64], invalidate func()) (unbind func()) {
	return mvvm.OneWay(obs, &p.Fraction, invalidate)
}

// ── Collection bindings (ObservableList → widget []string source) ───────────
//
// Wraps mvvm.BindList: on every list change it rebuilds the widget's []string
// backing slice via project and calls invalidate.

// BindListItems binds an ObservableList to a ListBox's Items slice
// (field ListBox.Items), projecting each element to its row string.
func BindListItems[T any](lb *toolkit.ListBox, l *mvvm.ObservableList[T], project func(T) string, invalidate func()) (unbind func()) {
	return mvvm.BindList(l, &lb.Items, project, invalidate)
}

// BindDropDownOptions binds an ObservableList to a DropDown's Options slice
// (field DropDown.Options), projecting each element to its option string.
func BindDropDownOptions[T any](d *toolkit.DropDown, l *mvvm.ObservableList[T], project func(T) string, invalidate func()) (unbind func()) {
	return mvvm.BindList(l, &d.Options, project, invalidate)
}

// BindViews binds an ObservableList to a ViewSwitcher's Views slice
// (field ViewSwitcher.Views), projecting each element to its segment label.
func BindViews[T any](v *toolkit.ViewSwitcher, l *mvvm.ObservableList[T], project func(T) string, invalidate func()) (unbind func()) {
	return mvvm.BindList(l, &v.Views, project, invalidate)
}

// ── Command binding ─────────────────────────────────────────────────────────

// BindCommand wires a Command to a Button: it composes Execute into the
// button's OnClick and reflects executability by greying the button — swapping
// its Style to ButtonSecondary when the command cannot execute and restoring
// the original Style when it can (a Button has no boolean "disabled" field, so
// this backend-specific greying is how CanExecute surfaces). invalidate (may be
// nil) requests a repaint on each executability change. The returned unbind
// restores the prior OnClick and detaches.
func BindCommand(b *toolkit.Button, c *mvvm.Command, invalidate func()) (unbind func()) {
	orig := b.Style
	setEnabled := func(enabled bool) {
		if enabled {
			b.Style = orig
		} else {
			b.Style = toolkit.ButtonSecondary
		}
		if invalidate != nil {
			invalidate()
		}
	}
	return mvvm.BindCommand(c, &b.OnClick, setEnabled)
}

// ── Tree binding (ObservableList → TreeTable forest) ────────────────────────

// BindTree binds an ObservableList to a TreeTable's Root forest: on every list
// change it rebuilds Root (field TreeTable.Root, a []*TreeTableNode forest — not
// a flat []string, so mvvm.BindList can't drive it) by projecting each element
// to a node, then calls invalidate. project owns building each node's Cells and
// Children; this helper only wires the rebuild-on-change. The returned unbind
// detaches the list subscription.
func BindTree[T any](tt *toolkit.TreeTable, l *mvvm.ObservableList[T], project func(T) *toolkit.TreeTableNode, invalidate func()) (unbind func()) {
	rebuild := func() {
		src := l.Slice()
		nodes := make([]*toolkit.TreeTableNode, len(src))
		for i, v := range src {
			nodes[i] = project(v)
		}
		tt.Root = nodes
		if invalidate != nil {
			invalidate()
		}
	}
	rebuild()
	return l.Subscribe(func(mvvm.ListEvent[T]) { rebuild() })
}
