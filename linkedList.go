package main

type SingleListNode struct {
	next *SingleListNode
	Value interface{}
}

type SingleList struct {
	first *SingleListNode
	last *SingleListNode
}

func (this *SingleList) Push(val interface{}) {

	if this.first == nil {

		this.first = &SingleListNode{nil, val}
		this.last = this.first

	} else {

		old := this.first
		this.first = &SingleListNode{nil, val}
		this.first.next = old

	}

}

func (this *SingleList) PushBack(val interface{}) {

	if this.first == nil {
		this.first = &SingleListNode{nil, val}
		this.last = this.first
	} else {
		this.last.next = &SingleListNode{nil, val}
		this.last = this.last.next
	}

}

func (this *SingleList) Length() int {

	current := this.first
	count := 0

	for current != nil {
		count++
		current = current.next
	}

	return count

}

func (this *SingleList) ToString() string {

	current := this.first
	out := ""

	for current != nil {
		out = out + "["+current.Value.(string)+"] "
		current = current.next
	}

	return out

}

func (this *SingleList) Get(index int) interface{} {

	counter := 0

	for current := this.first; ; current = current.next {

		if counter == index {
			return current.Value
		} else if current.next == nil {
			panic("Attempted to get an index greater than list length.")
		} else {
			counter++
		}

	}

}

func (this *SingleList) Delete(index int) {

	counter := 0
	current := this.first
	var prev *SingleListNode

	for  {

		if counter == index && prev == nil {
			// deleted first elem
			this.first = nil
			this.last = nil
			return
		} else if counter == index && current.next == nil  {
			// deleted last elem
			prev.next = current.next
			this.last = prev
			return
		} else if counter == index {
			// deleted any middle elem
			prev.next = current.next
			return
		}

		counter++
		prev = current
		current = current.next

	}


}

func NewLinkedList() *SingleList {
	return &SingleList{}
}
