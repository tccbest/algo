// 如何判断一个字符串是否是回文字符串
// 单链表实现
// 同类题
// 1、判断一条单向链表是不是“回文”
// 分析：对于单链表结构，可以用两个指针从两端或者中间遍历并判断对应字符是否相等。
// 但这里的关键就是如何朝两个方向遍历。
// 由于单链表是单向的，所以要向两个方向遍历的话，可以采取经典的快慢指针的方法，即先位到链表的中间位置，再将链表的后半逆置，最后用两个指针同时从链表头部和中间开始同时遍历并比较即可。
package _6_linkedlist

func IsPalindrome(head *ListNode) bool {
	var slow *ListNode = head
	var fast *ListNode = head
	var prev *ListNode = nil
	var temp *ListNode = nil

	if (head == nil || head.next == nil) {
		return true
	}

	for (fast != nil && fast.next != nil) {
		fast = fast.next.next
		temp = slow.next
		slow.next = prev
		prev = slow
		slow = temp
	}

	if (fast != slow) {
		slow = slow.next
	}

	var l1 *ListNode = prev
	var l2 *ListNode = slow

	for (l1 != nil && l2 != nil && l1.value == l2.value) {
		l1 = l1.next
		l2 = l2.next
	}

	return l1 == nil && l2 == nil
}
