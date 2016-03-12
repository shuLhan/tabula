// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabula

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
DatasetInterface is the interface for working with DSV data.
*/
type DatasetInterface interface {
	Clone() DatasetInterface
	Reset() error
	GetMode() int
	SetMode(mode int)
	GetNColumn() int
	GetNRow() int
	SetColumnsType(types []int)
	GetColumnsType() []int
	GetColumnTypeAt(colidx int) (int, error)
	SetColumnsName(names []string)
	GetColumnsName() []string

	AddColumn(tipe int, name string, vs []string)
	GetColumn(idx int) *Column
	GetColumns() *Columns
	GetColumnByName(name string) *Column
	GetRow(idx int) *Row
	GetRows() *Rows
	GetData() interface{}
	GetDataAsRows() Rows
	GetDataAsColumns() Columns
	TransposeToColumns()
	TransposeToRows()

	PushRow(r Row)
	PushRowToColumns(r Row)
	PushColumn(col Column)

	MergeColumns(DatasetInterface)
	MergeRows(DatasetInterface)
}

/*
ReadDatasetConfig open dataset configuration file and initialize dataset field
from there.
*/
func ReadDatasetConfig(ds interface{}, fcfg string) (e error) {
	cfg, e := ioutil.ReadFile(fcfg)

	if nil != e {
		return e
	}

	return json.Unmarshal(cfg, ds)
}

/*
SortColumnsByIndex will sort all columns using sorted index.
*/
func SortColumnsByIndex(di DatasetInterface, sortedIdx []int) {
	if di.GetMode() == DatasetModeRows {
		di.TransposeToColumns()
	}

	cols := di.GetColumns()
	for x, col := range *cols {
		colsorted := col.Records.SortByIndex(sortedIdx)
		(*cols)[x].SetRecords(colsorted)
	}
}

/*
SplitRowsByNumeric will split the data using splitVal in column `colidx`.

For example, given two continuous attribute,

	A: {1,2,3,4}
	B: {5,6,7,8}

if colidx is (1) B and splitVal is 7, the data will splitted into left set

	A': {1,2}
	B': {5,6}

and right set

	A'': {3,4}
	B'': {7,8}
*/
func SplitRowsByNumeric(di DatasetInterface, colidx int, splitVal float64) (
	splitLess DatasetInterface,
	splitGreater DatasetInterface,
	e error,
) {
	// check type of column
	coltype, e := di.GetColumnTypeAt(colidx)
	if e != nil {
		return
	}

	if !(coltype == TInteger || coltype == TReal) {
		return splitLess, splitGreater, ErrInvalidColType
	}

	// Should we convert the data mode back later.
	orgmode := di.GetMode()

	if orgmode == DatasetModeColumns {
		di.TransposeToRows()
	}

	if DEBUG >= 2 {
		fmt.Println("[tabula] dataset:", di)
	}

	splitLess = di.Clone()
	splitGreater = di.Clone()

	for _, row := range *di.GetRows() {
		if row[colidx].Float() < splitVal {
			splitLess.PushRow(row)
		} else {
			splitGreater.PushRow(row)
		}
	}

	if DEBUG >= 2 {
		fmt.Println("[tabula] split less:", splitLess)
		fmt.Println("[tabula] split greater:", splitGreater)
	}

	switch orgmode {
	case DatasetModeColumns:
		di.TransposeToColumns()
		splitLess.TransposeToColumns()
		splitGreater.TransposeToColumns()
	case DatasetModeMatrix:
		// do nothing, since its already filled when pushing new row.
	}

	return
}

/*
SplitRowsByCategorical will split the data using a set of split value in column
`colidx`.

For example, given two attributes,

	X: [A,B,A,B,C,D,C,D]
	Y: [1,2,3,4,5,6,7,8]

if colidx is (0) or A and split value is a set `[A,C]`, the data will splitted
into left set which contain all rows that have A or C,

	X': [A,A,C,C]
	Y': [1,3,5,7]

and the right set, excluded set, will contain all rows which is not A or C,

	X'': [B,B,D,D]
	Y'': [2,4,6,8]
*/
func SplitRowsByCategorical(di DatasetInterface, colidx int,
	splitVal []string) (
	splitIn DatasetInterface,
	splitEx DatasetInterface,
	e error,
) {
	// check type of column
	coltype, e := di.GetColumnTypeAt(colidx)
	if e != nil {
		return
	}

	if coltype != TString {
		return splitIn, splitEx, ErrInvalidColType
	}

	// should we convert the data mode back?
	orgmode := di.GetMode()

	if orgmode == DatasetModeColumns {
		di.TransposeToRows()
	}

	splitIn = di.Clone()
	splitEx = di.Clone()

	found := false

	for _, row := range *di.GetRows() {
		found = false
		for _, val := range splitVal {
			if row[colidx].String() == val {
				splitIn.PushRow(row)
				found = true
				break
			}
		}
		if !found {
			splitEx.PushRow(row)
		}
	}

	// convert all dataset based on original
	switch orgmode {
	case DatasetModeColumns:
		di.TransposeToColumns()
		splitIn.TransposeToColumns()
		splitEx.TransposeToColumns()
	case DatasetModeMatrix, DatasetNoMode:
		splitIn.TransposeToColumns()
		splitEx.TransposeToColumns()
	}

	return
}

/*
SplitRowsByValue generic function to split data by value. This function will
split data using value in column `colidx`. If value is numeric it will return
any rows that have column value less than `value` in `splitL`, and any column
value greater or equal to `value` in `splitR`.
*/
func SplitRowsByValue(di DatasetInterface, colidx int, value interface{}) (
	splitL DatasetInterface,
	splitR DatasetInterface,
	e error,
) {
	coltype, e := di.GetColumnTypeAt(colidx)
	if e != nil {
		return
	}

	if coltype == TString {
		splitL, splitR, e = SplitRowsByCategorical(di, colidx,
			value.([]string))
	} else {
		var splitval float64

		switch value.(type) {
		case int:
			splitval = float64(value.(int))
		case int64:
			splitval = float64(value.(int64))
		case float32:
			splitval = float64(value.(float32))
		case float64:
			splitval = value.(float64)
		}

		splitL, splitR, e = SplitRowsByNumeric(di, colidx,
			splitval)
	}

	if e != nil {
		return nil, nil, e
	}

	return
}
