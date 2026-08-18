// Copyright (c) 2026 the go-widgets/mvvmtk authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package mvvmtk

import (
	"testing"

	"github.com/go-widgets/mvvm"
	"github.com/go-widgets/toolkit"
)

// counter is a fake invalidate hook that records how many repaints a binding
// requested.
type counter struct{ n int }

func (c *counter) inc() { c.n++ }

// TestBindText covers the SearchEntry two-way binding: seed, VM→view push with
// invalidate, view→VM edit through OnChange, and unbind severing both edges.
func TestBindText(t *testing.T) {
	obs := mvvm.NewObservable("hi")
	e := toolkit.NewSearchEntry("stale")
	prior := 0
	e.OnChange = func(string) { prior++ } // pre-existing handler, must be composed
	c := &counter{}
	unbind := BindText(e, obs, c.inc)

	if e.Text != "hi" {
		t.Fatalf("seed: Text=%q, want %q", e.Text, "hi")
	}
	obs.Set("world") // VM→view
	if e.Text != "world" || c.n != 1 {
		t.Fatalf("vm→view: Text=%q n=%d", e.Text, c.n)
	}
	e.OnChange("typed") // view→VM, prior handler still runs
	if obs.Get() != "typed" || prior != 1 {
		t.Fatalf("view→vm: obs=%q prior=%d", obs.Get(), prior)
	}
	unbind()
	obs.Set("after")
	if e.Text != "typed" {
		t.Fatalf("after unbind VM→view must stop: Text=%q", e.Text)
	}
	e.OnChange("x") // only the restored prior handler now — must not touch obs
	if obs.Get() != "after" || prior != 2 {
		t.Fatalf("after unbind view→vm must stop: obs=%q prior=%d", obs.Get(), prior)
	}
}

// TestBindEntryText covers the Entry two-way binding both directions.
func TestBindEntryText(t *testing.T) {
	obs := mvvm.NewObservable("a")
	e := toolkit.NewEntry("")
	c := &counter{}
	unbind := BindEntryText(e, obs, c.inc)
	defer unbind()

	if e.Text != "a" {
		t.Fatalf("seed: Text=%q", e.Text)
	}
	obs.Set("b")
	if e.Text != "b" || c.n != 1 {
		t.Fatalf("vm→view: Text=%q n=%d", e.Text, c.n)
	}
	e.OnChange("c")
	if obs.Get() != "c" {
		t.Fatalf("view→vm: obs=%q", obs.Get())
	}
}

// TestBindChecked covers the CheckButton two-way binding both directions.
func TestBindChecked(t *testing.T) {
	obs := mvvm.NewObservable(false)
	cb := toolkit.NewCheckButton("on", true)
	c := &counter{}
	unbind := BindChecked(cb, obs, c.inc)
	defer unbind()

	if cb.Checked().Get() != false {
		t.Fatalf("seed should override widget: Checked=%v", cb.Checked().Get())
	}
	obs.Set(true)
	if !cb.Checked().Get() || c.n != 1 {
		t.Fatalf("vm→view: Checked=%v n=%d", cb.Checked().Get(), c.n)
	}
	cb.Checked().Set(false)
	if obs.Get() != false {
		t.Fatalf("view→vm: obs=%v", obs.Get())
	}
}

// TestBindSelectedIndex covers the DropDown two-way index binding.
func TestBindSelectedIndex(t *testing.T) {
	obs := mvvm.NewObservable(1)
	d := toolkit.NewDropDown([]string{"a", "b", "c"}, 0)
	c := &counter{}
	unbind := BindSelectedIndex(d, obs, c.inc)
	defer unbind()

	if d.Selected().Get() != 1 {
		t.Fatalf("seed: Selected=%d", d.Selected().Get())
	}
	obs.Set(2)
	if d.Selected().Get() != 2 || c.n != 1 {
		t.Fatalf("vm→view: Selected=%d n=%d", d.Selected().Get(), c.n)
	}
	d.Selected().Set(0)
	if obs.Get() != 0 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
}

// TestBindSpin covers the SpinButton two-way value binding.
func TestBindSpin(t *testing.T) {
	obs := mvvm.NewObservable(5)
	s := toolkit.NewSpinButton(0, 10, 0, 1)
	c := &counter{}
	unbind := BindSpin(s, obs, c.inc)
	defer unbind()

	if s.Value().Get() != 5 {
		t.Fatalf("seed: Value=%d", s.Value().Get())
	}
	obs.Set(7)
	if s.Value().Get() != 7 || c.n != 1 {
		t.Fatalf("vm→view: Value=%d n=%d", s.Value().Get(), c.n)
	}
	s.Value().Set(3)
	if obs.Get() != 3 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
}

