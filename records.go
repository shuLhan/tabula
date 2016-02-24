// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabula

/*
Records define slice of pointer to Record.
*/
type Records []*Record

/*
SortByIndex sort record in column by index.
*/
func (recs Records) SortByIndex(sortedIdx []int) (sorted []*Record) {
	sorted = make([]*Record, len(recs))

	for i := range sortedIdx {
		sorted[i] = recs[sortedIdx[i]]
	}
	return
}
