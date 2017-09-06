// Copyright 2017 M. Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package tabula_test

import (
	"github.com/shuLhan/tabula"
	"testing"
)

var data = []string{"9.987654321", "8.8", "7.7", "6.6", "5.5", "4.4", "3.3"}
var expFloat = []float64{9.987654321, 8.8, 7.7, 6.6, 5.5, 4.4, 3.3}

func initColReal(t *testing.T) (col *tabula.Column) {
	col = tabula.NewColumn(tabula.TReal, "TREAL")

	for x := range data {
		rec, e := tabula.NewRecordBy(data[x], tabula.TReal)
		if e != nil {
			t.Fatal(e)
		}

		col.PushBack(rec)
	}

	return col
}

func TestToFloatSlice(t *testing.T) {
	col := initColReal(t)
	got := col.ToFloatSlice()

	assert(t, expFloat, got, true)
}

func TestToStringSlice(t *testing.T) {
	var col tabula.Column

	for x := range data {
		rec, e := tabula.NewRecordBy(data[x], tabula.TString)
		if e != nil {
			t.Fatal(e)
		}

		col.PushBack(rec)
	}

	got := col.ToStringSlice()

	assert(t, data, got, true)
}

func TestDeleteRecordAt(t *testing.T) {
	var exp []float64
	del := 2

	exp = append(exp, expFloat[:del]...)
	exp = append(exp, expFloat[del+1:]...)

	col := initColReal(t)
	col.DeleteRecordAt(del)
	got := col.ToFloatSlice()

	assert(t, exp, got, true)
}
