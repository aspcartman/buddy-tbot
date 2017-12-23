// Copyright 2016 Huan Du. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package darkside

func getg() *g

type stack struct {
	lo uintptr
	hi uintptr
}

// v1.9.2
//noinspection ALL
type g struct {
	stack       stack
	stackguard0 uintptr
	stackguard1 uintptr

	_panic *Panic
}
