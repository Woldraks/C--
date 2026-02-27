package types

// Position определяет позицию для операций со списками
type Position int

const (
    Before Position = iota
    After
    Head
    Tail
)

// String возвращает строковое представление Position
func (p Position) String() string {
    return [...]string{"Before", "After", "Head", "Tail"}[p]
}