// Copyright (c) 2026 the go-widgets/mvvmtk authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package mvvmtk

import (
	"testing"

	"github.com/go-widgets/mvvm"
	"github.com/go-widgets/toolkit"
)

// Wave 7 (bindability): a helper per stateful widget whose state became
// observable this wave. Each test exercises the full two-way contract — seed
// (the observable overrides the widget), VM→view push with an invalidate count,
// view→VM through the widget callback, and unbind severing both edges.

// ── BindField-backed scalar binders ─────────────────────────────────────────

func TestBindSwitch(t *testing.T) {
	obs := mvvm.NewObservable(false)
	s := toolkit.NewSwitch(true)
	c := &counter{}
	unbind := BindSwitch(s, obs, c.inc)
	if s.On().Get() != false {
		t.Fatalf("seed: On=%v, want false", s.On().Get())
	}
	obs.Set(true)
	if !s.On().Get() || c.n != 1 {
		t.Fatalf("vm→view: On=%v n=%d", s.On().Get(), c.n)
	}
	s.On().Set(false)
	if obs.Get() != false {
		t.Fatalf("view→vm: obs=%v", obs.Get())
	}
	unbind()
	obs.Set(true)
	if s.On().Get() {
		t.Fatal("unbind: VM→view must stop")
	}
}

func TestBindToggle(t *testing.T) {
	obs := mvvm.NewObservable(false)
	tb := toolkit.NewToggleButton("b", true)
	c := &counter{}
	unbind := BindToggle(tb, obs, c.inc)
	if tb.Pressed().Get() != false {
		t.Fatalf("seed: Pressed=%v", tb.Pressed().Get())
	}
	obs.Set(true)
	if !tb.Pressed().Get() || c.n != 1 {
		t.Fatalf("vm→view: Pressed=%v n=%d", tb.Pressed().Get(), c.n)
	}
	tb.Pressed().Set(false)
	if obs.Get() != false {
		t.Fatalf("view→vm: obs=%v", obs.Get())
	}
	unbind()
}

func TestBindRadio(t *testing.T) {
	obs := mvvm.NewObservable(false)
	r := toolkit.NewRadioButton("r")
	r.Checked().Set(true)
	c := &counter{}
	unbind := BindRadio(r, obs, c.inc)
	if r.Checked().Get() != false {
		t.Fatalf("seed: Checked=%v", r.Checked().Get())
	}
	obs.Set(true)
	if !r.Checked().Get() || c.n != 1 {
		t.Fatalf("vm→view: Checked=%v n=%d", r.Checked().Get(), c.n)
	}
	r.Checked().Set(false)
	if obs.Get() != false {
		t.Fatalf("view→vm: obs=%v", obs.Get())
	}
	unbind()
}

func TestBindRating(t *testing.T) {
	obs := mvvm.NewObservable(0)
	r := toolkit.NewRating(5, 5)
	c := &counter{}
	unbind := BindRating(r, obs, c.inc)
	if r.Value().Get() != 0 {
		t.Fatalf("seed: Value=%d", r.Value().Get())
	}
	obs.Set(3)
	if r.Value().Get() != 3 || c.n != 1 {
		t.Fatalf("vm→view: Value=%d n=%d", r.Value().Get(), c.n)
	}
	r.Value().Set(4)
	if obs.Get() != 4 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
	unbind()
}

func TestBindScale(t *testing.T) {
	obs := mvvm.NewObservable(0.0)
	s := toolkit.NewScale(0, 10, 7)
	c := &counter{}
	unbind := BindScale(s, obs, c.inc)
	if s.Value().Get() != 0 {
		t.Fatalf("seed: Value=%v", s.Value().Get())
	}
	obs.Set(2.5)
	if s.Value().Get() != 2.5 || c.n != 1 {
		t.Fatalf("vm→view: Value=%v n=%d", s.Value().Get(), c.n)
	}
	s.Value().Set(4.5)
	if obs.Get() != 4.5 {
		t.Fatalf("view→vm: obs=%v", obs.Get())
	}
	unbind()
}

