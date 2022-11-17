package gobar

// Heap min heap
type Heap struct {
	data []int
	size int
}

func NewHeap(data []int) *Heap {
	h := &Heap{
		data: make([]int, len(data)+1),
		size: len(data),
	}
	for i := range data {
		h.data[i+1] = data[i]
	}
	s := h.size >> 1
	for s >= 1 {
		h.downAdjust(s)
		s--
	}
	return h
}

func (h *Heap) Add(item int) {
	h.data = append(h.data, item)
	h.upAdjust(len(h.data) - 1)
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) PopTop() int {
	if h.size == 0 {
		return -1
	}
	ans := h.data[1]
	h.data[1], h.data[h.size] = h.data[h.size], h.data[1]
	h.data = h.data[:h.size]
	h.size--
	h.downAdjust(1)
	return ans
}

func (h *Heap) upAdjust(s int) {
	data := h.data
	for s > 1 {
		f := s >> 1
		if data[f] > data[s] {
			data[f], data[s] = data[s], data[f]
			s = f
		} else {
			break
		}
	}
}

func (h *Heap) downAdjust(s int) {
	data := h.data
	for {
		l := s << 1
		r := s<<1 | 1
		if l >= len(data) {
			break
		} else if r >= len(data) {
			if data[l] < data[s] {
				data[s], data[l] = data[l], data[s]
				s = l
			} else {
				break
			}
		} else {
			if data[l] < data[r] {
				if data[l] < data[s] {
					data[s], data[l] = data[l], data[s]
					s = l
				} else {
					break
				}
			} else {
				if data[r] < data[s] {
					data[s], data[r] = data[r], data[s]
					s = r
				} else {
					break
				}
			}
		}
	}
}
