// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xterm256

import (
	"testing"
)

func TestXterm256(t *testing.T) {
	c, err := NewColor(5, 5, 5)
	if err != nil {
		t.Fatal(err)
	}

	if c.ForegroundColor != 231 {
		t.Errorf("ForegroundColor: expected %d, found %d",
			255, c.ForegroundColor)
	}

	if c.BackgroundColor != -1 {
		t.Errorf("BackgroundColor: expected %d, found %d",
			-1, c.BackgroundColor)
	}

	if err := c.SetBackground(0, 0, 0); err != nil {
		t.Fatal(err)
	}

	if c.BackgroundColor != 16 {
		t.Fatalf("BackgroundColor: expected %d, found %d",
			16, c.BackgroundColor)
	}

	output := wrapString(c, "Foo")
	expectedOutput := "\033[38;5;231m\033[48;5;16mFoo\033[0m"
	if output != expectedOutput {
		t.Errorf("wrapString(c, 'Foo'): expected '%s', found '%s'",
			expectedOutput, output)
	}

	if err := c.SetBackground(42, 42, 42); err == nil {
		t.Error("SetBackground(42, 42, 42): expected failure, found nil")
	}

	c, err = NewColor(42, 42, 42)
	if err == nil {
		t.Error("NewColor(42, 42, 42): expected failure, found nil")
	}

	c = Color{ForegroundColor: -1, BackgroundColor: -1}

	if code := foregroundCode(c); code != "" {
		t.Errorf("foregroundCode(c): expected '%s', found '%s'",
			"", code)
	}

	if code := backgroundCode(c); code != "" {
		t.Errorf("backgroundCode(c): expected '%s', found '%s'",
			"", code)
	}
}
