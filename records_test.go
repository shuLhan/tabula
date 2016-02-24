// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabula_test

import (
	"fmt"
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tabula/util/assert"
	"testing"
)

func TestSortByIndex(t *testing.T) {
	data := make(tabula.Records, 3)
	data[0], _ = tabula.NewRecord("3", tabula.TInteger)
	data[1], _ = tabula.NewRecord("2", tabula.TInteger)
	data[2], _ = tabula.NewRecord("1", tabula.TInteger)

	sortedIdx := []int{2, 1, 0}
	expect := []int{1, 2, 3}

	sorted := data.SortByIndex(sortedIdx)

	got := fmt.Sprint(sorted)
	exp := fmt.Sprint(expect)

	assert.Equal(t, exp, got)
}
