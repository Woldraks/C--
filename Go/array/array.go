package array

import (
    "errors"
    "fmt"
)

// Array представляет динамический массив строк
type Array struct {
    data     []string
    capacity int
    size     int
}

// New создает новый массив с начальной емкостью
func New(initialCapacity int) *Array {
    return &Array{
        data:     make([]string, initialCapacity),
        capacity: initialCapacity,
        size:     0,
    }
}

// GetLength возвращает текущее количество элементов
func (a *Array) GetLength() int {
    return a.size
}

// GetCapacity возвращает текущую емкость массива
func (a *Array) GetCapacity() int {
    return a.capacity
}

// extendIfNeeded увеличивает емкость массива вдвое, если нужно
func (a *Array) extendIfNeeded() {
    if a.size == a.capacity {
        a.capacity *= 2
        newData := make([]string, a.capacity)
        copy(newData, a.data[:a.size])
        a.data = newData
    }
}

// Push добавляет элемент в конец массива
func (a *Array) Push(key string) {
    a.extendIfNeeded()
    a.data[a.size] = key
    a.size++
}

// Insert вставляет элемент по указанному индексу
func (a *Array) Insert(index int, key string) error {
    if index < 0 {
        return errors.New("индекс не может быть отрицательным")
    }

    a.extendIfNeeded()

    // Если индекс указывает за пределы массива - добавить в конец
    if a.size <= index {
        a.Push(key)
        return nil
    }

    // Сдвиг элементов вправо для освобождения места
    for i := a.size; i > index; i-- {
        a.data[i] = a.data[i-1]
    }
    a.data[index] = key
    a.size++

    return nil
}

// ReplaceByIndex заменяет элемент по индексу
func (a *Array) ReplaceByIndex(index int, key string) error {
    if index < 0 || index >= a.size {
        return errors.New("выход за пределы массива")
    }
    a.data[index] = key
    return nil
}

// DeleteByIndex удаляет элемент по индексу
func (a *Array) DeleteByIndex(index int) error {
    if index < 0 || index >= a.size {
        return errors.New("выход за пределы массива")
    }

    for i := index; i < a.size-1; i++ {
        a.data[i] = a.data[i+1]
    }
    a.size--

    return nil
}

// GetByIndex возвращает элемент по индексу
func (a *Array) GetByIndex(index int) (string, error) {
    if index < 0 || index >= a.size {
        return "", errors.New("выход за пределы массива")
    }
    return a.data[index], nil
}

// Print выводит содержимое массива
func (a *Array) Print() {
    if a.size == 0 {
        fmt.Println("Массив пуст!")
        return
    }

    fmt.Print("Массив: [")
    for i := 0; i < a.size; i++ {
        fmt.Printf(" %s", a.data[i])
    }
    fmt.Println(" ]")
}

// Clear очищает массив
func (a *Array) Clear() {
    a.size = 0
}

// GetData возвращает копию данных массива
func (a *Array) GetData() []string {
    result := make([]string, a.size)
    copy(result, a.data[:a.size])
    return result
}