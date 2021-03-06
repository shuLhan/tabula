// Copyright 2017 M. Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package tabula_test

import (
	"github.com/shuLhan/tabula"
	"testing"
)

var dataFloat64 = []float64{0.1, 0.2, 0.3, 0.4, 0.5}

func createRow() (row tabula.Row) {
	for _, v := range dataFloat64 {
		row.PushBack(tabula.NewRecordReal(v))
	}
	return
}

func TestClone(t *testing.T) {
	row := createRow()
	rowClone := row.Clone()
	rowClone2 := row.Clone()

	assert(t, &row, rowClone, true)

	// changing the clone value should not change the original copy.
	(*rowClone2)[0].SetFloat(0)
	assert(t, &row, rowClone, true)
	assert(t, &row, rowClone2, false)
}
