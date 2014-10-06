// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xterm256

import (
	"errors"
)

var (
	// Define the 16 system colors. See http://pln.jonas.me/xterm-colors

	Black       = Color{ForegroundColor: 0, BackgroundColor: -1}
	DarkRed     = Color{ForegroundColor: 1, BackgroundColor: -1}
	DarkGreen   = Color{ForegroundColor: 2, BackgroundColor: -1}
	DarkYellow  = Color{ForegroundColor: 3, BackgroundColor: -1}
	DarkBlue    = Color{ForegroundColor: 4, BackgroundColor: -1}
	DarkMagenta = Color{ForegroundColor: 5, BackgroundColor: -1}
	DarkCyan    = Color{ForegroundColor: 6, BackgroundColor: -1}
	LightGray   = Color{ForegroundColor: 7, BackgroundColor: -1}
	DarkGray    = Color{ForegroundColor: 8, BackgroundColor: -1}
	Red         = Color{ForegroundColor: 9, BackgroundColor: -1}
	Green       = Color{ForegroundColor: 10, BackgroundColor: -1}
	Yellow      = Color{ForegroundColor: 11, BackgroundColor: -1}
	Blue        = Color{ForegroundColor: 12, BackgroundColor: -1}
	Magenta     = Color{ForegroundColor: 13, BackgroundColor: -1}
	Cyan        = Color{ForegroundColor: 14, BackgroundColor: -1}
	White       = Color{ForegroundColor: 15, BackgroundColor: -1}

	Orange = Color{ForegroundColor: 130, BackgroundColor: -1}
)

const (
	// ANSI escape code for disabling the coloration.
	resetCode = "\033[0m"
)

// Color structure holds the foreground and background colors.
//
// Colors are defined using the Xterm color code. For more details, see
// http://pln.jonas.me/xterm-colors
type Color struct {
	ForegroundColor int
	BackgroundColor int
}

func NewColor(r, g, b int) (Color, error) {
	c := Color{ForegroundColor: -1, BackgroundColor: -1}
	return c, c.SetForeground(r, g, b)
}

func (c *Color) SetBackground(r, g, b int) error {
	code, err := c.rgbToColorCode(r, g, b)
	if err != nil {
		return err
	}

	c.BackgroundColor = code
	return nil
}

func (c *Color) SetForeground(r, g, b int) error {
	code, err := c.rgbToColorCode(r, g, b)
	if err != nil {
		return err
	}

	c.ForegroundColor = code
	return nil
}

// rgbToColorCode converts RGB value to Xterm 256 code.
// For more details, see: http://pln.jonas.me/xterm-colors
func (c Color) rgbToColorCode(r, g, b int) (int, error) {
	code := 16 + r*36 + g*6 + b

	if code < 0 || code > 255 {
		return -1, errors.New("Invalid RGB component!")
	}

	return code, nil
}