// TestBindListSelection covers the ListBox Selected/OnActivate two-way binding.
func TestBindListSelection(t *testing.T) {
	obs := mvvm.NewObservable(0)
	lb := toolkit.NewListBox([]string{"x", "y", "z"})
	c := &counter{}
	unbind := BindListSelection(lb, obs, c.inc)
	defer unbind()

	if lb.Selected != 0 {
		t.Fatalf("seed: Selected=%d", lb.Selected)
	}
	obs.Set(2)
	if lb.Selected != 2 || c.n != 1 {
		t.Fatalf("vm→view: Selected=%d n=%d", lb.Selected, c.n)
	}
	lb.OnActivate(1)
	if obs.Get() != 1 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
}

// TestBindViewSwitcher covers the ViewSwitcher Current/OnChange two-way binding.
func TestBindViewSwitcher(t *testing.T) {
	obs := mvvm.NewObservable(1)
	v := toolkit.NewViewSwitcher([]string{"one", "two"}, 0)
	c := &counter{}
	unbind := BindViewSwitcher(v, obs, c.inc)
	defer unbind()

	if v.Current().Get() != 1 {
		t.Fatalf("seed: Current=%d", v.Current().Get())
	}
	obs.Set(0)
	if v.Current().Get() != 0 || c.n != 1 {
		t.Fatalf("vm→view: Current=%d n=%d", v.Current().Get(), c.n)
	}
	v.Current().Set(1)
	if obs.Get() != 1 {
		t.Fatalf("view→vm: obs=%d", obs.Get())
	}
}

// TestBindLabel covers the one-way Label text sink, including unbind.
func TestBindLabel(t *testing.T) {
	obs := mvvm.NewObservable("start")
	l := toolkit.NewLabel("old")
	c := &counter{}
	unbind := BindLabel(l, obs, c.inc)

	if l.Text != "start" {
		t.Fatalf("seed: Text=%q", l.Text)
	}
	obs.Set("next")
	if l.Text != "next" || c.n != 1 {
		t.Fatalf("push: Text=%q n=%d", l.Text, c.n)
	}
	unbind()
	obs.Set("ignored")
	if l.Text != "next" {
		t.Fatalf("after unbind Text changed: %q", l.Text)
	}
}

// TestBindProgress covers the one-way ProgressBar fraction sink.
func TestBindProgress(t *testing.T) {
	obs := mvvm.NewObservable(0.25)
	p := toolkit.NewProgressBar()
	c := &counter{}
	unbind := BindProgress(p, obs, c.inc)
	defer unbind()

	if p.Fraction != 0.25 {
		t.Fatalf("seed: Fraction=%v", p.Fraction)
	}
	obs.Set(0.75)
	if p.Fraction != 0.75 || c.n != 1 {
		t.Fatalf("push: Fraction=%v n=%d", p.Fraction, c.n)
	}
}

// TestBindListItems covers projecting an ObservableList into ListBox.Items,
// rebuild-on-change, and unbind.
func TestBindListItems(t *testing.T) {
	l := mvvm.NewObservableList(1, 2)
	lb := toolkit.NewListBox(nil)
	c := &counter{}
	project := func(n int) string {
		if n == 1 {
			return "one"
		}
		return "many"
	}
	unbind := BindListItems(lb, l, project, c.inc)

	if len(lb.Items) != 2 || lb.Items[0] != "one" || c.n != 1 {
		t.Fatalf("seed: items=%v n=%d", lb.Items, c.n)
	}
	l.Append(3)
	if len(lb.Items) != 3 || lb.Items[2] != "many" || c.n != 2 {
		t.Fatalf("append: items=%v n=%d", lb.Items, c.n)
	}
	unbind()
	l.Append(4)
	if len(lb.Items) != 3 {
		t.Fatalf("after unbind items rebuilt: %v", lb.Items)
	}
}

// TestBindDropDownOptions covers projecting a list into DropDown.Options.
func TestBindDropDownOptions(t *testing.T) {
	l := mvvm.NewObservableList("a")
	d := toolkit.NewDropDown(nil, 0)
	c := &counter{}
	unbind := BindDropDownOptions(d, l, func(s string) string { return s }, c.inc)
	defer unbind()

	if len(d.Options) != 1 || d.Options[0] != "a" {
		t.Fatalf("seed: options=%v", d.Options)
	}
	l.Append("b")
	if len(d.Options) != 2 || d.Options[1] != "b" || c.n != 2 {
		t.Fatalf("append: options=%v n=%d", d.Options, c.n)
	}
}

// TestBindViews covers projecting a list into ViewSwitcher.Views.
func TestBindViews(t *testing.T) {
	l := mvvm.NewObservableList("first")
	v := toolkit.NewViewSwitcher(nil, 0)
	c := &counter{}
	unbind := BindViews(v, l, func(s string) string { return s }, c.inc)
	defer unbind()

	if len(v.Views) != 1 || v.Views[0] != "first" {
		t.Fatalf("seed: views=%v", v.Views)
	}
	l.Append("second")
	if len(v.Views) != 2 || v.Views[1] != "second" || c.n != 2 {
		t.Fatalf("append: views=%v n=%d", v.Views, c.n)
	}
}

