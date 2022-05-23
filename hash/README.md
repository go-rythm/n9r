## 哈希表的原理

### 摘要

* HashMap和HashSet的联系和区别
* 哈希表Hash Table的基本原理
* 什么是哈希函数Hash Function
* 如何解决冲突Collision
    * 开散列Open Hash vs 闭散列Closed Hash
* 哈希表扩容问题

### 哈希表的基本原理

时间复杂度：`O(size of key)` O(L)

 ![hash](https://gitee.com/luxcgo/imgs4md/raw/master/img/20220524001404.jpeg)

### 开哈希与闭哈希

#### [开哈希](https://www.cs.usfca.edu/~galles/visualization/OpenHash.html)

每个位置存储一个链表

#### [闭哈希](https://www.cs.usfca.edu/~galles/visualization/ClosedHash.html)

每个位置存储一个 <key, value> pair

闭哈希在删除元素注意需要把位置置为`deleted`状态，防止查找时找不到因冲突而放到后面的元素

![closedhash](https://gitee.com/luxcgo/imgs4md/raw/master/img/20220524001951.jpeg)

### [128](https://www.lintcode.com/problem/hash-function) 哈希函数

取模过程要使用同余定理：` (a b ) % MOD = ((a % MOD) (b % MOD)) % MOD`

```go
func HashCode(key string, hASH_SIZE int) int {
	var ans int
	for _, r := range key {
		ans = (ans*33 + int(r)) % hASH_SIZE
	}
	return ans
}
```

### [129](https://www.lintcode.com/problem/rehashing/) 重哈希

```go
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
```