func TestBindExpander(t *testing.T) {
	obs := mvvm.NewObservable(false)
	e := toolkit.NewExpander("x", nil)
	e.Expanded().Set(true)
	c := &counter{}
	unbind := BindExpander(e, obs, c.inc)
	if e.Expanded().Get() != false {
		t.Fatalf("seed: Expanded=%v", e.Expanded().Get())
	}
	obs.Set(true)
	if !e.Expanded().Get() || c.n != 1 {
		t.Fatalf("vm→view: Expanded=%v n=%d", e.Expanded().Get(), c.n)
	}
	e.Expanded().Set(false)
	if obs.Get() != false {
		t.Fatalf("view→vm: obs=%v", obs.Get())
	}
	unbind()
}

func TestBindNotebook(t *testing.T) {
	obs := mvvm.NewObservable(0)
	n := toolkit.NewNotebook()
	n.Active().Set(3)
	c := &counter{}
	unbind := BindNotebook(n, obs, c.inc)
	if n.Active().Get() != 0 {
		t.Fatalf("seed: Active=%d", n.Active().Get())
	}
	obs.Set(2)
	if n.Active().Get() != 2 || c.n != 1 {
		t.Fatalf("vm→view: Active=%d n=%d", n.Active().Get(), c.n)
	}
	n.Active().Set(1)
	if obs.Get() != 1 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
	unbind()
}

func TestBindComboText(t *testing.T) {
	obs := mvvm.NewObservable("a")
	cb := toolkit.NewComboBox([]string{"a", "b", "c"})
	c := &counter{}
	unbind := BindComboText(cb, obs, c.inc)
	if cb.Text().Get() != "a" {
		t.Fatalf("seed: Text=%q", cb.Text().Get())
	}
	obs.Set("b")
	if cb.Text().Get() != "b" || c.n != 1 {
		t.Fatalf("vm→view: Text=%q n=%d", cb.Text().Get(), c.n)
	}
	cb.Text().Set("c")
	if obs.Get() != "c" {
		t.Fatalf("view→vm: obs=%q", obs.Get())
	}
	unbind()
}

func TestBindPagination(t *testing.T) {
	obs := mvvm.NewObservable(1)
	pg := toolkit.NewPagination(5, 10)
	c := &counter{}
	unbind := BindPagination(pg, obs, c.inc)
	if pg.Current().Get() != 1 {
		t.Fatalf("seed: Current=%d", pg.Current().Get())
	}
	obs.Set(4)
	if pg.Current().Get() != 4 || c.n != 1 {
		t.Fatalf("vm→view: Current=%d n=%d", pg.Current().Get(), c.n)
	}
	pg.Current().Set(7)
	if obs.Get() != 7 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
	unbind()
}

func TestBindPagingToolbar(t *testing.T) {
	obs := mvvm.NewObservable(1)
	pt := toolkit.NewPagingToolbar(5, 10)
	c := &counter{}
	unbind := BindPagingToolbar(pt, obs, c.inc)
	if pt.Page().Get() != 1 {
		t.Fatalf("seed: Page=%d", pt.Page().Get())
	}
	obs.Set(4)
	if pt.Page().Get() != 4 || c.n != 1 {
		t.Fatalf("vm→view: Page=%d n=%d", pt.Page().Get(), c.n)
	}
	pt.Page().Set(6)
	if obs.Get() != 6 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
	unbind()
}

func TestBindGanttSelection(t *testing.T) {
	obs := mvvm.NewObservable(-1)
	g := toolkit.NewGantt(nil)
	g.Selected().Set(5)
	c := &counter{}
	unbind := BindGanttSelection(g, obs, c.inc)
	if g.Selected().Get() != -1 {
		t.Fatalf("seed: Selected=%d", g.Selected().Get())
	}
	obs.Set(2)
	if g.Selected().Get() != 2 || c.n != 1 {
		t.Fatalf("vm→view: Selected=%d n=%d", g.Selected().Get(), c.n)
	}
	g.Selected().Set(3)
	if obs.Get() != 3 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
	unbind()
}

