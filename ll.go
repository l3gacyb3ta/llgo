package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	prev  *Node
	value int
}

type List struct {
	first  *Node
	last   *Node
	length int
}

// get the node at a location
func (list List) get(index int) *Node {
	/* if list.length < index-1 {
		return &Node{nil, nil, 0}
	}

	current_node := list.first

	for index > 0 {
		if current_node.next != nil {
			current_node = current_node.next
		} else {
			return &Node{nil, nil, 0}
		}
		index = index - 1
	}

	return current_node */

	return list.first
}


// Append a node to the end of the list
func (list List) push(value int) {
	old_last := *list.last
	
	list.last = &Node{list.last, nil, value}
	old_last.next = list.last
	list.last.prev = &old_last

	old_len := list.length +1
	list.length = old_len
}

// Remove a node at a given index
func (list List) remove(index int) {
	if index != list.length-1 && index != 0 {
		// Basic middle node logic
		node := list.get(index)
		node.prev.next = list.get(index).next

	} else if index == list.length-1 {
		new_last := list.get(index - 1)
		new_last.next = nil
		*list.last = *new_last


	} else {
		new_first := list.get(0)
		*list.first = *new_first
	}
}

// print every node in the list
func (list List) print() {
	current_node := list.first

	for current_node != nil {
		fmt.Println(current_node.value)

		current_node = current_node.next
	}
}

// Create a new list
func create(value int) List {
	f_node := &Node{nil, nil, value}
	new_list := List{f_node, f_node, 1}

	return new_list
}

func main() {
	list := create(2)
	list.push(1)
	//list.push(2)
	//list.push(3)

	fmt.Println(list.first)

	//list.print()

	//fmt.Println(list.get(1))

	//list.remove(1)

	//list.print()

	//fmt.Println(list.get(1))
}
