package heap

type Heap struct {
	list []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) left(i int) int {
	return 2*i + 1
}

func (h *Heap) right(i int) int {
	return 2*i + 2
}

func (h *Heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) heapify(i int) {
	n := len(h.list)
	if i >= n {
		return
	}
	l, r := h.left(i), h.right(i)
	smallest := i
	if l < n && h.list[l] < h.list[smallest] {
		smallest = l
	}
	if r < n && h.list[r] < h.list[smallest] {
		smallest = r
	}
	if smallest != i {
		h.list[i], h.list[smallest] = h.list[smallest], h.list[i]
		h.heapify(smallest)
	}
}

func (h *Heap) Len() int {
	return len(h.list)
}

func (h *Heap) Peek() (int, bool) {
	if len(h.list) == 0 {
		return 0, false
	}
	return h.list[0], true
}

func (h *Heap) Add(num int) {
	h.list = append(h.list, num)
	i := len(h.list) - 1
	for i > 0 && h.list[h.parent(i)] > h.list[i] {
		p := h.parent(i)
		h.list[p], h.list[i] = h.list[i], h.list[p]
		i = p
	}
}

func (h *Heap) Pop() (int, bool) {
	n := len(h.list)
	if n == 0 {
		return 0, false
	}
	var res int
	if n == 1 {
		res = h.list[0]
		h.list = h.list[:n-1]
		return res, true
	}
	res = h.list[0]
	h.list[0] = h.list[n-1]
	h.list = h.list[:n-1]
	h.heapify(0)
	return res, true
}
