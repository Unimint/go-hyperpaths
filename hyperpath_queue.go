package hyperpaths

import "container/heap"

type pqEntry struct {
	link     *Link
	priority float32 // u_j + c_a (used for prioritization)
	index    int     // The index of the item in the heap.
}

type PriorityQueue []*pqEntry

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqEntry)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an pqEntry in the queue.
func (pq *PriorityQueue) update(item *pqEntry, priority float32) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
