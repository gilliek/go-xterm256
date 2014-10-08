// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xterm256

import (
	"errors"
	"fmt"
	"io"
	"os"
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

	// Extra colors

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

// NewColor creates a new Color and initializes the ForegroundColor to the color
// corresponding to the given r (red), g (green) and b (blue) components.
//
// r, g and b must range from 0 to 5.
func NewColor(r, g, b int) (Color, error) {
	c := Color{ForegroundColor: -1, BackgroundColor: -1}
	return c, c.SetForeground(r, g, b)
}

// SetBackground sets the background color to the color corresponding to the
// given r (red), g (green) and b (blue) components.
//
// r, g and b must range from 0 to 5.
func (c *Color) SetBackground(r, g, b int) error {
	code, err := c.rgbToColorCode(r, g, b)
	if err != nil {
		return err
	}

	c.BackgroundColor = code
	return nil
}

// SetForeground sets the foreground color to the color corresponding to the
// given r (red), g (green) and b (blue) components.
//
// r, g and b must range from 0 to 5.
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
		return -1, errors.New("invalid RGB component")
	}

	return code, nil
}

// Fprint formats using the default formats for its operands and writes to w
// with the color c. Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
//
// This is just a wrapper for the fmt.Fprint function, adding the support of
// color.
func Fprint(c Color, w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, wrapString(c, fmt.Sprint(a...)))
}

// Fprintf formats according to a format specifier and writes to w with color c.
// It returns the number of bytes written and any write error encountered.
//
// This is just a wrapper for the fmt.Fprintf function, adding the support of
// color.
func Fprintf(c Color, w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, wrapString(c, fmt.Sprintf(format, a...)))
}

// Fprintln formats using the default formats for its operands and writes to w
// with color c. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered.
//
// This is just a wrapper for the fmt.Fprintln function, adding the support of
// color.
func Fprintln(c Color, w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, wrapString(c, fmt.Sprint(a...)))
}

// Print formats using the default formats for its operands and writes to
// standard output with color c. Spaces are added between operands when neither
// is a string. It returns the number of bytes written and any write error
// encountered.
//
// This is just a wrapper for the fmt.Print function, adding the support of
// color.
func Print(c Color, a ...interface{}) (n int, err error) {
	return fmt.Fprint(os.Stdout, wrapString(c, fmt.Sprint(a...)))
}

// Printf formats according to a format specifier and writes to standard output
// with color c. It returns the number of bytes written and any write error
// encountered.
//
// This is just a wrapper for the fmt.Printf function, adding the support of
// color.
func Printf(c Color, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprint(os.Stdout, wrapString(c, fmt.Sprintf(format, a...)))
}

// Println formats using the default formats for its operands and writes to
// standard output with color c. Spaces are always added between operands and a
// newline is appended. It returns the number of bytes written and any write
// error encountered.
//
// This is just a wrapper for the fmt.Println function, adding the support of
// color.
func Println(c Color, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, wrapString(c, fmt.Sprint(a...)))
}

// Sprint formats using the default formats for its operands and returns the
// resulting string surronded by the ANSI escape codes for the color c. Spaces
// are added between operands when neither is a string.
//
// This is just a wrapper for the fmt.Sprint function, adding the support of
// color.
func Sprint(c Color, a ...interface{}) string {
	return wrapString(c, fmt.Sprint(a...))
}

// Sprintf formats according to a format specifier and returns the resulting
// string surronded by the ANSI escape codes for the color c.
//
// This is just a wrapper for the fmt.Sprintf function, adding the support of
// color.
func Sprintf(c Color, format string, a ...interface{}) string {
	return wrapString(c, fmt.Sprintf(format, a...))
}

// Sprintln formats using the default formats for its operands and returns the
// resulting string surronded by the ANSI escape codes for the color c. Spaces
// are always added between operands and a newline is appended.
//
// This is just a wrapper for the fmt.Sprintf function, adding the support of
// color.
func Sprintln(c Color, a ...interface{}) string {
	return wrapString(c, fmt.Sprintln(a...))
}

// wrapString surrounds the given string str with the ANSI escape codes for the
// color c.
func wrapString(c Color, str string) string {
	return foregroundCode(c) + backgroundCode(c) + str + resetCode
}

// backgroundCode return the ANSI escape code for the background color c.
func backgroundCode(c Color) string {
	if c.BackgroundColor != -1 {
		return fmt.Sprintf("\033[48;5;%dm", c.BackgroundColor)
	}
	return ""
}

// foregroundCode return the ANSI escape code for the foreground color c.
func foregroundCode(c Color) string {
	if c.ForegroundColor != -1 {
		return fmt.Sprintf("\033[38;5;%dm", c.ForegroundColor)
	}
	return ""
}
