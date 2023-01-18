package bitmap

type bitmap struct {
	elem []uint8
}

func NewBitmap(size int) *bitmap {
	byteSize := (size + 7) / 8
	return &bitmap{
		elem: make([]uint8, byteSize),
	}
}

func (t *bitmap) Set(idx int) {
	if idx > len(t.elem)*8 {
		return
	}

	byteIndex := idx / 8
	bitIndex := idx % 8
	t.elem[byteIndex] |= (1 << bitIndex)
}

func (t *bitmap) Unset(idx int) {
	if idx > len(t.elem)*8 {
		return
	}

	byteIndex := idx / 8
	bitIndex := uint8(idx % 8)
	t.elem[byteIndex] &^= (1 << bitIndex)
}

func (t *bitmap) IsSet(idx int) bool {
	if idx > len(t.elem)*8 {
		return false
	}

	byteIndex := idx / 8
	bitIndex := idx % 8
	return t.elem[byteIndex]&(1<<bitIndex) > 1
}
