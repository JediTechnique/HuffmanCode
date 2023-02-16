package HuffmanCode

import (
	"container/heap"
)

type Node struct {
	Symbol      string
	Freq        int
	Left, Right *Node
}

// Min Heap implementation

type minHeap []*Node

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Freq < h[j].Freq }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(*Node))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//

func GetHuffmanTree(mp map[string]int) *Node {

	pq := &minHeap{}
	heap.Init(pq)

	for i, j := range mp {

		newNode := new(Node)
		newNode.Freq = j
		newNode.Symbol = i
		newNode.Left = nil
		newNode.Right = nil

		heap.Push(pq, newNode)

	}

	for len(*pq) > 1 {

		leftNode := (*pq)[0]
		heap.Pop(pq)

		rightNode := (*pq)[0]
		heap.Pop(pq)

		newFreq := leftNode.Freq + rightNode.Freq

		newNode := new(Node)
		newNode.Freq = newFreq
		newNode.Symbol = ""
		newNode.Left = leftNode
		newNode.Right = rightNode

		heap.Push(pq, newNode)
	}

	return (*pq)[0]

}

func dfs(curr *Node, code []byte, codes map[string]string) {

	if curr.Symbol != "" {
		codes[curr.Symbol] = string(code)
	}

	if curr.Left != nil {
		code = append(code, '0')
		dfs(curr.Left, code, codes)
		code = code[:len(code)-1]
	}

	if curr.Right != nil {
		code = append(code, '1')
		dfs(curr.Right, code, codes)
		code = code[:len(code)-1]
	}
}

func GetCodes(mp map[string]int) map[string]string {

	codes := make(map[string]string)

	root := GetHuffmanTree(mp)

	code := []byte{}

	dfs(root, code, codes)

	return codes

}
