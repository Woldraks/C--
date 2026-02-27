package stack

import "fmt"

// StackElement представляет элемент стека
type StackElement struct {
    Key  string
    Prev *StackElement
}

// Stack представляет стек
type Stack struct {
    Last *StackElement
}

// New создает новый пустой стек
func New() *Stack {
    return &Stack{Last: nil}
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
    return s.Last == nil
}

// CreateElement создает новый элемент стека
func CreateElement(key string) *StackElement {
    return &StackElement{Key: key, Prev: nil}
}

// Push добавляет элемент в вершину стека
func (s *Stack) Push(element *StackElement) {
    element.Prev = s.Last
    s.Last = element
}

// Pop удаляет элемент с вершины стека
func (s *Stack) Pop() (string, error) {
    if s.IsEmpty() {
        return "", fmt.Errorf("стек пуст")
    }

    lastPtr := s.Last
    s.Last = lastPtr.Prev
    return lastPtr.Key, nil
}

// Print выводит содержимое стека от вершины к основанию
func (s *Stack) Print() {
    fmt.Print("Стек (сверху вниз): [ ")
    for curr := s.Last; curr != nil; curr = curr.Prev {
        fmt.Printf("%s ", curr.Key)
    }
    fmt.Println("]")
}

// Clear очищает стек
func (s *Stack) Clear() {
    s.Last = nil
}

// Peek возвращает вершину стека без удаления
func (s *Stack) Peek() (string, error) {
    if s.IsEmpty() {
        return "", fmt.Errorf("стек пуст")
    }
    return s.Last.Key, nil
}