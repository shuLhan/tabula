// Copyright 2017 M. Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package tabula_test

import (
	"github.com/shuLhan/tabula"
	"testing"
)

func BenchmarkPushRow(b *testing.B) {
	dataset := tabula.NewDataset(tabula.DatasetModeRows, nil, nil)

	for i := 0; i < b.N; i++ {
		e := populateWithRows(dataset)
		if e != nil {
			b.Fatal(e)
		}
	}
}
