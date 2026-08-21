// Copyright (c) 2026 the go-widgets/mvvmtk authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package mvvmtk

import (
	"github.com/go-widgets/mvvm"
	"github.com/go-widgets/toolkit"
)

// ── Two-way scalar bindings (widget value ⇄ Observable) ─────────────────────
//
// A ViewModel change flows to the widget (and repaints via invalidate) and a
// user edit flows back to the observable. Widgets whose state is MVVM-migrated
// expose it as an Observable accessor, so their binder is bindObs2 (mirror two
// Observables); the remaining field-based widgets still use mvvm.BindField
// (pointer to a value field + a callback slot). invalidate may be nil. The
// returned unbind detaches the subscription(s).

// bindObs2 two-way-binds two Observables — a view-model one and a widget one that
// a migrated widget exposes through an accessor (e.g. Switch.On(), Rating.Value()).
// It seeds the widget from the VM, then mirrors each side's changes onto the
// other; because mvvm.Observable.Set is a no-op on an unchanged value, the mirror
// converges instead of looping. invalidate (may be nil) fires on a VM→widget push
// so the host repaints. The returned unbind detaches both subscriptions. This is
// the accessor-era replacement for mvvm.BindField (which needed a pointer to an
// exported value field + a callback slot, both gone from MVVM-migrated widgets).
func bindObs2[T any](vm, widget *mvvm.Observable[T], invalidate func()) (unbind func()) {
	widget.Set(vm.Get())
	u1 := vm.Subscribe(func(v T) {
		widget.Set(v)
		if invalidate != nil {
			invalidate()
		}
	})
	u2 := widget.Subscribe(func(v T) { vm.Set(v) })
	return func() {
		u1()
		u2()
	}
}

// BindText two-way-binds a SearchEntry's Text to a string observable
// (via the SearchEntry.Text() Observable).
func BindText(e *toolkit.SearchEntry, obs *mvvm.Observable[string], invalidate func()) (unbind func()) {
	return bindObs2(obs, e.Text(), invalidate)
}

// BindEntryText two-way-binds an Entry.s Text to a string observable
// (via the Entry.Text() Observable).
func BindEntryText(e *toolkit.Entry, obs *mvvm.Observable[string], invalidate func()) (unbind func()) {
	return bindObs2(obs, e.Text(), invalidate)
}

// BindChecked two-way-binds a CheckButton's Checked to a bool observable
// (via the CheckButton.Checked() Observable).
func BindChecked(cb *toolkit.CheckButton, obs *mvvm.Observable[bool], invalidate func()) (unbind func()) {
	return bindObs2(obs, cb.Checked(), invalidate)
}

// BindSelectedIndex two-way-binds a DropDown's Selected index to an int
// observable (via the DropDown.Selected() Observable).
func BindSelectedIndex(d *toolkit.DropDown, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, d.Selected(), invalidate)
}

// BindSpin two-way-binds a SpinButton's Value to an int observable
// (via the SpinButton.Value() Observable).
func BindSpin(s *toolkit.SpinButton, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, s.Value(), invalidate)
}

// BindListSelection two-way-binds a ListBox's Selected row to an int observable
// (via the ListBox.Selected() Observable).
func BindListSelection(lb *toolkit.ListBox, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, lb.Selected(), invalidate)
}

// BindViewSwitcher two-way-binds a ViewSwitcher's Current segment to an int
// observable (via the ViewSwitcher.Current() Observable).
func BindViewSwitcher(v *toolkit.ViewSwitcher, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, v.Current(), invalidate)
}

// BindSwitch two-way-binds a Switch's On to a bool observable
// (via the Switch.On() Observable).
func BindSwitch(s *toolkit.Switch, obs *mvvm.Observable[bool], invalidate func()) (unbind func()) {
	return bindObs2(obs, s.On(), invalidate)
}

// BindToggle two-way-binds a ToggleButton's Pressed to a bool observable
// (via the ToggleButton.Pressed() Observable).
func BindToggle(t *toolkit.ToggleButton, obs *mvvm.Observable[bool], invalidate func()) (unbind func()) {
	return bindObs2(obs, t.Pressed(), invalidate)
}

// BindRadio two-way-binds a standalone RadioButton's Checked to a bool
// observable (via the RadioButton.Checked() Observable). For a set of
// mutually-exclusive radios use BindRadioGroup, which binds the group's Active
// index; this helper is for a single, group-less radio behaving like a check.
func BindRadio(r *toolkit.RadioButton, obs *mvvm.Observable[bool], invalidate func()) (unbind func()) {
	return bindObs2(obs, r.Checked(), invalidate)
}

// BindRating two-way-binds a Rating's Value to an int observable
// (via the Rating.Value() Observable).
func BindRating(r *toolkit.Rating, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, r.Value(), invalidate)
}

// BindScale two-way-binds a Scale's Value to a float64 observable
// (via the Scale.Value() Observable).
func BindScale(s *toolkit.Scale, obs *mvvm.Observable[float64], invalidate func()) (unbind func()) {
	return bindObs2(obs, s.Value(), invalidate)
}