func TestBindTableSelection(t *testing.T) {
	obs := mvvm.NewObservable(-1)
	tb := toolkit.NewTable([]toolkit.TableColumn{{Title: "A"}}, [][]string{{"r0"}, {"r1"}})
	tb.Selected().Set(1)
	c := &counter{}
	unbind := BindTableSelection(tb, obs, c.inc)
	if tb.Selected().Get() != -1 {
		t.Fatalf("seed: Selected=%d", tb.Selected().Get())
	}
	obs.Set(1)
	if tb.Selected().Get() != 1 || c.n != 1 {
		t.Fatalf("vm→view: Selected=%d n=%d", tb.Selected().Get(), c.n)
	}
	tb.Selected().Set(0)
	if obs.Get() != 0 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
	unbind()
	obs.Set(1)
	if tb.Selected().Get() != 0 {
		t.Fatal("unbind: VM→view must stop")
	}
}

func TestBindCarousel(t *testing.T) {
	obs := mvvm.NewObservable(0)
	car := toolkit.NewCarousel([]toolkit.Widget{toolkit.NewLabel("a"), toolkit.NewLabel("b"), toolkit.NewLabel("c")})
	car.Current().Set(2)
	c := &counter{}
	unbind := BindCarousel(car, obs, c.inc)
	if car.Current().Get() != 0 {
		t.Fatalf("seed: Current=%d", car.Current().Get())
	}
	obs.Set(2)
	if car.Current().Get() != 2 || c.n != 1 {
		t.Fatalf("vm→view: Current=%d n=%d", car.Current().Get(), c.n)
	}
	car.Current().Set(1)
	if obs.Get() != 1 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
	unbind()
}

func TestBindCycle(t *testing.T) {
	obs := mvvm.NewObservable(0)
	cy := toolkit.NewCycleButton("List", "Grid", "Compact")
	cy.Index().Set(2)
	c := &counter{}
	unbind := BindCycle(cy, obs, c.inc)
	if cy.Index().Get() != 0 {
		t.Fatalf("seed: Index=%d", cy.Index().Get())
	}
	obs.Set(2)
	if cy.Index().Get() != 2 || c.n != 1 {
		t.Fatalf("vm→view: Index=%d n=%d", cy.Index().Get(), c.n)
	}
	cy.Index().Set(1)
	if obs.Get() != 1 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
	unbind()
}

func TestBindRadioGroup(t *testing.T) {
	obs := mvvm.NewObservable(-1)
	g := toolkit.NewRadioGroup()
	g.Add(toolkit.NewRadioButton("a"))
	g.Add(toolkit.NewRadioButton("b"))
	g.Active().Set(1)
	c := &counter{}
	unbind := BindRadioGroup(g, obs, c.inc)
	if g.Active().Get() != -1 {
		t.Fatalf("seed: Active=%d", g.Active().Get())
	}
	obs.Set(0)
	if g.Active().Get() != 0 || c.n != 1 {
		t.Fatalf("vm→view: Active=%d n=%d", g.Active().Get(), c.n)
	}
	g.Active().Set(1)
	if obs.Get() != 1 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
	unbind()
}

func TestBindColor(t *testing.T) {
	red := toolkit.RGBA{R: 255, A: 255}
	blue := toolkit.RGBA{B: 255, A: 255}
	obs := mvvm.NewObservable(red)
	cc := toolkit.NewColorChooser(blue)
	c := &counter{}
	unbind := BindColor(cc, obs, c.inc)
	if cc.Color().Get() != red {
		t.Fatalf("seed: Color=%+v, want red", cc.Color().Get())
	}
	obs.Set(blue)
	if cc.Color().Get() != blue || c.n != 1 {
		t.Fatalf("vm→view: Color=%+v n=%d", cc.Color().Get(), c.n)
	}
	cc.Color().Set(red)
	if obs.Get() != red {
		t.Fatalf("view→vm: obs=%+v", obs.Get())
	}
	unbind()
}

