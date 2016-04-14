// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabula

/*
Row represent slice of record.
*/
type Row []*Record

//
// Len return number of record in row.
//
func (row *Row) Len() int {
	return len(*row)
}

/*
PushBack will add new record to the end of row.
*/
func (row *Row) PushBack(r *Record) {
	*row = append(*row, r)
}

/*
GetTypes return type of all records.
*/
func (row *Row) GetTypes() (types []int) {
	for _, r := range *row {
		types = append(types, r.GetType())
	}
	return
}

//
// Clone create and return a clone of row.
//
func (row *Row) Clone() *Row {
	clone := make(Row, len(*row))

	for x, rec := range *row {
		clone[x] = rec.Clone()
	}
	return &clone
}

/*
IsNilAt return true if there is no record value in row at `idx`, otherwise
return false.
*/
func (row *Row) IsNilAt(idx int) bool {
	if idx < 0 {
		return true
	}
	if idx >= len(*row) {
		return true
	}
	if (*row)[idx] == nil {
		return true
	}
	return (*row)[idx].IsNil()
}

/*
SetValueAt will set the value of row at cell index `idx` with record `rec`.
*/
func (row *Row) SetValueAt(idx int, rec *Record) {
	(*row)[idx] = rec
}

//
// GetValueAt return the value of row record at index `idx`. If the index is
// out of range it will return nil and false
//
func (row *Row) GetValueAt(idx int) (interface{}, bool) {
	if row.Len() <= idx {
		return nil, false
	}
	return (*row)[idx].Value(), true
}

//
// GetIntAt return the integer value of row record at index `idx`.
// If the index is out of range it will return 0 and false.
//
func (row *Row) GetIntAt(idx int) (int64, bool) {
	if row.Len() <= idx {
		return 0, false
	}

	return (*row)[idx].Integer(), true
}

//
// IsEqual return true if row content equal with `other` row, otherwise return
// false.
//
func (row *Row) IsEqual(other *Row) bool {
	if len(*row) != len(*other) {
		return false
	}
	for x, xrec := range *row {
		if !xrec.IsEqual((*other)[x]) {
			return false
		}
	}
	return true
}
