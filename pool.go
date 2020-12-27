/*
------------------------------------------------------------------------------------------------------------------------
####### buffer ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package buffer

import "sync"

type (
	// Pool AFAIRE.
	Pool struct {
		sp *sync.Pool
	}
)

// NewPool AFAIRE.
func NewPool(bufSize int) *Pool {
	return &Pool{
		sp: &sync.Pool{
			New: func() interface{} {
				return New(bufSize)
			},
		},
	}
}

// Get AFAIRE.
func (p *Pool) Get() *Buffer {
	buf := p.sp.Get().(*Buffer)
	return buf
}

// Put AFAIRE.
func (p *Pool) Put(buf *Buffer) {
	buf.Reset()
	p.sp.Put(buf)
}

/*
######################################################################################################## @(°_°)@ #######
*/
