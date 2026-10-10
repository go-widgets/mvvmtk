# mvvmtk

[![Go Reference](https://pkg.go.dev/badge/github.com/go-widgets/mvvmtk.svg)](https://pkg.go.dev/github.com/go-widgets/mvvmtk)
[![License: BSD-3-Clause](https://img.shields.io/badge/License-BSD--3--Clause-blue.svg)](LICENSE)

Binding glue between [`go-widgets/mvvm`](https://github.com/go-widgets/mvvm)
(`Observable` / `Command` / `ObservableList`) and
[`go-widgets/toolkit`](https://github.com/go-widgets/toolkit) widgets.

## Why a separate module

`mvvm`'s core package is deliberately toolkit-agnostic: it binds through a
pointer to a widget's value field and a pointer to its callback slot, and
imports no widget package. `toolkit` does import that core (since v0.124.0:
its widgets expose state as `mvvm.Observable`s), but not the binders, which
would be a cycle. `mvvmtk` is the module that knows **both**, so an app wires
a ViewModel to a widget in a single call and never touches widget state fields
directly.

```go
unbind := mvvmtk.BindText(entry, vm.Query, win.Invalidate) // entry.Text ⇄ vm.Query
defer unbind()
```

Each helper is a thin, correct wrapper over the generic `mvvm` adapters
(`BindField` / `OneWay` / `BindList` / `BindCommand`) with the widget's real
field and callback names filled in — no business logic. Every helper returns an
`unbind func()` that detaches the binding and restores any prior callback.

## Helpers

| Helper | Widget field(s) | Direction |
| --- | --- | --- |
| `BindText(*SearchEntry, *Observable[string], invalidate)` | `Text` / `OnChange` | two-way |
| `BindEntryText(*Entry, *Observable[string], invalidate)` | `Text` / `OnChange` | two-way |
| `BindChecked(*CheckButton, *Observable[bool], invalidate)` | `Checked` / `OnToggle` | two-way |
| `BindSelectedIndex(*DropDown, *Observable[int], invalidate)` | `Selected` / `OnSelect` | two-way |
| `BindSpin(*SpinButton, *Observable[int], invalidate)` | `Value` / `OnChange` | two-way |
| `BindListSelection(*ListBox, *Observable[int], invalidate)` | `Selected` / `OnActivate` | two-way |
| `BindViewSwitcher(*ViewSwitcher, *Observable[int], invalidate)` | `Current` / `OnChange` | two-way |
| `BindLabel(*Label, *Observable[string], invalidate)` | `Text` | one-way |
| `BindProgress(*ProgressBar, *Observable[float64], invalidate)` | `Fraction` | one-way |
| `BindListItems[T](*ListBox, *ObservableList[T], project, invalidate)` | `Items` | list → widget |
| `BindDropDownOptions[T](*DropDown, *ObservableList[T], project, invalidate)` | `Options` | list → widget |
| `BindViews[T](*ViewSwitcher, *ObservableList[T], project, invalidate)` | `Views` | list → widget |
| `BindCommand(*Button, *Command, invalidate)` | `OnClick` + `Disabled` + `Style` greying | command |
| `BindTree[T](*TreeTable, *ObservableList[T], project, invalidate)` | `Root` forest | list → widget |

`BindCommand` binds the button's `Disabled` state to `!CanExecute`: a button whose
command cannot run takes no click, no Enter/Space and no keyboard focus. It also
greys the button as it always has — swapping its `Style` to `ButtonSecondary`
when the command cannot execute and restoring the original `Style` when it can.
The returned unbind puts `OnClick`, `Style` and `Disabled` back as they were.

`BindTree` rebuilds a `TreeTable.Root` (`[]*TreeTableNode` forest) from the
list; the caller's `project` owns each node's `Cells`/`Children`.

## Install

```sh
go get github.com/go-widgets/mvvmtk
```

## License

BSD-3-Clause. See [LICENSE](LICENSE). Copyright (c) 2026 the go-widgets/mvvmtk
authors.
