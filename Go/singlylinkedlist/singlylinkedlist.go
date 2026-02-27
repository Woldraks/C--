package singlylinkedlist

import (
	"dbms/types"
	"errors"
	"fmt"
)

// SingleNode представляет узел односвязного списка
type SingleNode struct {
	Key  string
	Next *SingleNode
}

// ForwardList представляет односвязный список
type ForwardList struct {
	Head *SingleNode
}

// New создает новый пустой список
func New() *ForwardList {
	return &ForwardList{Head: nil}
}

// IsEmpty проверяет, пуст ли список
func (fl *ForwardList) IsEmpty() bool {
	return fl.Head == nil
}

// CreateNode создает новый узел
func CreateNode(key string) *SingleNode {
	return &SingleNode{Key: key, Next: nil}
}

// findPrevNode находит узел, предшествующий заданному
func (fl *ForwardList) findPrevNode(node *SingleNode) *SingleNode {
	if fl.Head == nil || node == fl.Head {
		return nil
	}

	curr := fl.Head
	for curr != nil && curr.Next != node {
		curr = curr.Next
	}
	return curr
}

// FindNodeByValue находит узел по значению
func (fl *ForwardList) FindNodeByValue(key string) *SingleNode {
	curr := fl.Head
	for curr != nil {
		if curr.Key == key {
			return curr
		}
		curr = curr.Next
	}
	return nil
}

// AddNode добавляет узел в указанную позицию
func (fl *ForwardList) AddNode(pos types.Position, addingNode *SingleNode, node *SingleNode) error {
	if fl.IsEmpty() {
		fl.Head = addingNode
		return nil
	}

	switch pos {
	case types.Before:
		if node == nil {
			return errors.New("при добавлении до нужно указать целевой узел")
		}
		if node == fl.Head {
			fl.addToHead(addingNode)
		} else {
			prev := fl.findPrevNode(node)
			if prev == nil {
				return errors.New("целевой узел не найден в списке")
			}
			addingNode.Next = node
			prev.Next = addingNode
		}

	case types.After:
		if node == nil {
			return errors.New("при добавлении после нужно указать целевой узел")
		}
		addingNode.Next = node.Next
		node.Next = addingNode

	case types.Head:
		fl.addToHead(addingNode)

	case types.Tail:
		fl.addToTail(addingNode)
	}

	return nil
}

// addToHead добавляет узел в начало списка
func (fl *ForwardList) addToHead(addingNode *SingleNode) {
	addingNode.Next = fl.Head
	fl.Head = addingNode
}

// addToTail добавляет узел в конец списка
func (fl *ForwardList) addToTail(addingNode *SingleNode) {
	addingNode.Next = nil
	if fl.Head == nil {
		fl.Head = addingNode
		return
	}

	curr := fl.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = addingNode
}

// RemoveNode удаляет узел из указанной позиции
func (fl *ForwardList) RemoveNode(pos types.Position, node *SingleNode) error {
	if fl.IsEmpty() {
		return errors.New("список пуст")
	}

	switch pos {
	case types.Before:
		if node == nil {
			return errors.New("при удалении до нужно указать целевой узел")
		}
		if node == fl.Head {
			return errors.New("не существует узла перед головой")
		}
		prev := fl.findPrevNode(node)
		if prev == nil {
			return errors.New("целевой узел не найден")
		}
		beforePrev := fl.findPrevNode(prev)
		if beforePrev == nil {
			fl.Head = node
		} else {
			beforePrev.Next = node
		}

	case types.After:
		if node == nil {
			return errors.New("при удалении после нужно указать целевой узел")
		}
		if node.Next == nil {
			return errors.New("нет узла после хвоста")
		}
		toDelete := node.Next
		node.Next = toDelete.Next

	case types.Head:
		toDelete := fl.Head
		fl.Head = toDelete.Next

	case types.Tail:
		if fl.Head.Next == nil {
			fl.Head = nil
			return nil
		}

		curr := fl.Head
		for curr.Next != nil && curr.Next.Next != nil {
			curr = curr.Next
		}
		curr.Next = nil
	}

	return nil
}

// RemoveNodeByValue удаляет узел по значению
func (fl *ForwardList) RemoveNodeByValue(key string) error {
	foundNode := fl.FindNodeByValue(key)
	if foundNode == nil {
		return fmt.Errorf("узел со значением %s не найден", key)
	}

	if foundNode == fl.Head {
		return fl.RemoveNode(types.Head, nil)
	}

	prev := fl.findPrevNode(foundNode)
	if prev == nil {
		return errors.New("ошибка: не найден предыдущий узел")
	}
	prev.Next = foundNode.Next
	return nil
}

// Print выводит список в прямом порядке
func (fl *ForwardList) Print() {
	fmt.Print("Список: [ ")
	curr := fl.Head
	for curr != nil {
		fmt.Printf("%s ", curr.Key)
		curr = curr.Next
	}
	fmt.Println("]")
}

// printReversedRec вспомогательная рекурсивная функция для обратного вывода
func printReversedRec(node *SingleNode) {
	if node == nil {
		return
	}
	printReversedRec(node.Next)
	fmt.Printf("%s ", node.Key)
}

// PrintReversed выводит список в обратном порядке
func (fl *ForwardList) PrintReversed() {
	fmt.Print("Список в обратном порядке: [ ")
	printReversedRec(fl.Head)
	fmt.Println("]")
}

// Clear очищает список
func (fl *ForwardList) Clear() {
	fl.Head = nil
}
