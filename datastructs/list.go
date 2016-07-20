package main

type IntList struct {
	v    int
	next *IntList
}

// Insert will place the value into the list at the
// current location and will return the head of the
// list, which is now the node for the new value.
func (n *IntList) Insert(value int) *IntList {
	node := &IntList{value, n}
	return node
}

// Remove returns the next node of the list, effectively
// removing the current node from the list.
func (n *IntList) Remove() *IntList {
	return n.next
}

// Unique will walk the list and remove any duplicate
// values that it find.
func (n *IntList) Unique() *IntList {
	log := make(map[int]int, 1)
	log[n.v] = n.v

	node := n
	for node.next != nil {
		if _, exists := log[node.next.v]; !exists {
			// Mark that we've seen the value once and
			log[node.next.v] = node.next.v
		} else {
			// Remove the next node from the list, since the
			// value has already been seen.
			node.next = node.next.next
		}

		node = node.next
	}

	return n
}
