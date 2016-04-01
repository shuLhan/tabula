// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package util contain common function to work with data.
*/
package util

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	// SortThreshold for sort. When the data less than SortThreshold,
	// insertion sort will be used to replace the sort.
	SortThreshold = 7
)

var (
	// DEBUG level, can be set from environment using TABULA_UTIL_DEBUG.
	DEBUG = 0
)

func init() {
	v := os.Getenv("TABULA_UTIL_DEBUG")
	if v == "" {
		DEBUG = 0
	} else {
		DEBUG, _ = strconv.Atoi(v)
	}
}

/*
InsertionSortFloat64 will sort the data using insertion-sort algorithm.
*/
func InsertionSortFloat64(data []float64, idx []int, l, r int) {
	for x := l; x < r; x++ {
		for y := x + 1; y < r; y++ {
			if data[x] > data[y] {
				SwapInt(idx, x, y)
				SwapFloat64(data, x, y)
			}
		}
	}
}

/*
SwapInt swap two indices value of integer.
*/
func SwapInt(data []int, i, j int) {
	if i == j {
		return
	}

	tmp := data[i]
	data[i] = data[j]
	data[j] = tmp
}

/*
SwapFloat64 swap two indices value of 64bit float.
*/
func SwapFloat64(data []float64, i, j int) {
	if i == j {
		return
	}

	tmp := data[i]
	data[i] = data[j]
	data[j] = tmp
}

/*
SwapString swap two indices value of string.
*/
func SwapString(data []string, i, j int) {
	if i == j {
		return
	}

	tmp := data[i]
	data[i] = data[j]
	data[j] = tmp
}

//
// InplaceMergesortFloat64 in-place merge-sort without memory allocation.
//
// Algorithm,
//
// (0) If data length == Threshold, then
// (0.1) use insertion sort.
// (1) Divide into left and right.
// (2) Sort left.
// (3) Sort right.
// (4) Merge sorted left and right.
// (4.1) If the last element of the left is lower then the first element of the
//       right, i.e. [1 2] [3 4]; no merging needed, return immediately.
// (4.2) Let x be the first index of left-side, and y be the first index of
//       the right-side.
// (4.3) Loop until either x or y reached the maximum slice.
// (4.3.1) IF DATA[x] <= DATA[y]
// (4.3.1.1) INCREMENT x
// (4.3.1.2) IF x > y THEN GOTO 4.3
// (4.3.1.3) GOTO 4.3.4
// (4.3.2) LET YLAST := the next DATA[y] that is less DATA[x]
// (4.3.3) SWAP DATA, X, Y, YLAST
// (4.3.4) LET Y := the next DATA that has minimum value between x and r
//
func InplaceMergesortFloat64(data []float64, idx []int, l, r int) {
	// (0)
	if l+SortThreshold >= r {
		// (0.1)
		InsertionSortFloat64(data, idx, l, r)
		return
	}

	// (1)
	res := (r + l) % 2
	c := (r + l) / 2
	if res == 1 {
		c++
	}

	// (2)
	InplaceMergesortFloat64(data, idx, l, c)

	// (3)
	InplaceMergesortFloat64(data, idx, c, r)

	// (4)
	if data[c-1] <= data[c] {
		// (4.1)
		return
	}

	// (4.2)
	x := l
	y := c
	ylast := c

	// (4.3)
	for x < r && y < r {
		// (4.3.1)
		if data[x] <= data[y] {
			x++

			// (4.3.1.2)
			if x >= y {
				goto next
			}

			// (4.3.1.3)
			continue
		}

		// (4.3.2)
		ylast = movey(data, x, y, r)

		// (4.3.3)
		ylast = multiswap(data, idx, x, y, ylast)

	next:
		// (4.3.4)
		for x < r {
			y = min(data, x, r)
			if y == x {
				x++
			} else {
				break
			}
		}
	}
}

func movey(data []float64, x, y, r int) int {
	yorg := y
	y++
	for y < r {
		if data[y] >= data[x] {
			break
		}
		if data[y] < data[yorg] {
			break
		}
		y++
	}
	return y
}

func multiswap(data []float64, idx []int, x, y, ylast int) int {
	for y < ylast {
		SwapInt(idx, x, y)
		SwapFloat64(data, x, y)
		x++
		y++
		if y >= ylast {
			return y
		}
		if data[x] <= data[y] {
			return y
		}
	}

	return y
}

func min(data []float64, l, r int) (m int) {
	min := data[l]
	m = l
	for l++; l < r; l++ {
		if data[l] <= min {
			min = data[l]
			m = l
		}
	}
	return
}

/*
IndirectSortFloat64 will sort the data and return the sorted index.
*/
func IndirectSortFloat64(data []float64) (sortedIdx []int) {
	datalen := len(data)

	sortedIdx = make([]int, datalen)
	for i := 0; i < datalen; i++ {
		sortedIdx[i] = i
	}

	InplaceMergesortFloat64(data, sortedIdx, 0, datalen)

	return
}

/*
SortFloatSliceByIndex will sort the slice of float `data` using sorted index
`sortedIdx`.
*/
func SortFloatSliceByIndex(data *[]float64, sortedIdx *[]int) {
	newdata := make([]float64, len(*data))

	for i := range *sortedIdx {
		newdata[i] = (*data)[(*sortedIdx)[i]]
	}

	(*data) = newdata
}

/*
SortStringSliceByIndex will sort the slice of string `data` using sorted index
`sortedIdx`.
*/
func SortStringSliceByIndex(data *[]string, sortedIdx *[]int) {
	newdata := make([]string, len(*data))

	for i := range *sortedIdx {
		newdata[i] = (*data)[(*sortedIdx)[i]]
	}

	(*data) = newdata
}

/*
IntIsExist check if integer value exist in list of integer, return true if
exist, false otherwise.
*/
func IntIsExist(data []int, val int) bool {
	for _, v := range data {
		if val == v {
			return true
		}
	}
	return false
}

/*
GetRandomInteger return random integer value from 0 to maximum value `maxVal`.

The random value is checked with already picked index, `pickedIdx`.

If `dup` is true, allow duplicate value in `pickedIdx`, otherwise only single
unique value allowed in `pickedIdx`.

If excluding index `excIdx` is not empty, do not pick the integer value listed
in there.
*/
func GetRandomInteger(maxVal int, dup bool, pickedIdx []int, excIdx []int) (
	idx int,
) {
	rand.Seed(time.Now().UnixNano())

	for {
		idx = rand.Intn(maxVal)

		// check if its must not be selected
		excluded := false
		for _, excIdx := range excIdx {
			if idx == excIdx {
				excluded = true
				break
			}
		}
		if excluded {
			continue
		}

		if dup {
			// allow duplicate idx
			return
		}

		// check if its already picked
		isPicked := false
		for _, pastIdx := range pickedIdx {
			if idx == pastIdx {
				isPicked = true
				break
			}
		}
		// get another random idx again
		if isPicked {
			continue
		}

		// bingo, we found unique idx that has not been picked.
		return
	}
}

//
// IntCreateSequence will create and return sequence of integer from `min` to
// `max` value.
//
// E.g. if min is 0 and max is 5 then it will return `[0 1 2 3 4 5]`.
//
func IntCreateSequence(min, max int64) (seq []int64) {
	for ; min <= max; min++ {
		seq = append(seq, min)
	}
	return
}
