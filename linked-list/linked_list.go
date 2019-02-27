package linked_list

import "fmt"

//单链表节点
type ListNode struct {
    next  *ListNode
    value interface{}
}

//单链表
type LinkedList struct {
    head *ListNode
}

//打印链表
func (this *LinkedList) Print() {
    cur := this.head.next
    format := ""
    for nil != cur {
        format += fmt.Sprintf("%+v", cur.value)
        cur = cur.next
        if nil != cur {
            format += "->"
        }
    }
    fmt.Println(format)
}

//链表反转
//把后继节点往第一个节点放，最后把头结点的next和第一个节点连接起来
func (this *LinkedList) Reverse() {
    if nil == this.head || nil == this.head.next || nil == this.head.next.next {
        return
    }

    var pre *ListNode

    cur := this.head.next
    for nil != cur {
        cur.next, pre, cur = pre, cur, cur.next
    }

    this.head.next = pre
}

//判断单链表是否有环
/*
定义两个指针，同时从链表的头节点出发，一个指针一次走一步，另一个指针一次走两步。
如果走得快的指针追上了走得慢的指针，那么链表就是环形链表；
如果走得快的指针走到了链表的末尾（next指向nil）都没有追上第一个指针，那么链表就不是环形链表。
 */
func (this *LinkedList) HasCycle() bool {
    if nil != this.head {
        fast := this.head
        slow := this.head

        for fast != nil && fast.next != nil {
            slow = slow.next
            fast = fast.next.next

            if fast == slow {
                return true
            }
        }
    }

    return false
}

//判断单链表中是否有环，若有，找出环的入口节点
func (this *LinkedList) CycleNode() {

}

//有序链表合并
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
    //如果l1或l2为空链表，则返回
    if nil == l1 || nil == l1.head || nil == l1.head.next {
        return l2
    }

    if nil == l2 || nil == l2.head || nil == l2.head.next {
        return l1
    }

    l := &LinkedList{head: &ListNode{}}
}

//删除倒数第N个节点
func (this *LinkedList) DeleteBottomN(n int) {

}

//获取中间节点
func (this *LinkedList) FindMiddleNode() *ListNode {

}
