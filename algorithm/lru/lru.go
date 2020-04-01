package main

type Node struct {
	key,val int
	next,prev *Node
}

type LRUCache struct {
	cap int // 缓存容量
	cache map[int]*Node//判断数据是否存在
	head *Node //头指针
	tail *Node // 尾指针
}
func NewLRUCache(capacity int)*LRUCache{
	l :=&LRUCache{
		cap:  capacity,
		cache: make(map[int]*Node),
		head:  &Node{},
		tail:  &Node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}
// 双链表在头部插入数据
func (l *LRUCache)insert(node *Node){
	tmp := l.head.next
	l.head.next = node
	node.next = tmp
	tmp.prev = node
	node.prev = l.head
}
// 双链表插入数据
func (l *LRUCache)remove(node *Node){
	n := node.next// 获取后节点
	p := node.prev // 获取前节点
	p.next = n
	n.prev = p
	node.next = nil
	node.prev = nil
}

func(l *LRUCache)Get(key int)int{
	if n,ok:=l.cache[key];ok{
		l.remove(n)
		l.insert(n)
		return n.val
	}
	return -1
}
func(l *LRUCache)Put(key int,value int){
	if n,ok:=l.cache[key];ok{
		n.val =value
		l.remove(n)
		l.insert(n)
	}else {
		if len(l.cache)==l.cap {//需要判断当前容器是否已经满了
			deleteNode := l.tail.prev
			delete(l.cache,deleteNode.key)
			l.remove(deleteNode)
		}
		n :=&Node{
			key:  key,
			val:  value,
		}
		l.insert(n)
		l.cache[key]= n
	}
}

func main() {
	
}
