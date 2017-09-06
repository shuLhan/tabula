// Copyright 2017 M. Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package tabula_test

import (
	"os"
	"reflect"
	"runtime"
	"testing"

	"github.com/shuLhan/tabula"
)

var (
	traces = make([]byte, 1024)
)

func printStackTrace() {
	var lines, start, end int

	for x, b := range traces {
		if b != '\n' {
			continue
		}
		lines++
		if lines == 3 {
			start = x
		} else if lines == 5 {
			end = x + 1
			break
		}
	}

	os.Stderr.Write(traces[start:end])
}

func assert(t *testing.T, exp, got interface{}, equal bool) {
	if reflect.DeepEqual(exp, got) != equal {
		runtime.Stack(traces, true)
		printStackTrace()
		t.Fatalf("\n"+
			">>> Expecting '%v'\n"+
			"          got '%v'\n", exp, got)
	}
}

var testColTypes = []int{
	tabula.TInteger,
	tabula.TInteger,
	tabula.TInteger,
	tabula.TString,
}

var testColNames = []string{"int01", "int02", "int03", "class"}

// Testing data and function for Rows and MapRows
var rowsData = [][]string{
	{"1", "5", "9", "+"},
	{"2", "6", "0", "-"},
	{"3", "7", "1", "-"},
	{"4", "8", "2", "+"},
}

var testClassIdx = 3

var rowsExpect = []string{
	"&[1 5 9 +]",
	"&[2 6 0 -]",
	"&[3 7 1 -]",
	"&[4 8 2 +]",
}

var groupByExpect = "[{+ &[1 5 9 +]&[4 8 2 +]} {- &[2 6 0 -]&[3 7 1 -]}]"

func initRows() (rows tabula.Rows, e error) {
	for i := range rowsData {
		l := len(rowsData[i])
		row := make(tabula.Row, 0)

		for j := 0; j < l; j++ {
			rec, e := tabula.NewRecordBy(rowsData[i][j],
				testColTypes[j])

			if nil != e {
				return nil, e
			}

			row = append(row, rec)
		}

		rows.PushBack(&row)
	}
	return rows, nil
}
