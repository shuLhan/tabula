// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabula

/*
ColumnInterface define methods for working with Column.
*/
type ColumnInterface interface {
	SetType(tipe int)
	SetName(name string)

	GetType() int
	GetName() string

	SetRecords(recs *Records)

	Interface() interface{}
}
