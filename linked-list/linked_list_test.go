package linked_list

import "testing"

var l *LinkedList

func init() {
    n5 := &ListNode{value: 5}
    n4 := &ListNode{value: 4, next: n5}
    n3 := &ListNode{value: 3, next: n4}
    n2 := &ListNode{value: 2, next: n3}
    n1 := &ListNode{value: 1, next: n2}
    l = &LinkedList{head: &ListNode{next: n1}}
}

func TestLinkedList_Reverse(t *testing.T) {
    l.Print()
    l.Reverse()
    l.Print()
}

func TestLinkedList_HasCycle(t *testing.T) {
    t.Log(l.HasCycle())
    l.head.next.next.next.next.next.next = l.head.next.next.next
    t.Log(l.HasCycle())
}
