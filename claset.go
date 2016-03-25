// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabula

import (
	"github.com/shuLhan/tekstus"
	"strconv"
)

/*
Claset define a dataset with class attribute.
*/
type Claset struct {
	// Dataset embedded, for implementing the dataset interface.
	Dataset
	// ClassIndex contain index for target classification in columns.
	ClassIndex int `json:"ClassIndex"`

	// vs contain a copy of value space.
	vs []string
	// counts number of value space in current set.
	counts []int

	// major contain the name of majority class in dataset.
	major string
	// minor contain the name of minority class in dataset.
	minor string
}

/*
NewClaset create and return new Claset object.
*/
func NewClaset(mode int, types []int, names []string) (claset *Claset) {
	claset = &Claset{
		ClassIndex: -1,
	}

	claset.Init(mode, types, names)

	return
}

/*
Clone return a copy of current claset object.
*/
func (claset *Claset) Clone() DatasetInterface {
	clone := Claset{
		ClassIndex: claset.GetClassIndex(),
		major:      claset.MajorityClass(),
		minor:      claset.MinorityClass(),
	}
	clone.SetDataset(claset.GetDataset().Clone())
	return &clone
}

/*
GetDataset return the dataset.
*/
func (claset *Claset) GetDataset() DatasetInterface {
	return &claset.Dataset
}

/*
GetClassValueSpace return the class value space.
*/
func (claset *Claset) GetClassValueSpace() []string {
	return claset.Columns[claset.ClassIndex].ValueSpace
}

/*
GetClassColumn return dataset class values in column.
*/
func (claset *Claset) GetClassColumn() *Column {
	if claset.Mode == DatasetModeRows {
		claset.TransposeToColumns()
	}
	return &claset.Columns[claset.ClassIndex]
}

/*
GetClassAsStrings return all class values as slice of string.
*/
func (claset *Claset) GetClassAsStrings() []string {
	if claset.Mode == DatasetModeRows {
		claset.TransposeToColumns()
	}
	return claset.Columns[claset.ClassIndex].ToStringSlice()
}

/*
GetClassIndex return index of class attribute in dataset.
*/
func (claset *Claset) GetClassIndex() int {
	return claset.ClassIndex
}

/*
MajorityClass return the majority class of data.
*/
func (claset *Claset) MajorityClass() string {
	return claset.major
}

/*
MinorityClass return the minority class in dataset.
*/
func (claset *Claset) MinorityClass() string {
	return claset.minor
}

/*
SetDataset in class set.
*/
func (claset *Claset) SetDataset(dataset DatasetInterface) {
	claset.Dataset = *(dataset.(*Dataset))
}

/*
SetClassIndex will set the class index to `v`.
*/
func (claset *Claset) SetClassIndex(v int) {
	claset.ClassIndex = v
}

/*
SetMajorityClass will set the majority class to `v`.
*/
func (claset *Claset) SetMajorityClass(v string) {
	claset.major = v
}

/*
SetMinorityClass will set the minority class to `v`.
*/
func (claset *Claset) SetMinorityClass(v string) {
	claset.minor = v
}

//
// CountValueSpaces will count number of value space in current dataset.
//
func (claset *Claset) CountValueSpaces() {
	classv := claset.GetClassAsStrings()
	claset.vs = claset.GetClassValueSpace()

	claset.counts = tekstus.WordsCountTokens(classv, claset.vs, false)
}

//
// RecountMajorMinor recount major and minor class in claset.
//
func (claset *Claset) RecountMajorMinor() {
	claset.CountValueSpaces()

	_, maxIdx := tekstus.IntFindMax(claset.counts)
	_, minIdx := tekstus.IntFindMin(claset.counts)

	if maxIdx >= 0 {
		claset.major = claset.vs[maxIdx]
	}
	if minIdx >= 0 {
		claset.minor = claset.vs[minIdx]
	}
}

/*
IsInSingleClass check whether all target class contain only single value.
Return true and name of target if all rows is in the same class,
false and empty string otherwise.
*/
func (claset *Claset) IsInSingleClass() (single bool, class string) {
	classv := claset.GetClassAsStrings()

	for i, t := range classv {
		if i == 0 {
			single = true
			class = t
			continue
		}
		if t != class {
			return false, ""
		}
	}
	return
}

//
// String, yes it will pretty print the meta-data in JSON format.
//
func (claset *Claset) String() (s string) {
	if claset.vs == nil {
		claset.RecountMajorMinor()
	}

	s = `"claset": {
	"rows"    : ` + strconv.Itoa(claset.Len()) + `
,	"columns" : ` + strconv.Itoa(claset.GetNColumn()) + `
,	"vs"      :{
`

	for x, v := range claset.vs {
		if x > 0 {
			s += "\t,"
		} else {
			s += "\t"
		}
		s += "\t\"" + v + "\"\t:" + strconv.Itoa(claset.counts[x]) + "\n"
	}

	s += `	}
,	"major" :` + claset.major + `
,	"minor" :` + claset.minor + `
}`

	return
}
