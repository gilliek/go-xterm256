// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xterm256

import (
	"testing"
)

func TestNewColor(t *testing.T) {
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

	c, err = NewColor(42, 42, 42)
	if err == nil {
		t.Error("NewColor(42, 42, 42): expected failure, found nil")
	}
}
