package list

//定义链表的结构
type listNode struct {
	Val int
	Next *listNode
}

//定义链表
type MyLinkedList struct {
	Header *listNode
	Tail *listNode
	Lens int
}


//初始化
func Constructor() MyLinkedList {
	return MyLinkedList{
		Header:nil,
		Tail:nil,
		Lens:0,
	}
}


//获取链表中第 index 个节点的值。如果索引无效，则返回-1。
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Lens{
		return -1
	}
	if index == 0{
		return this.Header.Val
	}
	cur := this.Header
	for cur.Next != nil{
		index--
		cur = cur.Next
		if index == 0{
			return cur.Val
		}
	}
	return -1
}


//在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
func (this *MyLinkedList) AddAtHead(val int)  {
	this.Header = &listNode{
		Val:val,
		Next:this.Header,
	}
	if this.Lens == 0{
		this.Tail = this.Header
	}
	this.Lens++
}


//将值为 val 的节点追加到链表的最后一个元素。
func (this *MyLinkedList) AddAtTail(val int)  {
	if this.Lens == 0{
		this.AddAtHead(val)
		return
	}
	this.Tail.Next = &listNode{
		Val:val,
		Next:nil,
	}
	this.Tail = this.Tail.Next
	this.Lens++
}


//在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index <= 0{
		this.AddAtHead(val)
		return
	}
	if index > this.Lens{
		return
	}
	if index == this.Lens{
		this.AddAtTail(val)
		return
	}
	cur := this.Header
	for cur.Next != nil{
		index--
		if index == 0{
			newNode := &listNode{
				Val:val,
				Next:cur.Next,
			}
			cur.Next = newNode
			this.Lens++
			return
		}
		cur = cur.Next
	}
}


//如果索引 index 有效，则删除链表中的第 index 个节点。
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.Lens{
		return
	}
	if index == 0{
		this.Header = this.Header.Next
		this.Lens--
		return
	}
	cur := this.Header
	for cur.Next != nil{
		index--
		if index == 0{
			if cur.Next.Next == nil{
				cur.Next = nil
				this.Tail = cur
			}else {
				cur.Next = cur.Next.Next
			}
			this.Lens--
			return
		}
		cur = cur.Next
	}
}