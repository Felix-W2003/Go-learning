package main

import "fmt"

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。



示例 1：

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]


提示：

两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (list3 *ListNode) {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next
}

//辅助函数：数组转链表，方便测试
func sliceToList(arr []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range arr {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	fmt.Println(dummy)
	return dummy.Next
}

//辅助函数：链表打印输出
func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d", cur.Val)
		cur = cur.Next
	}

	fmt.Println()
}

func main() {
	// 示例1
	l1 := sliceToList([]int{1, 2, 3, 4})
	l2 := sliceToList([]int{1, 3, 4})
	res1 := mergeTwoLists(l1, l2)
	fmt.Print("示例1输出：")
	printList(res1) // 1 1 2 3 4 4

	// 示例2
	l3 := sliceToList([]int{})
	l4 := sliceToList([]int{})
	res2 := mergeTwoLists(l3, l4)
	fmt.Print("示例2输出：")
	printList(res2) // 空

	// 示例3
	l5 := sliceToList([]int{})
	l6 := sliceToList([]int{0})
	res3 := mergeTwoLists(l5, l6)
	fmt.Print("示例3输出：")
	printList(res3) // 0
}
