package heap

type MinHeap struct {
	items []Heapable
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (m *MinHeap) Size() int {
	return len(m.items)
}

func (m *MinHeap) leaf(index int) bool {
	if index >= (m.Size()) && index <= m.Size() {
		return true
	}
	return false
}


func (m *MinHeap) parent(index int) int {
	return (index - 1) / 2
}


func (m *MinHeap) leftChild(index int) int {
	return 2 * index + 1
}

func (m *MinHeap) rightChild(index int) int {
	return 2 * index + 2
}


func (m *MinHeap) Insert(item Heapable) {
	m.items = append(m.items, item)
	m.upHeapify(m.Size() - 1)
}

func (m *MinHeap) swap(first, second int) {
	temp := m.items[first]
	m.items[first] = m.items[second]
	m.items[second] = temp
}

func (m *MinHeap) upHeapify(index int) {
	for m.items[index].Priority() < m.items[m.parent(index)].Priority() {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *MinHeap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	smallest := current
	leftChildIndex := m.leftChild(current)
	rightChildIndex := m.rightChild(current)
	// If current is smallest then return
	if leftChildIndex < m.Size() && m.items[leftChildIndex].Priority() < m.items[smallest].Priority() {
		smallest = leftChildIndex
	}
	if rightChildIndex < m.Size() && m.items[rightChildIndex].Priority() < m.items[smallest].Priority() {
		smallest = rightChildIndex
	}

	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest)
	}
}

func (m *MinHeap) buildMinHeap() {
	for index := (m.Size() / 2) - 1; index >= 0; index-- {
		m.downHeapify(index)
	}
}


func (m *MinHeap) Remove() Heapable {
	if m.Size() == 0 {
		return nil
	}
	top := m.items[0]
	m.items[0] = m.items[m.Size()-1]
	m.items = m.items[:m.Size()-1]
	m.downHeapify(0)
	return top
}


func (m *MinHeap) Contents() []Heapable {
	return m.items
}

func (m *MinHeap) Peek() Heapable {
	if m.Size() == 0 {
		return nil
	}
	return m.items[0]
}
