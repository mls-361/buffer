/*
------------------------------------------------------------------------------------------------------------------------
####### buffer ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package buffer

import (
	"strconv"
	"strings"
	"time"
)

type (
	// Buffer AFAIRE.
	Buffer struct {
		bs []byte
	}
)

// New AFAIRE.
func New(bufSize int) *Buffer {
	return &Buffer{bs: make([]byte, 0, bufSize)}
}

// AppendByte AFAIRE.
func (b *Buffer) AppendByte(value byte) {
	b.bs = append(b.bs, value)
}

// AppendString AFAIRE.
func (b *Buffer) AppendString(value string) {
	b.bs = append(b.bs, value...)
}

// AppendLeftJustifiedString AFAIRE.
func (b *Buffer) AppendLeftJustifiedString(s string, r byte, size int) {
	if len(s) < size {
		b.AppendString(s)
		b.AppendString(strings.Repeat(string(r), size-len(s)))
	} else {
		b.AppendString(s[:size])
	}
}

// AppendRightJustifiedString AFAIRE.
func (b *Buffer) AppendRightJustifiedString(s string, r byte, size int) {
	if len(s) < size {
		b.AppendString(strings.Repeat(string(r), size-len(s)))
		b.AppendString(s)
	} else {
		b.AppendString(s[:size])
	}
}

// AppendInt AFAIRE.
func (b *Buffer) AppendInt(value int64) {
	b.bs = strconv.AppendInt(b.bs, value, 10)
}

// AppendFloat AFAIRE.
func (b *Buffer) AppendFloat(value float64, bitSize int) {
	b.bs = strconv.AppendFloat(b.bs, value, 'f', -1, bitSize)
}

// AppendBool AFAIRE.
func (b *Buffer) AppendBool(value bool) {
	b.bs = strconv.AppendBool(b.bs, value)
}

// AppendTime AFAIRE.
func (b *Buffer) AppendTime(value time.Time, layout string) {
	b.bs = value.AppendFormat(b.bs, layout)
}

// Write AFAIRE.
func (b *Buffer) Write(bs []byte) (int, error) {
	b.bs = append(b.bs, bs...)
	return len(bs), nil
}

// Len AFAIRE.
func (b *Buffer) Len() int {
	return len(b.bs)
}

// Bytes AFAIRE.
func (b *Buffer) Bytes() []byte {
	return b.bs
}

// String AFAIRE.
func (b *Buffer) String() string {
	return string(b.bs)
}

// Reset AFAIRE.
func (b *Buffer) Reset() {
	b.bs = b.bs[:0]
}

/*
######################################################################################################## @(°_°)@ #######
*/
