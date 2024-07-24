// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18

package idna

// Transitional processing is disabled by default in Go 1.18.
// https://golang.org/issue/47510
const transitionalLookup = false
