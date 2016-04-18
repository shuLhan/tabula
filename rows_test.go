// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabula_test

import (
	"fmt"
	"strings"
	"testing"
)

var exp = []string{
	"0\n",
	"1\n",
	"2\n",
	"3\n",
	"4\n",
}

func TestPushBack(t *testing.T) {
	rows, e := initRows()
	if e != nil {
		t.Fatal(e)
	}

	exp := strings.Join(rowsExpect, "")
	got := fmt.Sprint(rows)

	assert(t, exp, got, true)
}

func TestPopFront(t *testing.T) {
	rows, e := initRows()
	if e != nil {
		t.Fatal(e)
	}

	l := len(rows) - 1
	for i := range rows {
		row := rows.PopFront()

		exp := rowsExpect[i]
		got := fmt.Sprint(row)

		assert(t, exp, got, true)

		if i < l {
			exp = strings.Join(rowsExpect[i+1:], "")
		} else {
			exp = ""
		}
		got = fmt.Sprint(rows)

		assert(t, exp, got, true)
	}

	// empty rows
	row := rows.PopFront()

	exp := "<nil>"
	got := fmt.Sprint(row)

	assert(t, exp, got, true)
}

func TestPopFrontRow(t *testing.T) {
	rows, e := initRows()
	if e != nil {
		t.Fatal(e)
	}

	l := len(rows) - 1
	for i := range rows {
		newRows := rows.PopFrontAsRows()

		exp := rowsExpect[i]
		got := fmt.Sprint(newRows)

		assert(t, exp, got, true)

		if i < l {
			exp = strings.Join(rowsExpect[i+1:], "")
		} else {
			exp = ""
		}
		got = fmt.Sprint(rows)

		assert(t, exp, got, true)
	}

	// empty rows
	row := rows.PopFrontAsRows()

	exp := ""
	got := fmt.Sprint(row)

	assert(t, exp, got, true)
}

func TestGroupByValue(t *testing.T) {
	rows, e := initRows()
	if e != nil {
		t.Fatal(e)
	}

	mapRows := rows.GroupByValue(testClassIdx)

	got := fmt.Sprint(mapRows)

	assert(t, groupByExpect, got, true)
}

func TestRandomPick(t *testing.T) {
	rows, e := initRows()
	if e != nil {
		t.Fatal(e)
	}

	// random pick with duplicate
	for i := 0; i < 5; i++ {
		picked, unpicked, pickedIdx, unpickedIdx := rows.RandomPick(6,
			true)

		// check if unpicked item exist in picked items.
		isin, _ := picked.Contains(unpicked)

		if isin {
			fmt.Println("Random pick with duplicate rows")
			fmt.Println("==> picked rows   :", picked)
			fmt.Println("==> picked idx    :", pickedIdx)
			fmt.Println("==> unpicked rows :", unpicked)
			fmt.Println("==> unpicked idx  :", unpickedIdx)
			t.Fatal("random pick: unpicked is false")
		}
	}

	// random pick without duplication
	for i := 0; i < 5; i++ {
		picked, unpicked, pickedIdx, unpickedIdx := rows.RandomPick(3,
			false)

		// check if picked rows is duplicate
		assert(t, picked[0], picked[1], false)

		// check if unpicked item exist in picked items.
		isin, _ := picked.Contains(unpicked)

		if isin {
			fmt.Println("Random pick with no duplicate rows")
			fmt.Println("==> picked rows   :", picked)
			fmt.Println("==> picked idx    :", pickedIdx)
			fmt.Println("==> unpicked rows :", unpicked)
			fmt.Println("==> unpicked idx  :", unpickedIdx)
			t.Fatal("random pick: unpicked is false")
		}
	}
}

func TestRowsDel(t *testing.T) {
	rows, e := initRows()
	if e != nil {
		t.Fatal(e)
	}

	rows.Del(0)

	exp := strings.Join(rowsExpect[1:], "")
	got := fmt.Sprint(rows)

	assert(t, exp, got, true)
}
