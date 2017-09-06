// Copyright 2017 M. Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package tabula_test

import (
	"fmt"
	"github.com/shuLhan/tabula"
	"testing"
)

func TestAddRow(t *testing.T) {
	mapRows := tabula.MapRows{}
	rows, e := initRows()

	if e != nil {
		t.Fatal(e)
	}

	for _, row := range rows {
		key := fmt.Sprint((*row)[testClassIdx].Interface())
		mapRows.AddRow(key, row)
	}

	got := fmt.Sprint(mapRows)

	assert(t, groupByExpect, got, true)
}

func TestGetMinority(t *testing.T) {
	mapRows := tabula.MapRows{}
	rows, e := initRows()

	if e != nil {
		t.Fatal(e)
	}

	for _, row := range rows {
		key := fmt.Sprint((*row)[testClassIdx].Interface())
		mapRows.AddRow(key, row)
	}

	// remove the first row in the first key, so we can make it minority.
	mapRows[0].Value.PopFront()

	_, minRows := mapRows.GetMinority()

	exp := rowsExpect[3]
	got := fmt.Sprint(minRows)

	assert(t, exp, got, true)
}
