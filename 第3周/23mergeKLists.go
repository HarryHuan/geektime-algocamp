package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 法1：按照作业要求，使用分治。代码如下，结果OOM，额，请助教评论中给予提示。
func mergeKLists1(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	former := mergeKLists1(lists[:len(lists)/2])
	latter := mergeKLists1(lists[len(lists)/2:])
	return mergeTowLists(former, latter)
}

// 法2：直接，一个一个拼接上
func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	var result *ListNode
	for _, list := range lists {
		result = mergeTowLists(result, list)
	}
	return result
}

func mergeTowLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else {
			cur.Next = l2
			break
		}
	}
	return head.Next
}
