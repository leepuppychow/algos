package heap

type MinHeap []int

func (h *MinHeap) Min() int {
	return (*h)[0]
}

func (h *MinHeap) Length() int {
	return len(*h)
}

func (h *MinHeap) Insert(currentValue int) {
	// 1. Insert number at last index of MinHeap
	// 2. If parent is larger, then swap current with parent  --> Check again
	// 3. Else if parent is smaller OR there is no parent then leave as is --> STOP

	*h = append(*h, currentValue)
	currentIndex := len(*h) - 1

	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2
		parentValue := (*h)[parentIndex]

		if currentValue < parentValue {
			(*h)[currentIndex], (*h)[parentIndex] = (*h)[parentIndex], (*h)[currentIndex]
			currentIndex = parentIndex
		} else {
			currentIndex = 0
		}
	}
}

func (h *MinHeap) PopMin() int {
	// Save the min value, swap start and end, then remove the last value (min)
	min := (*h)[0]
	if h.Length() == 1 { // check in case MinHeap only has 1 item
		*h = (*h)[0:0]
		return min
	}
	
	endIndex := len(*h) - 1
	(*h)[0], (*h)[endIndex] = (*h)[endIndex], (*h)[0]
	*h = (*h)[0:endIndex]
	endIndex--
	currentIndex := 0
	currentValue := (*h)[currentIndex] // currentValue stays the same, only currentIndex changes as it bubbles down

	for true {
		leftChildIndex := 2*currentIndex + 1
		rightChildIndex := 2*currentIndex + 2

		if leftChildIndex > endIndex && rightChildIndex > endIndex { // here no more children
			break
		} else if rightChildIndex > endIndex { // only has left child
			if currentValue > (*h)[leftChildIndex] {
				(*h)[leftChildIndex], (*h)[currentIndex] = (*h)[currentIndex], (*h)[leftChildIndex]
				currentIndex = leftChildIndex
			} else {
				currentIndex = endIndex
			}
		} else { // has both children, swap with the smallest
			if currentValue > (*h)[leftChildIndex] && (*h)[leftChildIndex] < (*h)[rightChildIndex] {
				(*h)[leftChildIndex], (*h)[currentIndex] = (*h)[currentIndex], (*h)[leftChildIndex]
				currentIndex = leftChildIndex
			} else if currentValue > (*h)[rightChildIndex] && (*h)[rightChildIndex] < (*h)[leftChildIndex] {
				(*h)[rightChildIndex], (*h)[currentIndex] = (*h)[currentIndex], (*h)[rightChildIndex]
				currentIndex = leftChildIndex
			} else {
				currentIndex = endIndex
			}
		}
	}

	return min
}
