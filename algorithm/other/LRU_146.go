package main

import "fmt"

/**
LRU（least recently used），最近最少使用算法，也是redis一种近似算法；
思想：最近使用的数据会在未来一段时期内仍然被使用，已经很久没有使用的数据大概率在未来很长一段时间仍然不会被使用
*/
func main() {
	constructor := Constructor(3)
	constructor.put(1, 1)
	constructor.put(2, 2)
	constructor.put(3, 3)
	fmt.Println("get 1: ", constructor.get(1))
	fmt.Println("get 2: ", constructor.get(2))
	fmt.Println("get 3: ", constructor.get(3))
	constructor.put(4, 4)
	fmt.Println("get 4: ", constructor.get(4))
	fmt.Println("get 1: ", constructor.get(1))
}

type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

type LRUCache struct {
	container  map[int]*LinkNode
	capacity   int
	head, tail *LinkNode
}

// Constructor 构造器
func Constructor(capacity int) *LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return &LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

/**
移除一个节点
*/
func (this *LRUCache) removeNode(node *LinkNode) {
	// 抽离当前节点
	node.pre.next = node.next
	node.next.pre = node.pre
}

/**
往头部添加一个节点
*/
func (this *LRUCache) addToHead(node *LinkNode) {
	// 先把头节点拿出来
	head := this.head

	// 移动至头节点,这里有个坑，这里的必须先将head的后一个节点和要加在头部的节点先关联起来，
	// 如果是先将该节点和头节点关联起来，那么 head的next就会一直是 当前node节点的值
	node.next = head.next
	head.next.pre = node

	node.pre = head
	head.next = node

}

// 将当前节点移动至头节点
func (this *LRUCache) toHead(current *LinkNode) {
	this.removeNode(current)
	this.addToHead(current)
}

/**
get 方法，这里有两种情况：
	1. 如果当前数据不存在，则直接返回 -1
	2. 当前数据存在，则先把该数据放到缓存的头节点，然后在返回当前节点对应的value
*/
func (this *LRUCache) get(key int) int {
	container := this.container
	if node, exist := container[key]; exist {
		this.toHead(node)
		return node.val
	} else {
		return -1
	}
}

/**
put 方法主要有三种情况需要考虑：
	1. 数据已经存在的，这里其实只是做更新的操作，然后将当前节点移动到头部即可
	2. 数据不存在，但是容器已经满了的情况下，需要删除最后一个节点，再把新的节点放到头节点
	3. 容器未满的情况下，只要把新节点放到头部位置即可
*/
func (this *LRUCache) put(key, val int) {
	tail := this.tail
	container := this.container
	// 当该键已经存在，则更新值，然后将其移动到头节点
	if node, exist := container[key]; exist {
		node.val = val
		this.toHead(node)
	} else {
		// 初始化当前新增节点
		newNode := &LinkNode{key, val, nil, nil}
		// 如果容器已经满了，这个时候需要删除最后一个尾节点
		if len(container) == this.capacity {
			delete(container, tail.pre.key)
			this.removeNode(tail.pre)
		}
		// 把节点添加到头部
		this.addToHead(newNode)
		// 然后记录在缓存中的位置
		container[key] = newNode
	}
}
