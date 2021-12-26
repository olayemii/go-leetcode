/*
21. Merge Two Sorted Lists

https://leetcode.com/problems/merge-two-sorted-lists/

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.



Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ans := new(ListNode)
	head := ans

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			ans.Next = list1
			list1 = list1.Next
		} else {
			ans.Next = list2
			list2 = list2.Next
		}
		ans = ans.Next
	}

	if list1 != nil {
		ans.Next = list1
	}
	if list2 != nil {
		ans.Next = list2
	}
	return head.Next
}

// Space Complexity: O(n)
// Time Complexity: O(n)