// BindExpander two-way-binds an Expander's Expanded to a bool observable
// (via the Expander.Expanded() Observable).
func BindExpander(e *toolkit.Expander, obs *mvvm.Observable[bool], invalidate func()) (unbind func()) {
	return bindObs2(obs, e.Expanded(), invalidate)
}

// BindNotebook two-way-binds a Notebook's Active tab to an int observable
// (via the Notebook.Active() Observable).
func BindNotebook(n *toolkit.Notebook, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, n.Active(), invalidate)
}

// BindComboText two-way-binds a ComboBox's Text to a string observable
// (via the ComboBox.Text() Observable).
func BindComboText(c *toolkit.ComboBox, obs *mvvm.Observable[string], invalidate func()) (unbind func()) {
	return bindObs2(obs, c.Text(), invalidate)
}

// BindPagination two-way-binds a Pagination's Current page to an int observable
// (via the Pagination.Current() Observable).
func BindPagination(pg *toolkit.Pagination, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, pg.Current(), invalidate)
}

// BindPagingToolbar two-way-binds a PagingToolbar's Page to an int observable
// (via the PagingToolbar.Page() Observable).
func BindPagingToolbar(pt *toolkit.PagingToolbar, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, pt.Page(), invalidate)
}

// BindGanttSelection two-way-binds a Gantt's Selected task to an int observable
// (via the Gantt.Selected() Observable).
func BindGanttSelection(g *toolkit.Gantt, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, g.Selected(), invalidate)
}

// BindTableSelection two-way-binds a Table's Selected row to an int observable
// (via the Table.Selected() Observable). While MultiSelect is on, Selected
// doubles as the anchor row, so the same edge tracks a range-anchor move.
func BindTableSelection(t *toolkit.Table, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, t.Selected(), invalidate)
}

// BindCarousel two-way-binds a Carousel's Current slide to an int observable
// (via the Carousel.Current() Observable).
func BindCarousel(c *toolkit.Carousel, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, c.Current(), invalidate)
}

// BindCycle two-way-binds a CycleButton.s Index to an int observable (via the
// CycleButton.Index() Observable). A host that also wants the option string reads
// CycleButton.Value() itself.
func BindCycle(c *toolkit.CycleButton, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, c.Index(), invalidate)
}

// BindRadioGroup two-way-binds a RadioGroup's Active member to an int observable
// (via the RadioGroup.Active() Observable). This binds the whole group
// (which member is checked); BindRadio binds a single group-less radio.
func BindRadioGroup(g *toolkit.RadioGroup, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, g.Active(), invalidate)
}

// BindColor two-way-binds a ColorChooser's Color to an RGBA observable (fields
// ColorChooser.Color / ColorChooser.OnChange). ColorPicker, which exposes no
// Color field, is bound with BindColorPicker instead.
func BindColor(c *toolkit.ColorChooser, obs *mvvm.Observable[toolkit.RGBA], invalidate func()) (unbind func()) {
	return bindObs2(obs, c.Color(), invalidate)
}

// ── One-way sinks (Observable → widget value field) ─────────────────────────
//
// For view-only widgets that carry no user-edit callback. Wraps mvvm.OneWay.

// BindLabel one-way-binds a string observable to a Label's Text() Observable
// (a Label has no edit callback, so this is vm→view only): it seeds the label
// from obs now, then pushes every later obs value onto Label.Text().
func BindLabel(l *toolkit.Label, obs *mvvm.Observable[string], invalidate func()) (unbind func()) {
	l.Text().Set(obs.Get())
	return obs.Subscribe(func(s string) {
		l.Text().Set(s)
		if invalidate != nil {
			invalidate()
		}
	})
}

