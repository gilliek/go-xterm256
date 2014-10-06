// Copyright 2014 The project AUTHORS. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math/rand"
	"time"

	"github.com/gilliek/go-xterm256/xterm256"
)

func main() {
	rand.Seed(time.Now().Unix())

	r := rand.Intn(6)
	g := rand.Intn(6)
	b := rand.Intn(6)

	randColor, _ := xterm256.NewColor(r, g, b)

	xterm256.Println(randColor, "Hello, World!")
}
