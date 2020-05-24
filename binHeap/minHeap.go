package binHeap

import "errors"

type Heap struct {
	heapData []int
	size     int
	maxSize  int
}

func (h Heap) Parent(x int) int {
	return (x - 1) / 2
}

func (h Heap) LeftChild(x int) int {
	return 2*x + 1
}

func (h Heap) RightChild(x int) int {
	return 2*x + 2
}

func (h Heap) Leaf(x int) bool {
	if x >= (h.size)/2 && x <= h.size {
		return true
	}

	return false
}

func NewMinHeap(x int) *Heap {
	minHeap := &Heap{
		heapData: []int{},
		size:     0,
		maxSize:  x,
	}

	return minHeap
}

func (h *Heap) Swap(x, y int) {
	temp := h.heapData[x]
	h.heapData[x] = h.heapData[y]
	h.heapData[y] = temp
}

func (h *Heap) Insert(item int) error {
	if h.size > h.maxSize {
		return errors.New("Heap is full")
	}

	h.heapData = append(h.heapData, item)
	h.size++
	h.upHeapify(h.size - 1)
	return nil
}

func (h *Heap) upHeapify(x int) {
	for h.heapData[x] < h.heapData[h.Parent(x)] { //For MaxHeap ise greater comparison
		h.Swap(x, h.Parent(x))
		x = h.Parent(x)
	}
}

func (h *Heap) DownHeapify(x int) {
	if h.Leaf(x) {
		return
	}

	smallest := x
	left := h.LeftChild(x)
	right := h.RightChild(x)

	if left < h.size && h.heapData[left] < h.heapData[smallest] {
		smallest = left
	}

	if right < h.size && h.heapData[right] < h.heapData[smallest] {
		smallest = right
	}

	if smallest != x {
		h.Swap(x, smallest)
		h.DownHeapify(smallest)
	}

	return
}

func (h *Heap) BuildHeap() {
	for index := ((h.size / 2) - 1); index >= 0; index-- {
		h.DownHeapify(index)
	}
}

func (h *Heap) Remove() int {
	top := h.heapData[0]
	h.heapData[0] = h.heapData[h.size-1]
	h.heapData = h.heapData[:(h.size)-1]
	h.size--
	h.DownHeapify(0)
	return top
}
