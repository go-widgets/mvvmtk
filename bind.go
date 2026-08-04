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

// BindSwitch two-way-binds a Switch's On to a bool observable
// (fields Switch.On / Switch.OnToggle).
func BindSwitch(s *toolkit.Switch, obs *mvvm.Observable[bool], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &s.On, &s.OnToggle, invalidate)
}

// BindToggle two-way-binds a ToggleButton's Pressed to a bool observable
// (fields ToggleButton.Pressed / ToggleButton.OnToggle).
func BindToggle(t *toolkit.ToggleButton, obs *mvvm.Observable[bool], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &t.Pressed, &t.OnToggle, invalidate)
}

// BindRadio two-way-binds a standalone RadioButton's Checked to a bool
// observable (fields RadioButton.Checked / RadioButton.OnToggle). For a set of
// mutually-exclusive radios use BindRadioGroup, which binds the group's Active
// index; this helper is for a single, group-less radio behaving like a check.
func BindRadio(r *toolkit.RadioButton, obs *mvvm.Observable[bool], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &r.Checked, &r.OnToggle, invalidate)
}

// BindRating two-way-binds a Rating's Value to an int observable
// (fields Rating.Value / Rating.OnChange).
func BindRating(r *toolkit.Rating, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &r.Value, &r.OnChange, invalidate)
}

// BindScale two-way-binds a Scale's Value to a float64 observable
// (fields Scale.Value / Scale.OnChange).
func BindScale(s *toolkit.Scale, obs *mvvm.Observable[float64], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &s.Value, &s.OnChange, invalidate)
}

// BindExpander two-way-binds an Expander's Expanded to a bool observable
// (fields Expander.Expanded / Expander.OnExpand).
func BindExpander(e *toolkit.Expander, obs *mvvm.Observable[bool], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &e.Expanded, &e.OnExpand, invalidate)
}

// BindNotebook two-way-binds a Notebook's Active tab to an int observable
// (fields Notebook.Active / Notebook.OnTabChanged).
func BindNotebook(n *toolkit.Notebook, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &n.Active, &n.OnTabChanged, invalidate)
}

// BindComboText two-way-binds a ComboBox's Text to a string observable
// (fields ComboBox.Text / ComboBox.OnChange).
func BindComboText(c *toolkit.ComboBox, obs *mvvm.Observable[string], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &c.Text, &c.OnChange, invalidate)
}

// BindPagination two-way-binds a Pagination's Current page to an int observable
// (fields Pagination.Current / Pagination.OnChange).
func BindPagination(pg *toolkit.Pagination, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &pg.Current, &pg.OnChange, invalidate)
}

// BindPagingToolbar two-way-binds a PagingToolbar's Page to an int observable
// (fields PagingToolbar.Page / PagingToolbar.OnChange).
func BindPagingToolbar(pt *toolkit.PagingToolbar, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &pt.Page, &pt.OnChange, invalidate)
}

// BindGanttSelection two-way-binds a Gantt's Selected task to an int observable
// (fields Gantt.Selected / Gantt.OnSelect).
func BindGanttSelection(g *toolkit.Gantt, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &g.Selected, &g.OnSelect, invalidate)
}

// BindTableSelection two-way-binds a Table's Selected row to an int observable
// (fields Table.Selected / Table.OnSelect). OnSelect fires on a MultiSelect
// anchor-moving click and on a keyboard cursor move, so the view→VM edge tracks
// either interaction; the VM→view edge sets Selected on any observable change.
func BindTableSelection(t *toolkit.Table, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &t.Selected, &t.OnSelect, invalidate)
}

// BindCarousel two-way-binds a Carousel's Current slide to an int observable
// (fields Carousel.Current / Carousel.OnChange).
func BindCarousel(c *toolkit.Carousel, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &c.Current, &c.OnChange, invalidate)
}

// BindCycle two-way-binds a CycleButton's Index to an int observable (fields
// CycleButton.Index / CycleButton.OnChangeIndex). It binds the single-argument
// OnChangeIndex slot; the multi-argument OnChange (index, value) is left free
// for a host that also wants the value.
func BindCycle(c *toolkit.CycleButton, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &c.Index, &c.OnChangeIndex, invalidate)
}

// BindRadioGroup two-way-binds a RadioGroup's Active member to an int observable
// (fields RadioGroup.Active / RadioGroup.OnChange). This binds the whole group
// (which member is checked); BindRadio binds a single group-less radio.
func BindRadioGroup(g *toolkit.RadioGroup, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &g.Active, &g.OnChange, invalidate)
}

