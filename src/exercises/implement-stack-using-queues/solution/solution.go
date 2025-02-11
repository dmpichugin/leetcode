package solution

type MyStack struct {
	head *Element
}

type Element struct {
	Value int
	Next  *Element
}

func Constructor() MyStack {

	return MyStack{
		head: nil,
	}
}

func (this *MyStack) Push(x int) {
	if this.head == nil {
		this.head = &Element{Value: x, Next: nil}
	} else {
		e := &Element{Value: x, Next: this.head}
		this.head = e
	}
}

func (this *MyStack) Pop() int {
	if this.head == nil {
		return 0
	}
	val := this.head.Value
	this.head = this.head.Next
	return val
}

func (this *MyStack) Top() int {
	if this.head == nil {
		return 0
	}
	return this.head.Value
}

func (this *MyStack) Empty() bool {
	return this.head == nil
}