// ── Custom / accessor / multi-value binders ─────────────────────────────────
// These cover both the prior-callback-composed path (prev != nil, invalidate
// != nil) and the bare path (prev == nil, invalidate == nil).

func TestBindAccordion(t *testing.T) {
	obs := mvvm.NewObservable(-1)
	a := toolkit.NewAccordion([]toolkit.AccordionSection{{Title: "One"}, {Title: "Two"}})
	a.Expanded().Set(1)
	c := &counter{}
	unbind := BindAccordion(a, obs, c.inc)
	if a.Expanded().Get() != -1 {
		t.Fatalf("seed should override widget: Expanded=%d", a.Expanded().Get())
	}
	obs.Set(0)
	if a.Expanded().Get() != 0 || c.n != 1 {
		t.Fatalf("vm→view: Expanded=%d n=%d", a.Expanded().Get(), c.n)
	}
	// A user toggle flips the widget's Expanded Observable; the binder publishes it.
	a.Expanded().Set(1)
	if obs.Get() != 1 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
	unbind()
	obs.Set(2)
	if a.Expanded().Get() != 1 {
		t.Fatal("unbind: VM→view must stop")
	}
	a.Expanded().Set(0)
	if obs.Get() != 2 {
		t.Fatalf("unbind: view→vm must stop, obs=%d", obs.Get())
	}

	// Bare path: nil invalidate — must not panic.
	obs2 := mvvm.NewObservable(-1)
	a2 := toolkit.NewAccordion([]toolkit.AccordionSection{{Title: "X"}})
	u2 := BindAccordion(a2, obs2, nil)
	obs2.Set(0) // invalidate == nil branch
	if a2.Expanded().Get() != 0 {
		t.Fatalf("bare vm→view: Expanded=%d", a2.Expanded().Get())
	}
	a2.Expanded().Set(0) // no-op change, still fine
	if obs2.Get() != 0 {
		t.Fatalf("bare: obs=%d", obs2.Get())
	}
	u2()
}

func TestBindRange(t *testing.T) {
	obs := mvvm.NewObservableEq([2]float64{0, 1}, func(a, b [2]float64) bool { return a == b })
	rs := toolkit.NewRangeSlider(0, 100, 10, 90)
	c := &counter{}
	unbind := BindRange(rs, obs, c.inc)
	if rs.Low().Get() != 0 || rs.High().Get() != 1 {
		t.Fatalf("seed: Low=%v High=%v", rs.Low().Get(), rs.High().Get())
	}
	obs.Set([2]float64{20, 80})
	if rs.Low().Get() != 20 || rs.High().Get() != 80 || c.n != 1 {
		t.Fatalf("vm→view: Low=%v High=%v n=%d", rs.Low().Get(), rs.High().Get(), c.n)
	}
	// A handle move (either Observable) publishes the current pair to the VM.
	rs.Low().Set(30)
	rs.High().Set(70)
	if obs.Get() != ([2]float64{30, 70}) {
		t.Fatalf("view→vm: obs=%v", obs.Get())
	}
	unbind()
	obs.Set([2]float64{1, 2})
	if rs.Low().Get() != 30 || rs.High().Get() != 70 {
		t.Fatal("unbind: VM→view must stop")
	}
	rs.Low().Set(9)
	if obs.Get() != ([2]float64{1, 2}) {
		t.Fatal("unbind: view→vm must stop")
	}

	// Bare path.
	obs2 := mvvm.NewObservableEq([2]float64{0, 0}, func(a, b [2]float64) bool { return a == b })
	rs2 := toolkit.NewRangeSlider(0, 100, 0, 0)
	u2 := BindRange(rs2, obs2, nil)
	obs2.Set([2]float64{5, 6})
	if rs2.Low().Get() != 5 || rs2.High().Get() != 6 {
		t.Fatalf("bare vm→view: Low=%v High=%v", rs2.Low().Get(), rs2.High().Get())
	}
	rs2.Low().Set(7) // publishes [7, High=6]
	if obs2.Get() != ([2]float64{7, 6}) {
		t.Fatalf("bare: obs=%v", obs2.Get())
	}
	u2()
}