// TestBindCommand covers wiring Execute into OnClick and the greying of the
// button through both the enabled and disabled Style branches, plus unbind.
func TestBindCommand(t *testing.T) {
	allowed := mvvm.NewObservable(true)
	runs := 0
	cmd := mvvm.NewCommand(func() { runs++ }, allowed.Get)
	mvvm.BindCanExecute(cmd, allowed) // re-raise CanExecuteChanged when allowed flips

	b := toolkit.NewButton("Go", nil)
	priorClicks := 0
	b.OnClick = func() { priorClicks++ } // pre-existing handler, must be composed + restored
	b.Style = toolkit.ButtonProminent    // a non-secondary original to observe restore
	c := &counter{}
	unbind := BindCommand(b, cmd, c.inc)

	// Enabled at bind time: original style kept, one repaint requested.
	if b.Style != toolkit.ButtonProminent {
		t.Fatalf("enabled seed Style=%v, want Prominent", b.Style)
	}
	if c.n != 1 {
		t.Fatalf("seed should invalidate once, got %d", c.n)
	}
	b.OnClick() // Execute, prior handler also runs
	if runs != 1 || priorClicks != 1 {
		t.Fatalf("OnClick: runs=%d prior=%d", runs, priorClicks)
	}

	// Disable → greyed to Secondary, repaint.
	allowed.Set(false)
	if b.Style != toolkit.ButtonSecondary || c.n != 2 {
		t.Fatalf("disabled: Style=%v n=%d", b.Style, c.n)
	}
	b.OnClick() // guarded Execute is a no-op when CanExecute is false
	if runs != 1 {
		t.Fatalf("disabled Execute ran: runs=%d", runs)
	}

	// Re-enable → restored to the original Prominent style.
	allowed.Set(true)
	if b.Style != toolkit.ButtonProminent || c.n != 3 {
		t.Fatalf("re-enabled: Style=%v n=%d", b.Style, c.n)
	}

	unbind()
	allowed.Set(false) // no more greying after unbind
	if b.Style != toolkit.ButtonProminent {
		t.Fatalf("after unbind Style changed: %v", b.Style)
	}
	b.OnClick() // OnClick restored to the prior handler — must not Execute
	if runs != 1 || priorClicks != 3 {
		t.Fatalf("after unbind OnClick: runs=%d prior=%d", runs, priorClicks)
	}
}

// TestBindCommandNilInvalidate covers BindCommand's nil-invalidate branch.
func TestBindCommandNilInvalidate(t *testing.T) {
	allowed := mvvm.NewObservable(false)
	cmd := mvvm.NewCommand(func() {}, allowed.Get)
	mvvm.BindCanExecute(cmd, allowed)
	b := toolkit.NewButton("x", nil)
	unbind := BindCommand(b, cmd, nil) // nil invalidate must not panic
	defer unbind()
	if b.Style != toolkit.ButtonSecondary {
		t.Fatalf("disabled seed Style=%v", b.Style)
	}
	allowed.Set(true) // enable branch, still nil invalidate
	if b.Style != toolkit.ButtonDefault {
		t.Fatalf("re-enabled Style=%v, want Default", b.Style)
	}
}

// TestBindTree covers rebuilding a TreeTable's Root forest from a list, on
// change and on unbind, plus the nil-invalidate branch.
func TestBindTree(t *testing.T) {
	l := mvvm.NewObservableList("root-a", "root-b")
	cols := []toolkit.TreeTableColumn{{Title: "Name"}}
	tt := toolkit.NewTreeTable(cols, nil)
	c := &counter{}
	project := func(s string) *toolkit.TreeTableNode {
		return &toolkit.TreeTableNode{Cells: []string{s}}
	}
	unbind := BindTree(tt, l, project, c.inc)

	if len(tt.Root) != 2 || tt.Root[0].Cells[0] != "root-a" || c.n != 1 {
		t.Fatalf("seed: root=%d n=%d", len(tt.Root), c.n)
	}
	l.Append("root-c")
	if len(tt.Root) != 3 || tt.Root[2].Cells[0] != "root-c" || c.n != 2 {
		t.Fatalf("append: root=%d n=%d", len(tt.Root), c.n)
	}
	unbind()
	l.Append("root-d")
	if len(tt.Root) != 3 {
		t.Fatalf("after unbind Root rebuilt: %d", len(tt.Root))
	}

	// nil invalidate must not panic.
	tt2 := toolkit.NewTreeTable(cols, nil)
	u2 := BindTree(tt2, mvvm.NewObservableList("solo"), project, nil)
	if len(tt2.Root) != 1 {
		t.Fatalf("nil-invalidate seed: %d", len(tt2.Root))
	}
	u2()
}
