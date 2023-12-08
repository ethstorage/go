// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !math_pure_go

#include "textflag.h"

// func archFloor(x float64) float64
TEXT ·archFloor(SB),NOSPLIT,$0
	FMOVD	x+0(FP), F0
	FRINTMD	F0, F0
	FMOVD	F0, ret+8(FP)
	RET

// func archCeil(x float64) float64
TEXT ·archCeil(SB),NOSPLIT,$0
	FMOVD	x+0(FP), F0
	FRINTPD	F0, F0
	FMOVD	F0, ret+8(FP)
	RET

// func archTrunc(x float64) float64
TEXT ·archTrunc(SB),NOSPLIT,$0
	FMOVD	x+0(FP), F0
	FRINTZD	F0, F0
	FMOVD	F0, ret+8(FP)
	RET