func TestBindColorPicker(t *testing.T) {
	red := toolkit.RGBA{R: 255, A: 255}
	blue := toolkit.RGBA{B: 255, A: 255}
	green := toolkit.RGBA{G: 255, A: 255}
	obs := mvvm.NewObservable(red)
	cp := toolkit.NewColorPicker(blue)
	prior := 0
	cp.OnChange = func(toolkit.RGBA) { prior++ }
	c := &counter{}
	unbind := BindColorPicker(cp, obs, c.inc)
	if cp.Color() != red {
		t.Fatalf("seed: Color=%+v", cp.Color())
	}
	obs.Set(green)
	if cp.Color() != green || c.n != 1 {
		t.Fatalf("vm→view: Color=%+v n=%d", cp.Color(), c.n)
	}
	cp.OnChange(blue)
	if obs.Get() != blue || prior != 1 {
		t.Fatalf("view→vm: obs=%+v prior=%d", obs.Get(), prior)
	}
	unbind()
	obs.Set(red)
	if cp.Color() != blue {
		t.Fatal("unbind: VM→view must stop")
	}

	// Bare path.
	obs2 := mvvm.NewObservable(red)
	cp2 := toolkit.NewColorPicker(blue)
	u2 := BindColorPicker(cp2, obs2, nil)
	obs2.Set(green)
	cp2.OnChange(red)
	if obs2.Get() != red {
		t.Fatalf("bare: obs=%+v", obs2.Get())
	}
	u2()
}

func TestBindDate(t *testing.T) {
	obs := mvvm.NewObservable(Date{2021, 6, 15})
	cal := toolkit.NewCalendar(2000, 1, 1)
	c := &counter{}
	unbind := BindDate(cal, obs, c.inc)
	if cal.Year().Get() != 2021 || cal.Month().Get() != 6 || cal.Day().Get() != 15 {
		t.Fatalf("seed: %d-%d-%d", cal.Year().Get(), cal.Month().Get(), cal.Day().Get())
	}
	obs.Set(Date{2022, 3, 4})
	if cal.Year().Get() != 2022 || cal.Month().Get() != 3 || cal.Day().Get() != 4 || c.n != 1 {
		t.Fatalf("vm→view: %d-%d-%d n=%d", cal.Year().Get(), cal.Month().Get(), cal.Day().Get(), c.n)
	}
	// A user date pick sets the Calendar's date Observables; the binder publishes it.
	cal.SetDate(2023, 12, 25)
	if obs.Get() != (Date{2023, 12, 25}) {
		t.Fatalf("view→vm: obs=%+v", obs.Get())
	}
	unbind()
	obs.Set(Date{2000, 1, 1})
	if cal.Year().Get() != 2023 {
		t.Fatal("unbind: VM→view must stop")
	}

	// Bare path.
	obs2 := mvvm.NewObservable(Date{2010, 5, 5})
	cal2 := toolkit.NewCalendar(2000, 1, 1)
	u2 := BindDate(cal2, obs2, nil)
	obs2.Set(Date{2011, 6, 6})
	if cal2.Year().Get() != 2011 {
		t.Fatalf("bare vm→view: %d", cal2.Year().Get())
	}
	u2()
}

