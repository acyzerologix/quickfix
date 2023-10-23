package quickfix

import "sync"

type BytesPool struct {
	pool sync.Pool
}

func NewBytesPool(bytesLength int) *BytesPool {
	bp := &BytesPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 0, bytesLength)
			},
		},
	}
	return bp
}

func (p *BytesPool) Get() []byte {
	return p.pool.Get().([]byte)
}

func (p *BytesPool) Put(slice []byte) {
	slice = slice[:0]
	p.pool.Put(slice)
	return
}
