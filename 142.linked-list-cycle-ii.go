// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	// loop 
		for fast != nil && fast.Next != nil {
				// move both pointers
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// collision: cycle detected
			// phase 2 goes here
			ptr1 := head
			ptr2 := slow
			for ptr1 != ptr2 {
				ptr1 = ptr1.Next
				ptr2 = ptr2.Next
			}
			return ptr1
		}

	}
	return nil // no cycle
}
// @leet end
