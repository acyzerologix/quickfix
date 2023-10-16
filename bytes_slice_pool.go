package quickfix

type BytesPool struct {
	pool chan []byte
}

func NewBytesPool(size int, bytesLength int) *BytesPool {
	bp := &BytesPool{
		pool: make(chan []byte, size),
	}
	for i := 0; i < size; i++ {
		bp.pool <- make([]byte, 0, bytesLength)
	}
	return bp
}

func (p *BytesPool) Get() []byte {
	slice := <-p.pool
	slice = slice[:0]
	return slice
}

func (p *BytesPool) Put(slice []byte) {
	p.pool <- slice
	return
}
