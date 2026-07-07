package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

// addTwoNumbers 两数相加（逆序链表）
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 虚拟头节点（不存数值，仅做哨兵）
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.value
			l1 = l1.next
		}
		if l2 != nil {
			sum += l2.value
			l2 = l2.next
		}
		carry = sum / 10
		// 新建节点保存个位
		current.next = &ListNode{value: sum % 10}
		current = current.next
	}
	return dummy.next // 返回虚拟头下一个，跳过无效哨兵
}

// 工具1：数组快速构建链表
func buildList(arr []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range arr {
		cur.next = &ListNode{value: v}
		cur = cur.next
	}
	return dummy.next
}

// 工具2：打印链表（输出数字）
func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.value)
		cur = cur.next
	}
	fmt.Println()
}

func main() {
	// 测试用例1：基础案例 l1=[2,4,3] l2=[5,6,4] 预期 [7,0,8]
	fmt.Println("测试用例1:")
	l1 := buildList([]int{2, 4, 3})
	l2 := buildList([]int{5, 6, 4})
	res1 := addTwoNumbers(l1, l2)
	printList(res1) // 输出：7 0 8

	// 测试用例2：有进位到最高位 l1=[9,9,9,9,9,9,9] l2=[9,9,9,9] 预期 [8,9,9,9,0,0,0,1]
	fmt.Println("测试用例2:")
	l3 := buildList([]int{9, 9, 9, 9, 9, 9, 9})
	l4 := buildList([]int{9, 9, 9, 9})
	res2 := addTwoNumbers(l3, l4)
	printList(res2) // 输出：8 9 9 9 0 0 0 1

	// 测试用例3：其中一个链表为空 l1=[0] l2=[0] 预期 [0]
	fmt.Println("测试用例3:")
	l5 := buildList([]int{0})
	l6 := buildList([]int{0})
	res3 := addTwoNumbers(l5, l6)
	printList(res3) // 输出：0

	
}