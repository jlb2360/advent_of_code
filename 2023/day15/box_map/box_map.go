package boxmap

import "fmt"

type Lense struct {
	Label string
	Power int
}

type LenseNode struct {
	info Lense
	next *LenseNode
}

type LenseList struct {
	head *LenseNode
	tail *LenseNode
}

type BoxNode struct {
	key  int
	info LenseList
	next *BoxNode
}

type BoxList struct {
	head *BoxNode
	tail *BoxNode
}

// Implement the LenseList methods Insert(), Change(), Delete(), and PerformEquals() below

func (ll *LenseList) Insert(info Lense) {
	// Insert a new node with the specified info at the end of the list
	if ll.head == nil {
		ll.head = &LenseNode{info: info}
		ll.tail = ll.head
	} else {
		ll.tail.next = &LenseNode{info: info}
		ll.tail = ll.tail.next
	}
}

func (ll *LenseList) Change(key int, info Lense) {
	// Change the info of the node with the specified key
	current := ll.head
	for current != nil {
		if current.info.Label == info.Label {
			current.info.Power = info.Power
			break
		}
		current = current.next
	}
}

func (ll *LenseList) PerformEquals(info Lense) {
	// This has two interpretations:
	// 1. Insert the Lense into the list with the specified key at the end
	// 2. Change the Lense in the list with the specified key with this Lense

	// Interpretation 1
	current := ll.head
	for current != nil {
		if current.info.Label == info.Label {
			current.info.Power = info.Power
			return
		}
		current = current.next
	}

	// Interpretation 2
	ll.Insert(info)
}

func (ll *LenseList) Delete(key string) {
	// Delete the node with the specified key
	current := ll.head
	if current == nil {
		return // Empty list
	}
	if current.info.Label == key {
		ll.head = current.next
		return
	}
	for current.next != nil {
		if current.next.info.Label == key {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (ll *LenseList) PrintList() {
	// Print the list
	current := ll.head
	for current != nil {
		fmt.Printf("%s:%d \n", current.info.Label, current.info.Power)
		current = current.next
	}
}

// Implement the BoxList methods Insert(), Insert2List(), PerformMinus(), PerformEquals(), Construct(),  below

func (bl *BoxList) Insert2List(key int, info Lense) {
	// Insert the specified info into the list with the specified key
	current := bl.head
	for current != nil {
		if current.key == key {
			current.info.PerformEquals(info)
			return
		}
		current = current.next
	}
}

func (bl *BoxList) PerformMinus(key int, info Lense) {
	// Perform the minus operation on the list with the specified key
	// Delete the Lense from list with the specified key
	current := bl.head
	for current != nil {
		if current.key == key {
			current.info.Delete(info.Label)
			return
		}
		current = current.next
	}
}

func (bl *BoxList) PerformEquals(key int, info Lense) {
	// This has two interpretations:
	// 1. Insert the Lense into the list with the specified key at the end
	// 2. Change the Lense in the list with the specified key with this Lense

	// Interpretation 1
	current := bl.head
	for current != nil {
		if current.key == key {
			current.info.PerformEquals(info)
			return
		}
		current = current.next
	}
}

func (bl *BoxList) Insert(key int) {
	// Insert a new box with the specified key at the end of the list
	if bl.head == nil {
		bl.head = &BoxNode{key: key}
		bl.tail = bl.head
	} else {
		bl.tail.next = &BoxNode{key: key}
		bl.tail = bl.tail.next
	}
}

func (bl *BoxList) Construct() {
	// Construct the list with the specified key and info
	for i := 0; i < 257; i++ {
		bl.Insert(i)
	}
}

func (bl *BoxList) PrintBox() {
	// Print the list
	current := bl.head
	box := 0
	for current != nil {
		if current.info.head == nil {
			current = current.next
			box += 1
			continue
		}
		fmt.Printf("Box[%d]:\n", box)
		current.info.PrintList()
		current = current.next
		box += 1
	}
}

func (bl *BoxList) Sum() int {
	// Return the focusing power of all the lenses in the box
	// that is (1+key)*slow_num*power
	current := bl.head
	sum := 0
	for i := 0; i < 258; i++ {
		if current != nil {
			fmt.Println(current.key)
			lense_lis := current.info
			current_lense := lense_lis.head
			lense_idx := 1
			for current_lense != nil {
				sum += (i + 1) * lense_idx * current_lense.info.Power
				lense_idx += 1
				current_lense = current_lense.next
			}
		}
		if current.next != nil {
			current = current.next
		}
	}
	return sum
}
