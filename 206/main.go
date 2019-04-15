package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	d := []int{1, 2, 3}

	list, _ := NewList(d)

	fmt.Println(list)

	fmt.Println(reverseList(list))

}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p := head

	if head.Next == nil {
		return head
	}

	var reversed *ListNode

	p = head
	for p != nil {
		newElement := &ListNode{
			Val:  p.Val,
			Next: reversed,
		}

		reversed = newElement
		p = p.Next
	}

	return reversed
}

func (l *ListNode) String() string {
	builder := &strings.Builder{}

	_, err := builder.WriteString("[ ")
	if err != nil {
		log.Println(err)
	}

	for p := l; p != nil; p = p.Next {
		StringValue := fmt.Sprintf("%d ", p.Val)
		_, err = builder.WriteString(StringValue)
		if err != nil {
			log.Println(err)
		}
	}
	_, err = builder.WriteRune(']')
	if err != nil {
		log.Println(err)
	}

	return builder.String()
}

// Push помещает в список элемент
func (l *ListNode) Push(value int) {
	var p *ListNode

	for p = l.Next; p.Next != nil; p = p.Next {
	}

	p.Next = &ListNode{
		Val:  value,
		Next: nil,
	}
}

// NewList создает список из слайса
func NewList(values []int) (*ListNode, error) {

	if len(values) == 0 {
		return nil, fmt.Errorf("array is empty")
	}

	head := &ListNode{Val: values[0]}
	if len(values) == 1 {
		return head, nil
	}

	var newElements *ListNode
	for i := len(values) - 1; i >= 1; i-- {
		newElement := &ListNode{
			Val:  values[i],
			Next: newElements,
		}

		newElements = newElement
	}

	head.Next = newElements

	return head, nil
}

// ListNode - элемент списка
type ListNode struct {
	Val  int
	Next *ListNode
}