func TestBindDatePicker(t *testing.T) {
	obs := mvvm.NewObservable(Date{2021, 6, 15})
	dp := toolkit.NewDatePicker(2000, 1, 1)
	c := &counter{}
	unbind := BindDatePicker(dp, obs, c.inc)
	if dp.Cal.Year().Get() != 2021 || dp.Cal.Month().Get() != 6 || dp.Cal.Day().Get() != 15 {
		t.Fatalf("seed: %d-%d-%d", dp.Cal.Year().Get(), dp.Cal.Month().Get(), dp.Cal.Day().Get())
	}
	obs.Set(Date{2022, 3, 4})
	if dp.Cal.Year().Get() != 2022 || dp.Cal.Month().Get() != 3 || dp.Cal.Day().Get() != 4 || c.n != 1 {
		t.Fatalf("vm→view: %d-%d-%d n=%d", dp.Cal.Year().Get(), dp.Cal.Month().Get(), dp.Cal.Day().Get(), c.n)
	}
	dp.Cal.SetDate(2023, 12, 25) // user picks a day in the popup
	if obs.Get() != (Date{2023, 12, 25}) {
		t.Fatalf("view→vm: obs=%+v", obs.Get())
	}
	unbind()

	// Bare path.
	obs2 := mvvm.NewObservable(Date{2010, 5, 5})
	dp2 := toolkit.NewDatePicker(2000, 1, 1)
	u2 := BindDatePicker(dp2, obs2, nil)
	obs2.Set(Date{2011, 6, 6})
	if dp2.Cal.Year().Get() != 2011 {
		t.Fatalf("bare: %d", dp2.Cal.Year().Get())
	}
	u2()
}

func TestBindTime(t *testing.T) {
	obs := mvvm.NewObservable(TimeOfDay{9, 30})
	tp := toolkit.NewTimePicker(0, 0)
	c := &counter{}
	unbind := BindTime(tp, obs, c.inc)
	if tp.Hour().Get() != 9 || tp.Minute().Get() != 30 {
		t.Fatalf("seed: %d:%d", tp.Hour().Get(), tp.Minute().Get())
	}
	obs.Set(TimeOfDay{14, 45})
	if tp.Hour().Get() != 14 || tp.Minute().Get() != 45 || c.n != 1 {
		t.Fatalf("vm→view: %d:%d n=%d", tp.Hour().Get(), tp.Minute().Get(), c.n)
	}
	// A user spin sets the Hour/Minute Observables; the binder publishes the pair.
	tp.Hour().Set(6)
	tp.Minute().Set(15)
	if obs.Get() != (TimeOfDay{6, 15}) {
		t.Fatalf("view→vm: obs=%+v", obs.Get())
	}
	unbind()

	// Bare path.
	obs2 := mvvm.NewObservable(TimeOfDay{1, 1})
	tp2 := toolkit.NewTimePicker(0, 0)
	u2 := BindTime(tp2, obs2, nil)
	obs2.Set(TimeOfDay{2, 2})
	if tp2.Hour().Get() != 2 || tp2.Minute().Get() != 2 {
		t.Fatalf("bare: %d:%d", tp2.Hour().Get(), tp2.Minute().Get())
	}
	u2()
}

func TestBindTextView(t *testing.T) {
	obs := mvvm.NewObservable("hello")
	tv := toolkit.NewTextView("stale")
	c := &counter{}
	unbind := BindTextView(tv, obs, c.inc)
	if tv.Text().Get() != "hello" {
		t.Fatalf("seed: Text=%q", tv.Text().Get())
	}
	obs.Set("world")
	if tv.Text().Get() != "world" || c.n != 1 {
		t.Fatalf("vm→view: Text=%q n=%d", tv.Text().Get(), c.n)
	}
	// A view edit publishes through Text(); the binder reads it back into obs.
	tv.SetText("typed")
	if obs.Get() != "typed" {
		t.Fatalf("view→vm: obs=%q", obs.Get())
	}
	unbind()
	obs.Set("after")
	if tv.Text().Get() != "typed" {
		t.Fatal("unbind: VM→view must stop")
	}

	// Bare path.
	obs2 := mvvm.NewObservable("a")
	tv2 := toolkit.NewTextView("")
	u2 := BindTextView(tv2, obs2, nil)
	obs2.Set("b")
	tv2.SetText("c")
	if obs2.Get() != "c" {
		t.Fatalf("bare: obs=%q", obs2.Get())
	}
	u2()
}

