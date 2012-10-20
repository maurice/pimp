// Copyright 2009 The pimp Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// pimp is a simple command lint utility to calculate the "Performance IMProvement"
// of a `before` and `after` figure
//
// Example:
//
// $ pimp 436.8 90.5
// Before:  436.80
// After:   90.50
// Delta:   346.30
// Change:  79.28%
package main

import (
	"log"
	"os"
	"strconv"
)

func init() {
	log.SetPrefix("")
	log.SetFlags(0)
}

func parseInt(s string) (f float64) {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Fatalf("Not a number: %s", s)
	}
	return
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <before> <after>\n", os.Args[0])
	}
	before := parseInt(os.Args[1])
	after := parseInt(os.Args[2])
	log.Printf("Before:  %0.2f\nAfter:   %0.2f\n", before, after)
	delta := -(after - before)
	change := (delta / before) * 100
	log.Printf("Delta:   %0.2f\nChange:  %0.2f%%\n", delta, change)
}
