// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabula_test

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tabula/util/assert"
	"testing"
)

var data = []string{"9.987654321", "8.8", "7.7", "6.6", "5.5", "4.4", "3.3"}
var expFloat = []float64{9.987654321, 8.8, 7.7, 6.6, 5.5, 4.4, 3.3}

func TestToFloatSlice(t *testing.T) {
	var col tabula.Column

	for x := range data {
		rec, e := tabula.NewRecordBy(data[x], tabula.TReal)
		if e != nil {
			t.Fatal(e)
		}

		col.PushBack(rec)
	}

	got := col.ToFloatSlice()

	assert.Equal(t, expFloat, got)
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

	assert.Equal(t, data, got)
}
