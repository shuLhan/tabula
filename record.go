// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabula

import (
	"math"
	"reflect"
	"strconv"
)

const (
	// TUndefined for undefined type
	TUndefined = -1
	// TString string type.
	TString = 0
	// TInteger integer type (64 bit).
	TInteger = 1
	// TReal float type (64 bit).
	TReal = 2
)

/*
Record represent the smallest building set of data-set.
*/
type Record struct {
	V interface{}
}

/*
NewRecord create new record from string with specific type.
Return record object or error when fail to convert the byte to type.
*/
func NewRecord(v string, t int) (r *Record, e error) {
	r = &Record{}

	e = r.SetValue(v, t)
	if e != nil {
		return nil, e
	}

	return
}

//
// NewRecordInt create new record from integer value.
//
func NewRecordInt(v int64) (r *Record) {
	return &Record{V: v}
}

/*
NewRecordReal create new record from float value.
*/
func NewRecordReal(v float64) (r *Record) {
	return &Record{
		V: v,
	}
}

/*
GetType of record.
*/
func (r *Record) GetType() int {
	switch r.V.(type) {
	case int64:
		return TInteger
	case float64:
		return TReal
	}
	return TString
}

/*
SetValue set the record values from string. If value can not be converted
to type, it will return an error.
*/
func (r *Record) SetValue(v string, t int) error {
	switch t {
	case TString:
		r.V = v

	case TInteger:
		i64, e := strconv.ParseInt(v, 10, 64)
		if nil != e {
			return e
		}

		r.V = i64

	case TReal:
		f64, e := strconv.ParseFloat(v, 64)
		if nil != e {
			return e
		}

		r.V = f64
	}
	return nil
}

/*
SetString will set the record content with string value.
*/
func (r *Record) SetString(v string) {
	r.V = v
}

/*
SetFloat will set the record content with float 64bit.
*/
func (r *Record) SetFloat(v float64) {
	r.V = v
}

/*
SetInteger will set the record value with integer 64bit.
*/
func (r *Record) SetInteger(v int64) {
	r.V = v
}

/*
Value return value of record based on their type.
*/
func (r *Record) Value() interface{} {
	switch r.V.(type) {
	case int64:
		return r.V.(int64)
	case float64:
		return r.V.(float64)
	}

	return r.V.(string)
}

/*
ToByte convert record value to byte.
*/
func (r *Record) ToByte() (b []byte) {
	switch r.V.(type) {
	case string:
		b = []byte(r.V.(string))

	case int64:
		b = []byte(strconv.FormatInt(r.V.(int64), 10))

	case float64:
		b = []byte(strconv.FormatFloat(r.V.(float64), 'f', -1, 64))
	}

	return b
}

/*
IsMissingValue check wether the value is a missing attribute.

If its string the missing value is indicated by character '?'.

If its integer the missing value is indicated by minimum negative integer, or
math.MinInt64.

If its real the missing value is indicated by -Inf.
*/
func (r *Record) IsMissingValue() bool {
	switch r.V.(type) {
	case string:
		str := r.V.(string)
		if str == "?" {
			return true
		}

	case int64:
		i64 := r.V.(int64)
		if i64 == math.MinInt64 {
			return true
		}

	case float64:
		f64 := r.V.(float64)
		return math.IsInf(f64, -1)
	}

	return false
}

/*
String convert record value to string.
*/
func (r Record) String() (s string) {
	switch r.V.(type) {
	case string:
		s = r.V.(string)

	case int64:
		s = strconv.FormatInt(r.V.(int64), 10)

	case float64:
		s = strconv.FormatFloat(r.V.(float64), 'f', -1, 64)
	}
	return
}

/*
Float convert given record to float value.
*/
func (r *Record) Float() (f64 float64) {
	var e error

	switch r.V.(type) {
	case string:
		f64, e = strconv.ParseFloat(r.V.(string), 64)

		if nil != e {
			f64 = math.Inf(-1)
		}

	case int64:
		f64 = float64(r.V.(int64))

	case float64:
		f64 = r.V.(float64)
	}

	return
}

/*
Integer convert given record to integer value.
*/
func (r *Record) Integer() (i64 int64) {
	var e error

	switch r.V.(type) {
	case string:
		i64, e = strconv.ParseInt(r.V.(string), 10, 64)

		if nil != e {
			i64 = math.MinInt64
		}

	case int64:
		i64 = r.V.(int64)

	case float64:
		i64 = int64(r.V.(float64))
	}

	return
}

//
// IsEqualToString return true if string representation of record value is
// equal to string `v`.
//
func (r *Record) IsEqualToString(v string) bool {
	if r.String() == v {
		return true
	}
	return false
}

//
// IsEqual return true if interface type and value equal to record type and
// value.
//
func (r *Record) IsEqual(v interface{}) bool {
	return reflect.DeepEqual(r.V, v)
}
