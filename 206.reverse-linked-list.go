// @leet start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		next := current.Next // save next
		current.Next = prev // reverse pointer
		prev = current // move prev forward
		current = next //move current forward
	}
	return prev
}
// @leet end
