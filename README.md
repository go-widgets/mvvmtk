# mvvmtk

[![Go Reference](https://pkg.go.dev/badge/github.com/go-widgets/mvvmtk.svg)](https://pkg.go.dev/github.com/go-widgets/mvvmtk)
[![License: BSD-3-Clause](https://img.shields.io/badge/License-BSD--3--Clause-blue.svg)](LICENSE)

Binding glue between [`go-widgets/mvvm`](https://github.com/go-widgets/mvvm)
(`Observable` / `Command` / `ObservableList`) and
[`go-widgets/toolkit`](https://github.com/go-widgets/toolkit) widgets.

## Why a separate module

`mvvm` is deliberately toolkit-agnostic: it binds through a pointer to a
widget's value field and a pointer to its callback slot, and never imports any
widget package. `toolkit`, in turn, never imports `mvvm`. `mvvmtk` is the one
module that knows **both**, so an app wires a ViewModel to a widget in a single
call and never touches widget state fields directly.

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
| `BindCommand(*Button, *Command, invalidate)` | `OnClick` + `Style` greying | command |
| `BindTree[T](*TreeTable, *ObservableList[T], project, invalidate)` | `Root` forest | list → widget |

`BindCommand` reflects `CanExecute` by greying the button — swapping its `Style`
to `ButtonSecondary` when the command cannot execute and restoring the original
`Style` when it can — because a `Button` has no boolean "disabled" field.

`BindTree` rebuilds a `TreeTable.Root` (`[]*TreeTableNode` forest) from the
list; the caller's `project` owns each node's `Cells`/`Children`.

## Install

```sh
go get github.com/go-widgets/mvvmtk
```

## License

BSD-3-Clause. See [LICENSE](LICENSE). Copyright (c) 2026 the go-widgets/mvvmtk
authors.
