// using the map method

func hasCycle(head *ListNode) bool {
    seenNodes := make(map[*ListNode]bool)
    current := head

    for current != nil {
        if seenNodes[current] {
            return true // cycle detected
        }
        seenNodes[current] = true
        current = current.Next
    }

    return false // no cycle
}

