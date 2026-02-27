package doublylinkedlist

import (
	"dbms/types"
	"errors"
	"fmt"
)

// DoubleNode представляет узел двусвязного списка
type DoubleNode struct {
	Key  string
	Next *DoubleNode
	Prev *DoubleNode
}

// LinkedList представляет двусвязный список
type LinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

// New создает новый пустой список
func New() *LinkedList {
	return &LinkedList{Head: nil, Tail: nil}
}

// IsEmpty проверяет, пуст ли список
func (ll *LinkedList) IsEmpty() bool {
	return ll.Head == nil
}

// CreateNode создает новый узел
func CreateNode(key string) *DoubleNode {
	return &DoubleNode{Key: key, Next: nil, Prev: nil}
}

// FindNodeByValue находит узел по значению
func (ll *LinkedList) FindNodeByValue(key string) *DoubleNode {
	if ll.IsEmpty() {
		return nil
	}

	curr := ll.Head
	for curr != nil {
		if curr.Key == key {
			return curr
		}
		curr = curr.Next
	}
	return nil
}

// AddNode добавляет узел в указанную позицию
func (ll *LinkedList) AddNode(pos types.Position, addingNode *DoubleNode, node *DoubleNode) error {
	if ll.IsEmpty() {
		ll.Head = addingNode
		ll.Tail = addingNode
		return nil
	}

	switch pos {
	case types.Before:
		if node == nil {
			return errors.New("при добавлении узла ДО нужно указать узел, перед которым нужно добавить узел")
		}
		if node == ll.Head {
			ll.addToHead(addingNode)
		} else {
			addingNode.Next = node
			addingNode.Prev = node.Prev
			node.Prev.Next = addingNode
			node.Prev = addingNode
		}

	case types.After:
		if node == nil {
			return errors.New("при добавлении узла ПОСЛЕ нужно указать узел, после которого нужно добавить узел")
		}
		if node == ll.Tail {
			ll.addToTail(addingNode)
		} else {
			addingNode.Prev = node
			addingNode.Next = node.Next
			node.Next.Prev = addingNode
			node.Next = addingNode
		}

	case types.Head:
		ll.addToHead(addingNode)

	case types.Tail:
		ll.addToTail(addingNode)
	}

	return nil
}

// addToHead добавляет узел в начало списка
func (ll *LinkedList) addToHead(addingNode *DoubleNode) {
	addingNode.Next = ll.Head
	addingNode.Prev = nil
	if ll.Head != nil {
		ll.Head.Prev = addingNode
	}
	ll.Head = addingNode

	if ll.Tail == nil {
		ll.Tail = addingNode
	}
}

// addToTail добавляет узел в конец списка
func (ll *LinkedList) addToTail(addingNode *DoubleNode) {
	addingNode.Prev = ll.Tail
	addingNode.Next = nil
	if ll.Tail != nil {
		ll.Tail.Next = addingNode
	}
	ll.Tail = addingNode

	if ll.Head == nil {
		ll.Head = addingNode
	}
}

// RemoveNode удаляет узел из указанной позиции
func (ll *LinkedList) RemoveNode(pos types.Position, node *DoubleNode) error {
	if ll.IsEmpty() {
		return errors.New("список пуст")
	}

	switch pos {
	case types.Before:
		if node == nil {
			return errors.New("при удалении узла ДО нужно указать узел, перед которым нужно удалить узел")
		}
		if node == ll.Head {
			return errors.New("невозможно удалить узел ДО головы списка")
		}

		prevNode := node.Prev
		if prevNode.Prev == nil {
			ll.Head = node
			node.Prev = nil
		} else {
			prevNode.Prev.Next = node
			node.Prev = prevNode.Prev
		}

	case types.After:
		if node == nil {
			return errors.New("при удалении узла ПОСЛЕ нужно указать узел, после которого нужно удалить узел")
		}
		if node == ll.Tail {
			return errors.New("нечего удалять, мы в хвосте")
		}

		nextPtr := node.Next
		if nextPtr == ll.Tail {
			ll.Tail = node
			node.Next = nil
		} else {
			node.Next = nextPtr.Next
			nextPtr.Next.Prev = node
		}

	case types.Head:
		headPtr := ll.Head
		if ll.Head == ll.Tail {
			ll.Head = nil
			ll.Tail = nil
		} else {
			ll.Head = headPtr.Next
			ll.Head.Prev = nil
		}

	case types.Tail:
		tailPtr := ll.Tail
		if ll.Head == tailPtr {
			ll.Head = nil
			ll.Tail = nil
		} else {
			ll.Tail = tailPtr.Prev
			ll.Tail.Next = nil
		}
	}

	return nil
}

// RemoveNodeByValue удаляет узел по значению
func (ll *LinkedList) RemoveNodeByValue(key string) error {
	foundNode := ll.FindNodeByValue(key)
	if foundNode == nil {
		return fmt.Errorf("узел со значением %s не найден", key)
	}

	if foundNode == ll.Head {
		return ll.RemoveNode(types.Head, nil)
	}
	if foundNode == ll.Tail {
		return ll.RemoveNode(types.Tail, nil)
	}

	foundNode.Prev.Next = foundNode.Next
	foundNode.Next.Prev = foundNode.Prev
	return nil
}

// Print выводит список в прямом порядке
func (ll *LinkedList) Print() {
	fmt.Print("Список: [ ")
	curr := ll.Head
	for curr != nil {
		fmt.Printf("%s ", curr.Key)
		curr = curr.Next
	}
	fmt.Println("]")
}

// PrintReversed выводит список в обратном порядке
func (ll *LinkedList) PrintReversed() {
	fmt.Print("Список в обратном порядке: [ ")
	curr := ll.Tail
	for curr != nil {
		fmt.Printf("%s ", curr.Key)
		curr = curr.Prev
	}
	fmt.Println("]")
}

// Clear очищает список
func (ll *LinkedList) Clear() {
	ll.Head = nil
	ll.Tail = nil
}
