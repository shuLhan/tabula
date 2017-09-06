// Copyright 2017 M. Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package tabula_test

import (
	"fmt"
	"github.com/shuLhan/tabula"
	"testing"
)

func TestSortByIndex(t *testing.T) {
	data := make(tabula.Records, 3)
	data[0] = tabula.NewRecordInt(3)
	data[1] = tabula.NewRecordInt(2)
	data[2] = tabula.NewRecordInt(1)

	sortedIdx := []int{2, 1, 0}
	expect := []int{1, 2, 3}

	sorted := data.SortByIndex(sortedIdx)

	got := fmt.Sprint(sorted)
	exp := fmt.Sprint(&expect)

	assert(t, exp, got, true)
}