// BindProgress one-way-binds a float64 observable to a ProgressBar's Fraction()
// Observable (a ProgressBar is display-only, with no edit callback). Callers keep
// the value in the widget's [0,1] domain.
func BindProgress(p *toolkit.ProgressBar, obs *mvvm.Observable[float64], invalidate func()) (unbind func()) {
	p.Fraction().Set(obs.Get()) // seed widget from the source, like mvvm.OneWay
	return obs.Subscribe(func(v float64) {
		p.Fraction().Set(v)
		if invalidate != nil {
			invalidate()
		}
	})
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

// BindAccordion two-way-binds an Accordion.s exclusive-mode Expanded section
// index to an int observable (via the Accordion.Expanded() Observable; -1 when
// all are collapsed). A Multiple-mode accordion does not track a single index.
func BindAccordion(a *toolkit.Accordion, obs *mvvm.Observable[int], invalidate func()) (unbind func()) {
	return bindObs2(obs, a.Expanded(), invalidate)
}

// BindRange two-way-binds a RangeSlider's [Low, High] pair to a [2]float64
// observable (fields RangeSlider.Low / RangeSlider.High, callback
// RangeSlider.OnChange(low, high)). The tuple is a single [2]float64 so a
// two-handle drag round-trips atomically; index 0 is Low, index 1 is High.
func BindRange(rs *toolkit.RangeSlider, obs *mvvm.Observable[[2]float64], invalidate func()) (unbind func()) {
	// syncing suppresses the widget→VM publish while a VM→widget push sets the two
	// handles one at a time (otherwise the intermediate [new, old] pair would
	// re-enter obs and double-fire). set() pushes both handles atomically.
	syncing := false
	set := func(v [2]float64) {
		syncing = true
		rs.Low().Set(v[0])
		rs.High().Set(v[1])
		syncing = false
	}
	set(obs.Get())
	// widget → VM: a handle move publishes the current pair.
	publish := func(float64) {
		if syncing {
			return
		}
		obs.Set([2]float64{rs.Low().Get(), rs.High().Get()})
	}
	uL := rs.Low().Subscribe(publish)
	uH := rs.High().Subscribe(publish)
	// VM → widget: set both handles, then repaint.
	uO := obs.Subscribe(func(v [2]float64) {
		set(v)
		if invalidate != nil {
			invalidate()
		}
	})
	return func() {
		uL()
		uH()
		uO()
	}
}

// BindColorPicker two-way-binds a ColorPicker's colour to an RGBA observable via
// its Color() / SetColor() accessors (it has no exported colour field) and its
// OnChange(c) callback. ColorChooser, which does expose a Color field, uses the
// field-based BindColor instead.
func BindColorPicker(cp *toolkit.ColorPicker, obs *mvvm.Observable[toolkit.RGBA], invalidate func()) (unbind func()) {
	return bindObs2(obs, cp.Color(), invalidate)
}

// BindDate two-way-binds a Calendar's selected date to a Date observable via its
// SetDate(y, m, d) seed and OnSelect(y, m, d) callback. OnSelect's three-argument
// signature is not func(Date), so BindField cannot drive it.
func BindDate(c *toolkit.Calendar, obs *mvvm.Observable[Date], invalidate func()) (unbind func()) {
	return bindCalendarDate(c, obs, invalidate)
}

// bindCalendarDate two-way-binds a Calendar's Year()/Month()/Day() Observables to
// a Date observable. A syncing guard suppresses the widget→VM publish while a
// VM→widget SetDate sets the three components, so the intermediate partial dates
// are not published back (like BindRange's guard for the [2]float64 pair).
func bindCalendarDate(c *toolkit.Calendar, obs *mvvm.Observable[Date], invalidate func()) (unbind func()) {
	syncing := false
	set := func(d Date) {
		syncing = true
		c.SetDate(d.Year, d.Month, d.Day)
		syncing = false
	}
	set(obs.Get())
	push := func(int) {
		if syncing {
			return
		}
		obs.Set(Date{Year: c.Year().Get(), Month: c.Month().Get(), Day: c.Day().Get()})
	}
	uY := c.Year().Subscribe(push)
	uM := c.Month().Subscribe(push)
	uD := c.Day().Subscribe(push)
	uO := obs.Subscribe(func(d Date) {
		set(d)
		if invalidate != nil {
			invalidate()
		}
	})
	return func() {
		uY()
		uM()
		uD()
		uO()
	}
}

// BindDatePicker two-way-binds a DatePicker's date to a Date observable via its
// SetDate(y, m, d) seed and OnChange(y, m, d) callback — the DatePicker analogue
// of BindDate (a DatePicker reports its date through OnChange, a Calendar through
// OnSelect).
func BindDatePicker(dp *toolkit.DatePicker, obs *mvvm.Observable[Date], invalidate func()) (unbind func()) {
	return bindCalendarDate(dp.Cal, obs, invalidate)
}

// BindTime two-way-binds a TimePicker's [Hour, Minute] to a TimeOfDay observable
// (fields TimePicker.Hour / TimePicker.Minute, callback
// TimePicker.OnChange(hour, minute)). The pair round-trips atomically as one
// TimeOfDay.
func BindTime(tp *toolkit.TimePicker, obs *mvvm.Observable[TimeOfDay], invalidate func()) (unbind func()) {
	syncing := false
	set := func(t TimeOfDay) {
		syncing = true
		tp.Hour().Set(t.Hour)
		tp.Minute().Set(t.Minute)
		syncing = false
	}
	set(obs.Get())
	push := func(int) {
		if syncing {
			return
		}
		obs.Set(TimeOfDay{Hour: tp.Hour().Get(), Minute: tp.Minute().Get()})
	}
	uH := tp.Hour().Subscribe(push)
	uM := tp.Minute().Subscribe(push)
	uO := obs.Subscribe(func(t TimeOfDay) {
		set(t)
		if invalidate != nil {
			invalidate()
		}
	})
	return func() {
		uH()
		uM()
		uO()
	}
}

// BindTextView two-way-binds a TextView's text to a string observable through its
// Text() / SetText() accessors (it has no exported Text field) and its OnChange
// callback (which carries no value, so the view→VM edge reads Text() when it
// fires). This is the accessor-based binding pattern.
func BindTextView(tv *toolkit.TextView, obs *mvvm.Observable[string], invalidate func()) (unbind func()) {
	return bindObs2(obs, tv.Text(), invalidate)
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