func TestBindAgendaRename(t *testing.T) {
	list := mvvm.NewObservableList("Work", "Home")
	sb := toolkit.NewAgendaSidebar([]toolkit.AgendaCalendar{
		{Name: "old0"}, {Name: "old1"},
	})
	prior := 0
	sb.OnRename = func(int, string) { prior++ }
	c := &counter{}
	unbind := BindAgendaRename(sb, list, c.inc)
	// Seed: names come from the list (source of truth).
	if sb.Calendars[0].Name != "Work" || sb.Calendars[1].Name != "Home" {
		t.Fatalf("seed: %q,%q", sb.Calendars[0].Name, sb.Calendars[1].Name)
	}
	if c.n != 1 {
		t.Fatalf("seed apply should invalidate once: n=%d", c.n)
	}
	// VM→view: a list change rewrites the matching calendar name.
	list.Set(1, "Family")
	if sb.Calendars[1].Name != "Family" {
		t.Fatalf("vm→view: name1=%q", sb.Calendars[1].Name)
	}
	// view→VM: an inline rename pushes into the list, prior handler runs.
	sb.OnRename(0, "Personal")
	if list.At(0) != "Personal" || prior != 1 {
		t.Fatalf("view→vm: list0=%q prior=%d", list.At(0), prior)
	}
	unbind()
	list.Set(0, "Ignored")
	if sb.Calendars[0].Name == "Ignored" {
		t.Fatal("unbind: VM→view must stop")
	}

	// Bare path + length-mismatch guard: fewer calendars than list entries, nil
	// invalidate, out-of-range rename index is a no-op.
	longList := mvvm.NewObservableList("a", "b", "c")
	sb2 := toolkit.NewAgendaSidebar([]toolkit.AgendaCalendar{{Name: "z"}})
	u2 := BindAgendaRename(sb2, longList, nil)
	if sb2.Calendars[0].Name != "a" {
		t.Fatalf("bare seed: %q", sb2.Calendars[0].Name)
	}
	sb2.OnRename(9, "oob") // index past the list -> ignored, no panic
	sb2.OnRename(0, "A")
	if longList.At(0) != "A" {
		t.Fatalf("bare view→vm: %q", longList.At(0))
	}
	u2()
}

func TestBindMenuChecks(t *testing.T) {
	list := mvvm.NewObservableList(true, false)
	noop := func() {}
	m := &toolkit.Menu{Items: []toolkit.MenuItem{
		{Label: "A", Checkable: true, Action: noop},
		{Label: "B", Checkable: true, Action: noop},
	}}
	prior := 0
	m.OnItemToggle = func(int, bool) { prior++ }
	c := &counter{}
	unbind := BindMenuChecks(m, list, c.inc)
	if !m.Items[0].Checked || m.Items[1].Checked {
		t.Fatalf("seed: %v,%v", m.Items[0].Checked, m.Items[1].Checked)
	}
	if c.n != 1 {
		t.Fatalf("seed apply should invalidate once: n=%d", c.n)
	}
	list.Set(1, true)
	if !m.Items[1].Checked {
		t.Fatalf("vm→view: item1=%v", m.Items[1].Checked)
	}
	m.OnItemToggle(0, false)
	if list.At(0) != false || prior != 1 {
		t.Fatalf("view→vm: list0=%v prior=%d", list.At(0), prior)
	}
	unbind()
	list.Set(0, true)
	if m.Items[0].Checked {
		t.Fatal("unbind: VM→view must stop")
	}

	// Bare path + length-mismatch + out-of-range index.
	longList := mvvm.NewObservableList(true, true, true)
	m2 := &toolkit.Menu{Items: []toolkit.MenuItem{{Label: "X", Checkable: true, Action: noop}}}
	u2 := BindMenuChecks(m2, longList, nil)
	if !m2.Items[0].Checked {
		t.Fatalf("bare seed: %v", m2.Items[0].Checked)
	}
	m2.OnItemToggle(9, false) // out of range -> ignored
	m2.OnItemToggle(0, false)
	if longList.At(0) != false {
		t.Fatalf("bare view→vm: %v", longList.At(0))
	}
	u2()
}
