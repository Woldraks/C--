package queue

import "fmt"

// QueueElement представляет элемент очереди
type QueueElement struct {
    Key  string
    Next *QueueElement
}

// Queue представляет очередь
type Queue struct {
    First *QueueElement
    Last  *QueueElement
}

// New создает новую пустую очередь
func New() *Queue {
    return &Queue{First: nil, Last: nil}
}

// IsEmpty проверяет, пуста ли очередь
func (q *Queue) IsEmpty() bool {
    return q.First == nil
}

// CreateElement создает новый элемент очереди
func CreateElement(key string) *QueueElement {
    return &QueueElement{Key: key, Next: nil}
}

// Push добавляет элемент в конец очереди
func (q *Queue) Push(element *QueueElement) {
    element.Next = nil

    if q.IsEmpty() {
        q.First = element
        q.Last = element
        return
    }

    q.Last.Next = element
    q.Last = element
}

// Pop извлекает элемент из начала очереди
func (q *Queue) Pop() (string, error) {
    if q.IsEmpty() {
        return "", fmt.Errorf("очередь пуста")
    }

    firstPtr := q.First
    q.First = firstPtr.Next
    if q.First == nil {
        q.Last = nil
    }
    return firstPtr.Key, nil
}

// Print выводит содержимое очереди
func (q *Queue) Print() {
    fmt.Print("Очередь: [ ")
    for curr := q.First; curr != nil; curr = curr.Next {
        fmt.Printf("%s ", curr.Key)
    }
    fmt.Println("]")
}

// Clear очищает очередь
func (q *Queue) Clear() {
    q.First = nil
    q.Last = nil
}

// Peek возвращает первый элемент без удаления
func (q *Queue) Peek() (string, error) {
    if q.IsEmpty() {
        return "", fmt.Errorf("очередь пуста")
    }
    return q.First.Key, nil
}