// BindColor two-way-binds a ColorChooser's Color to an RGBA observable (fields
// ColorChooser.Color / ColorChooser.OnChange). ColorPicker, which exposes no
// Color field, is bound with BindColorPicker instead.
func BindColor(c *toolkit.ColorChooser, obs *mvvm.Observable[toolkit.RGBA], invalidate func()) (unbind func()) {
	return mvvm.BindField(obs, &c.Color, &c.OnChange, invalidate)
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

// ── Accessor / multi-value bindings ─────────────────────────────────────────
//
// These widgets cannot be driven by mvvm.BindField: either their change
// callback carries more than one argument (so its signature is not func(T)), or
// their value is reached through accessor methods rather than an exported field,
// or the observed value is a tuple/struct that no single field holds. Each helper
// hand-rolls the same three-part contract BindField provides — seed the widget
// from the observable, compose (not clobber) the prior callback so a user edit
// flows to the observable, and push observable→widget on change (invalidate may
// be nil) — and returns an unbind that restores the prior callback and detaches.

// Date is a plain calendar date, the value BindDate / BindDatePicker observe. It
// is comparable, so mvvm.NewObservable[Date] works and equal Sets are deduped.
type Date struct {
	Year  int
	Month int // 1..12
	Day   int
}

// TimeOfDay is a 24-hour wall-clock time, the value BindTime observes. It is
// comparable, so mvvm.NewObservable[TimeOfDay] works and equal Sets are deduped.
type TimeOfDay struct {
	Hour   int // 0..23
	Minute int // 0..59
}

// BindAccordion two-way-binds an Accordion's Expanded section index to an int
// observable (field Accordion.Expanded, callback Accordion.OnToggle). OnToggle's
// (i, expanded) signature is not func(int), so BindField cannot drive it; this
// helper composes OnToggle and, after each toggle, publishes the resulting
// Expanded index. It binds the exclusive-mode single-expanded index (-1 when all
// are collapsed); a Multiple-mode accordion does not track a single index.
func BindAccordion(a *toolkit.Accordion, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	a.Expanded = obs.Get()
	prev := a.OnToggle
	a.OnToggle = func(i int, expanded bool) {
		if prev != nil {
			prev(i, expanded)
		}
		obs.Set(a.Expanded)
	}
	unsub := obs.Subscribe(func(v int) {
		a.Expanded = v
		if invalidate != nil {
			invalidate()
		}
	})
	return func() {
		a.OnToggle = prev
		unsub()
	}
}

// BindRange two-way-binds a RangeSlider's [Low, High] pair to a [2]float64
// observable (fields RangeSlider.Low / RangeSlider.High, callback
// RangeSlider.OnChange(low, high)). The tuple is a single [2]float64 so a
// two-handle drag round-trips atomically; index 0 is Low, index 1 is High.
func BindRange(rs *toolkit.RangeSlider, obs *mvvm.Observable[[2]float64], invalidate func()) (unbind func()) {
	v := obs.Get()
	rs.Low, rs.High = v[0], v[1]
	prev := rs.OnChange
	rs.OnChange = func(low, high float64) {
		if prev != nil {
			prev(low, high)
		}
		obs.Set([2]float64{low, high})
	}
	unsub := obs.Subscribe(func(v [2]float64) {
		rs.Low, rs.High = v[0], v[1]
		if invalidate != nil {
			invalidate()
		}
	})
	return func() {
		rs.OnChange = prev
		unsub()
	}
}

// BindColorPicker two-way-binds a ColorPicker's colour to an RGBA observable via
// its Color() / SetColor() accessors (it has no exported colour field) and its
// OnChange(c) callback. ColorChooser, which does expose a Color field, uses the
// field-based BindColor instead.
func BindColorPicker(cp *toolkit.ColorPicker, obs *mvvm.Observable[toolkit.RGBA], invalidate func()) (unbind func()) {
	cp.SetColor(obs.Get())
	prev := cp.OnChange
	cp.OnChange = func(c toolkit.RGBA) {
		if prev != nil {
			prev(c)
		}
		obs.Set(c)
	}
	unsub := obs.Subscribe(func(c toolkit.RGBA) {
		cp.SetColor(c)
		if invalidate != nil {
			invalidate()
		}
	})
	return func() {
		cp.OnChange = prev
		unsub()
	}
}

// BindDate two-way-binds a Calendar's selected date to a Date observable via its
// SetDate(y, m, d) seed and OnSelect(y, m, d) callback. OnSelect's three-argument
// signature is not func(Date), so BindField cannot drive it.
func BindDate(c *toolkit.Calendar, obs *mvvm.Observable[Date], invalidate func()) (unbind func()) {
	d := obs.Get()
	c.SetDate(d.Year, d.Month, d.Day)
	prev := c.OnSelect
	c.OnSelect = func(y, m, day int) {
		if prev != nil {
			prev(y, m, day)
		}
		obs.Set(Date{Year: y, Month: m, Day: day})
	}
	unsub := obs.Subscribe(func(d Date) {
		c.SetDate(d.Year, d.Month, d.Day)
		if invalidate != nil {
			invalidate()
		}
	})
	return func() {
		c.OnSelect = prev
		unsub()
	}
}

// BindDatePicker two-way-binds a DatePicker's date to a Date observable via its
// SetDate(y, m, d) seed and OnChange(y, m, d) callback — the DatePicker analogue
// of BindDate (a DatePicker reports its date through OnChange, a Calendar through
// OnSelect).
func BindDatePicker(dp *toolkit.DatePicker, obs *mvvm.Observable[Date], invalidate func()) (unbind func()) {
	d := obs.Get()
	dp.SetDate(d.Year, d.Month, d.Day)
	prev := dp.OnChange
	dp.OnChange = func(y, m, day int) {
		if prev != nil {
			prev(y, m, day)
		}
		obs.Set(Date{Year: y, Month: m, Day: day})
	}
	unsub := obs.Subscribe(func(d Date) {
		dp.SetDate(d.Year, d.Month, d.Day)
		if invalidate != nil {
			invalidate()
		}
	})
	return func() {
		dp.OnChange = prev
		unsub()
	}
}

// BindTime two-way-binds a TimePicker's [Hour, Minute] to a TimeOfDay observable
// (fields TimePicker.Hour / TimePicker.Minute, callback
// TimePicker.OnChange(hour, minute)). The pair round-trips atomically as one
// TimeOfDay.
func BindTime(tp *toolkit.TimePicker, obs *mvvm.Observable[TimeOfDay], invalidate func()) (unbind func()) {
	t := obs.Get()
	tp.Hour, tp.Minute = t.Hour, t.Minute
	prev := tp.OnChange
	tp.OnChange = func(h, m int) {
		if prev != nil {
			prev(h, m)
		}
		obs.Set(TimeOfDay{Hour: h, Minute: m})
	}
	unsub := obs.Subscribe(func(t TimeOfDay) {
		tp.Hour, tp.Minute = t.Hour, t.Minute
		if invalidate != nil {
			invalidate()
		}
	})
	return func() {
		tp.OnChange = prev
		unsub()
	}
}

// BindTextView two-way-binds a TextView's text to a string observable through its
// Text() / SetText() accessors (it has no exported Text field) and its OnChange
// callback (which carries no value, so the view→VM edge reads Text() when it
// fires). This is the accessor-based binding pattern.
func BindTextView(tv *toolkit.TextView, obs *mvvm.Observable[string], invalidate func()) (unbind func()) {
	tv.SetText(obs.Get())
	prev := tv.OnChange
	tv.OnChange = func() {
		if prev != nil {
			prev()
		}
		obs.Set(tv.Text())
	}
	unsub := obs.Subscribe(func(s string) {
		tv.SetText(s)
		if invalidate != nil {
			invalidate()
		}
	})
	return func() {
		tv.OnChange = prev
		unsub()
	}
}

// BindAgendaRename two-way-binds an AgendaSidebar's calendar names to a string
// ObservableList (callback AgendaSidebar.OnRename(i, name)). An inline rename
// pushes the new name into the list at that index; a list change writes the
// names back onto the matching calendars (up to the shorter length, leaving each
// calendar's Color/Hidden untouched). The list is the source of truth: the
// sidebar's names are seeded from it. Returns an unbind that restores the prior
// OnRename and detaches.
func BindAgendaRename(s *toolkit.AgendaSidebar, l *mvvm.ObservableList[string], invalidate func()) (unbind func()) {
	apply := func() {
		src := l.Slice()
		n := len(src)
		if len(s.Calendars) < n {
			n = len(s.Calendars)
		}
		for i := 0; i < n; i++ {
			s.Calendars[i].Name = src[i]
		}
		if invalidate != nil {
			invalidate()
		}
	}
	apply()
	prev := s.OnRename
	s.OnRename = func(i int, name string) {
		if prev != nil {
			prev(i, name)
		}
		if i >= 0 && i < l.Len() {
			l.Set(i, name)
		}
	}
	unsub := l.Subscribe(func(mvvm.ListEvent[string]) { apply() })
	return func() {
		s.OnRename = prev
		unsub()
	}
}

// BindMenuChecks two-way-binds a Menu's item checked states to a bool
// ObservableList (callback Menu.OnItemToggle(i, checked)). Toggling a checkable
// or radio row pushes its new state into the list at that index; a list change
// writes the flags back onto the matching items (up to the shorter length). The
// list is the source of truth: the menu's Checked flags are seeded from it.
// Returns an unbind that restores the prior OnItemToggle and detaches.
func BindMenuChecks(m *toolkit.Menu, l *mvvm.ObservableList[bool], invalidate func()) (unbind func()) {
	apply := func() {
		src := l.Slice()
		n := len(src)
		if len(m.Items) < n {
			n = len(m.Items)
		}
		for i := 0; i < n; i++ {
			m.Items[i].Checked = src[i]
		}
		if invalidate != nil {
			invalidate()
		}
	}
	apply()
	prev := m.OnItemToggle
	m.OnItemToggle = func(i int, checked bool) {
		if prev != nil {
			prev(i, checked)
		}
		if i >= 0 && i < l.Len() {
			l.Set(i, checked)
		}
	}
	unsub := l.Subscribe(func(mvvm.ListEvent[bool]) { apply() })
	return func() {
		m.OnItemToggle = prev
		unsub()
	}
}
