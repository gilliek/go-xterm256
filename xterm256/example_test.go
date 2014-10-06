// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xterm256_test

import (
	"xterm256"
)

func ExampleXterm256() {
	// Use a predefined color
	xterm256.Println(xterm256.Red, "Foo")

	// Define a new color
	orange, _ := xterm256.NewColor(3, 1, 0)
	xterm256.Println(orange, "Bar")

	// Add a background color
	orange.SetBackground(5, 5, 5)

	// Output:
	// \033[38;5;9mFoo\033[0m
	// \033[38;5;214mBar\033[0m
	// \033[38;5;231m\033[38;5;255mBar\033[0m
}
