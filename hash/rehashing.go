package hash

// https://www.lintcode.com/problem/rehashing/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @param hashTable: A list of The first node of linked list
 * @return: A list of The first node of linked list which have twice size
 */
func Rehashing(hashTable []*ListNode) []*ListNode {
	newTable := make([]*ListNode, len(hashTable)*2)
	for _, node := range hashTable {
		p := node
		for p != nil {
			addNode(newTable, p.Val)
			p = p.Next
		}
	}
	return newTable
}

func addListNode(node *ListNode, num int) {
	if node.Next != nil {
		addListNode(node.Next, num)
	} else {
		node.Next = &ListNode{Val: num}
	}
}

func addNode(newTable []*ListNode, num int) {
	idx := num % len(newTable)
	if newTable[idx] == nil {
		newTable[idx] = &ListNode{Val: num}
	} else {
		addListNode(newTable[idx], num)
	}
}
