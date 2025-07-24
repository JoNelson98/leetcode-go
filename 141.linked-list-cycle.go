// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// fast and slow pointer method
	//build fast and slow pointers
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		//move pointers
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
// @leet end
