// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabula_test

import (
	"fmt"
	"github.com/shuLhan/tabula"
	"testing"
)

/*
TestRecord simply check how the stringer work.
*/
func TestRecord(t *testing.T) {
	expec := []string{"test", "1", "2"}
	expType := []int{tabula.TString, tabula.TInteger, tabula.TInteger}

	row := make(tabula.Row, 0)

	for i := range expec {
		r, e := tabula.NewRecordBy(expec[i], expType[i])
		if nil != e {
			t.Error(e)
		}

		row = append(row, r)
	}

	exp := fmt.Sprint(expec)
	got := fmt.Sprint(row)
	assert(t, exp, got, true)
}
