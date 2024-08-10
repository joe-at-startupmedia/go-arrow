// Code generated by column_reader_types.gen.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file

import (
	"unsafe"

	"github.com/joe-at-startupmedia/go-arrow/arrow"
	"github.com/joe-at-startupmedia/go-arrow/parquet"
	"github.com/joe-at-startupmedia/go-arrow/parquet/internal/encoding"
)

// Int32ColumnChunkReader is the Typed Column chunk reader instance for reading
// Int32 column data.
type Int32ColumnChunkReader struct {
	columnChunkReader
}

// Skip skips the next nvalues so that the next call to ReadBatch
// will start reading *after* the skipped values.
func (cr *Int32ColumnChunkReader) Skip(nvalues int64) (int64, error) {
	return cr.columnChunkReader.skipValues(nvalues,
		func(batch int64, buf []byte) (int64, error) {
			vals, _, err := cr.ReadBatch(batch,
				arrow.Int32Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf))
			return vals, err
		})
}

// ReadBatch reads batchSize values from the column.
//
// Returns error if values is not at least big enough to hold the number of values that will be read.
//
// defLvls and repLvls can be nil, or will be populated if not nil. If not nil, they must be
// at least large enough to hold the number of values that will be read.
//
// total is the number of rows that were read, valuesRead is the actual number of physical values
// that were read excluding nulls
func (cr *Int32ColumnChunkReader) ReadBatch(batchSize int64, values []int32, defLvls, repLvls []int16) (total int64, valuesRead int, err error) {
	return cr.readBatch(batchSize, defLvls, repLvls, func(start, len int64) (int, error) {
		return cr.curDecoder.(encoding.Int32Decoder).Decode(values[start : start+len])
	})
}

// Int64ColumnChunkReader is the Typed Column chunk reader instance for reading
// Int64 column data.
type Int64ColumnChunkReader struct {
	columnChunkReader
}

// Skip skips the next nvalues so that the next call to ReadBatch
// will start reading *after* the skipped values.
func (cr *Int64ColumnChunkReader) Skip(nvalues int64) (int64, error) {
	return cr.columnChunkReader.skipValues(nvalues,
		func(batch int64, buf []byte) (int64, error) {
			vals, _, err := cr.ReadBatch(batch,
				arrow.Int64Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf))
			return vals, err
		})
}

// ReadBatch reads batchSize values from the column.
//
// Returns error if values is not at least big enough to hold the number of values that will be read.
//
// defLvls and repLvls can be nil, or will be populated if not nil. If not nil, they must be
// at least large enough to hold the number of values that will be read.
//
// total is the number of rows that were read, valuesRead is the actual number of physical values
// that were read excluding nulls
func (cr *Int64ColumnChunkReader) ReadBatch(batchSize int64, values []int64, defLvls, repLvls []int16) (total int64, valuesRead int, err error) {
	return cr.readBatch(batchSize, defLvls, repLvls, func(start, len int64) (int, error) {
		return cr.curDecoder.(encoding.Int64Decoder).Decode(values[start : start+len])
	})
}

// Int96ColumnChunkReader is the Typed Column chunk reader instance for reading
// Int96 column data.
type Int96ColumnChunkReader struct {
	columnChunkReader
}

// Skip skips the next nvalues so that the next call to ReadBatch
// will start reading *after* the skipped values.
func (cr *Int96ColumnChunkReader) Skip(nvalues int64) (int64, error) {
	return cr.columnChunkReader.skipValues(nvalues,
		func(batch int64, buf []byte) (int64, error) {
			vals, _, err := cr.ReadBatch(batch,
				parquet.Int96Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf))
			return vals, err
		})
}

// ReadBatch reads batchSize values from the column.
//
// Returns error if values is not at least big enough to hold the number of values that will be read.
//
// defLvls and repLvls can be nil, or will be populated if not nil. If not nil, they must be
// at least large enough to hold the number of values that will be read.
//
// total is the number of rows that were read, valuesRead is the actual number of physical values
// that were read excluding nulls
func (cr *Int96ColumnChunkReader) ReadBatch(batchSize int64, values []parquet.Int96, defLvls, repLvls []int16) (total int64, valuesRead int, err error) {
	return cr.readBatch(batchSize, defLvls, repLvls, func(start, len int64) (int, error) {
		return cr.curDecoder.(encoding.Int96Decoder).Decode(values[start : start+len])
	})
}

// Float32ColumnChunkReader is the Typed Column chunk reader instance for reading
// Float32 column data.
type Float32ColumnChunkReader struct {
	columnChunkReader
}

// Skip skips the next nvalues so that the next call to ReadBatch
// will start reading *after* the skipped values.
func (cr *Float32ColumnChunkReader) Skip(nvalues int64) (int64, error) {
	return cr.columnChunkReader.skipValues(nvalues,
		func(batch int64, buf []byte) (int64, error) {
			vals, _, err := cr.ReadBatch(batch,
				arrow.Float32Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf))
			return vals, err
		})
}

// ReadBatch reads batchSize values from the column.
//
// Returns error if values is not at least big enough to hold the number of values that will be read.
//
// defLvls and repLvls can be nil, or will be populated if not nil. If not nil, they must be
// at least large enough to hold the number of values that will be read.
//
// total is the number of rows that were read, valuesRead is the actual number of physical values
// that were read excluding nulls
func (cr *Float32ColumnChunkReader) ReadBatch(batchSize int64, values []float32, defLvls, repLvls []int16) (total int64, valuesRead int, err error) {
	return cr.readBatch(batchSize, defLvls, repLvls, func(start, len int64) (int, error) {
		return cr.curDecoder.(encoding.Float32Decoder).Decode(values[start : start+len])
	})
}

