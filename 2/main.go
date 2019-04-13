package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {

	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}

	firstList, err := NewList(arr1)
	if err != nil {
		log.Fatal(err)
	}
	secondList, err := NewList(arr2)
	if err != nil {
		log.Fatal(err)
	}

	resultList := addTwoNumbers(firstList, secondList)

	fmt.Println(resultList)

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

// GetLastElement возвращает последний элемент списка
func (l *ListNode) GetLastElement() *ListNode {
	if l.Next == nil {
		return l
	}

	var p *ListNode
	for p = l.Next; p.Next != nil; {
		p = p.Next
	}

	return p
}

// ListNode - элемент списка
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var resultArray []int

	p1 := l1
	p2 := l2

	var remainder int
	for p1 != nil || p2 != nil {
		v := remainder
		if p1 != nil {
			v += p1.Val
			p1 = p1.Next
		}

		if p2 != nil {
			v += p2.Val
			p2 = p2.Next
		}

		resultArray = append(resultArray, v%10)
		remainder = v / 10
	}

	if remainder > 0 {
		resultArray = append(resultArray, remainder)
	}

	resultList, _ := NewList(resultArray)

	return resultList
}
