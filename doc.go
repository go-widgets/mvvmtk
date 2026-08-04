// Copyright (c) 2026 the go-widgets/mvvmtk authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

// Package mvvmtk is the binding glue between github.com/go-widgets/mvvm
// (Observable / Command / ObservableList) and github.com/go-widgets/toolkit
// widgets.
//
// mvvm is deliberately toolkit-agnostic: it binds through a pointer to a
// widget's value field and a pointer to its callback slot, and never imports
// any widget package. toolkit, in turn, never imports mvvm. This module is the
// one place that knows BOTH, so an app can wire a ViewModel to a widget in a
// single call and never touch widget state fields directly:
//
//	unbind := mvvmtk.BindText(entry, vm.Query, win.Invalidate)
//
// Each helper is a thin, correct wrapper over the generic mvvm adapters
// (BindField / OneWay / BindList / BindCommand) with the widget's real field
// and callback names filled in — no business logic. Every helper returns an
// unbind func that detaches the binding and restores any prior callback.
package mvvmtk