// Float64ColumnChunkReader is the Typed Column chunk reader instance for reading
// Float64 column data.
type Float64ColumnChunkReader struct {
	columnChunkReader
}

// Skip skips the next nvalues so that the next call to ReadBatch
// will start reading *after* the skipped values.
func (cr *Float64ColumnChunkReader) Skip(nvalues int64) (int64, error) {
	return cr.columnChunkReader.skipValues(nvalues,
		func(batch int64, buf []byte) (int64, error) {
			vals, _, err := cr.ReadBatch(batch,
				arrow.Float64Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf))
			return vals, err
		})
}

// ReadBatch reads batchSize values from the column.
//
// Returns error if values is not at least big enough to hold the number of values that will be read.
//
// defLvls and repLvls can be nil, or will be populated if not nil. If not nil, they must be
// at least large enough to hold the number of values that will be read.
//
// total is the number of rows that were read, valuesRead is the actual number of physical values
// that were read excluding nulls
func (cr *Float64ColumnChunkReader) ReadBatch(batchSize int64, values []float64, defLvls, repLvls []int16) (total int64, valuesRead int, err error) {
	return cr.readBatch(batchSize, defLvls, repLvls, func(start, len int64) (int, error) {
		return cr.curDecoder.(encoding.Float64Decoder).Decode(values[start : start+len])
	})
}

// BooleanColumnChunkReader is the Typed Column chunk reader instance for reading
// Boolean column data.
type BooleanColumnChunkReader struct {
	columnChunkReader
}

// Skip skips the next nvalues so that the next call to ReadBatch
// will start reading *after* the skipped values.
func (cr *BooleanColumnChunkReader) Skip(nvalues int64) (int64, error) {
	return cr.columnChunkReader.skipValues(nvalues,
		func(batch int64, buf []byte) (int64, error) {
			vals, _, err := cr.ReadBatch(batch,
				*(*[]bool)(unsafe.Pointer(&buf)),
				nil,
				nil)
			return vals, err
		})
}

// ReadBatch reads batchSize values from the column.
//
// Returns error if values is not at least big enough to hold the number of values that will be read.
//
// defLvls and repLvls can be nil, or will be populated if not nil. If not nil, they must be
// at least large enough to hold the number of values that will be read.
//
// total is the number of rows that were read, valuesRead is the actual number of physical values
// that were read excluding nulls
func (cr *BooleanColumnChunkReader) ReadBatch(batchSize int64, values []bool, defLvls, repLvls []int16) (total int64, valuesRead int, err error) {
	return cr.readBatch(batchSize, defLvls, repLvls, func(start, len int64) (int, error) {
		return cr.curDecoder.(encoding.BooleanDecoder).Decode(values[start : start+len])
	})
}

// ByteArrayColumnChunkReader is the Typed Column chunk reader instance for reading
// ByteArray column data.
type ByteArrayColumnChunkReader struct {
	columnChunkReader
}

// Skip skips the next nvalues so that the next call to ReadBatch
// will start reading *after* the skipped values.
func (cr *ByteArrayColumnChunkReader) Skip(nvalues int64) (int64, error) {
	return cr.columnChunkReader.skipValues(nvalues,
		func(batch int64, buf []byte) (int64, error) {
			vals, _, err := cr.ReadBatch(batch,
				parquet.ByteArrayTraits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf))
			return vals, err
		})
}

// ReadBatch reads batchSize values from the column.
//
// Returns error if values is not at least big enough to hold the number of values that will be read.
//
// defLvls and repLvls can be nil, or will be populated if not nil. If not nil, they must be
// at least large enough to hold the number of values that will be read.
//
// total is the number of rows that were read, valuesRead is the actual number of physical values
// that were read excluding nulls
func (cr *ByteArrayColumnChunkReader) ReadBatch(batchSize int64, values []parquet.ByteArray, defLvls, repLvls []int16) (total int64, valuesRead int, err error) {
	return cr.readBatch(batchSize, defLvls, repLvls, func(start, len int64) (int, error) {
		return cr.curDecoder.(encoding.ByteArrayDecoder).Decode(values[start : start+len])
	})
}

// FixedLenByteArrayColumnChunkReader is the Typed Column chunk reader instance for reading
// FixedLenByteArray column data.
type FixedLenByteArrayColumnChunkReader struct {
	columnChunkReader
}

// Skip skips the next nvalues so that the next call to ReadBatch
// will start reading *after* the skipped values.
func (cr *FixedLenByteArrayColumnChunkReader) Skip(nvalues int64) (int64, error) {
	return cr.columnChunkReader.skipValues(nvalues,
		func(batch int64, buf []byte) (int64, error) {
			vals, _, err := cr.ReadBatch(batch,
				parquet.FixedLenByteArrayTraits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf),
				arrow.Int16Traits.CastFromBytes(buf))
			return vals, err
		})
}

// ReadBatch reads batchSize values from the column.
//
// Returns error if values is not at least big enough to hold the number of values that will be read.
//
// defLvls and repLvls can be nil, or will be populated if not nil. If not nil, they must be
// at least large enough to hold the number of values that will be read.
//
// total is the number of rows that were read, valuesRead is the actual number of physical values
// that were read excluding nulls
func (cr *FixedLenByteArrayColumnChunkReader) ReadBatch(batchSize int64, values []parquet.FixedLenByteArray, defLvls, repLvls []int16) (total int64, valuesRead int, err error) {
	return cr.readBatch(batchSize, defLvls, repLvls, func(start, len int64) (int, error) {
		return cr.curDecoder.(encoding.FixedLenByteArrayDecoder).Decode(values[start : start+len])
	})
}
