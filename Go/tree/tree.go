package tree

import (
	"errors"
	"fmt"
	"strings"
)

// TreeNode представляет узел бинарного дерева
type TreeNode struct {
	Key    string
	Left   *TreeNode
	Right  *TreeNode
}

// BST представляет бинарное дерево поиска
type BST struct {
	Root *TreeNode
	size int
}

// New создает новое пустое дерево
func New() *BST {
	return &BST{Root: nil, size: 0}
}

// IsEmpty проверяет, пусто ли дерево
func (t *BST) IsEmpty() bool {
	return t.Root == nil
}

// Size возвращает количество узлов в дереве
func (t *BST) Size() int {
	return t.size
}

// CreateNode создает новый узел
func CreateNode(key string) *TreeNode {
	return &TreeNode{Key: key, Left: nil, Right: nil}
}

// Insert вставляет узел в дерево
func (t *BST) Insert(key string) {
	newNode := CreateNode(key)
	if t.IsEmpty() {
		t.Root = newNode
		t.size++
		return
	}
	t.insertRecursive(t.Root, newNode)
}

// insertRecursive рекурсивно вставляет узел
func (t *BST) insertRecursive(current, newNode *TreeNode) {
	if newNode.Key < current.Key {
		if current.Left == nil {
			current.Left = newNode
			t.size++
		} else {
			t.insertRecursive(current.Left, newNode)
		}
	} else if newNode.Key > current.Key {
		if current.Right == nil {
			current.Right = newNode
			t.size++
		} else {
			t.insertRecursive(current.Right, newNode)
		}
	}
	// Если ключ равен существующему - не вставляем (в дереве уникальные значения)
}

// Find ищет узел по значению
func (t *BST) Find(key string) *TreeNode {
	return t.findRecursive(t.Root, key)
}

func (t *BST) findRecursive(current *TreeNode, key string) *TreeNode {
	if current == nil || current.Key == key {
		return current
	}
	
	if key < current.Key {
		return t.findRecursive(current.Left, key)
	}
	return t.findRecursive(current.Right, key)
}

// Delete удаляет узел по значению
func (t *BST) Delete(key string) error {
	if t.IsEmpty() {
		return errors.New("дерево пусто")
	}
	
	var parent *TreeNode
	current := t.Root
	
	// Поиск удаляемого узла и его родителя
	for current != nil && current.Key != key {
		parent = current
		if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	
	if current == nil {
		return fmt.Errorf("узел со значением %s не найден", key)
	}
	
	// Случай 1: Узел - лист (нет детей)
	if current.Left == nil && current.Right == nil {
		if parent == nil {
			t.Root = nil
		} else if parent.Left == current {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		t.size--
		return nil
	}
	
	// Случай 2: Узел имеет только одного ребенка
	if current.Left == nil || current.Right == nil {
		var child *TreeNode
		if current.Left != nil {
			child = current.Left
		} else {
			child = current.Right
		}
		
		if parent == nil {
			t.Root = child
		} else if parent.Left == current {
			parent.Left = child
		} else {
			parent.Right = child
		}
		t.size--
		return nil
	}
	
	// Случай 3: Узел имеет двух детей
	// Находим наименьший узел в правом поддереве (преемник)
	successorParent := current
	successor := current.Right
	for successor.Left != nil {
		successorParent = successor
		successor = successor.Left
	}
	
	// Заменяем значение текущего узла на значение преемника
	current.Key = successor.Key
	
	// Удаляем преемника
	if successorParent.Left == successor {
		successorParent.Left = successor.Right
	} else {
		successorParent.Right = successor.Right
	}
	t.size--
	return nil
}

// Min находит минимальный элемент в дереве
func (t *BST) Min() (string, error) {
	if t.IsEmpty() {
		return "", errors.New("дерево пусто")
	}
	
	current := t.Root
	for current.Left != nil {
		current = current.Left
	}
	return current.Key, nil
}

// Max находит максимальный элемент в дереве
func (t *BST) Max() (string, error) {
	if t.IsEmpty() {
		return "", errors.New("дерево пусто")
	}
	
	current := t.Root
	for current.Right != nil {
		current = current.Right
	}
	return current.Key, nil
}

// InOrder обход дерева (Левый-Корень-Правый) - сортированный порядок
func (t *BST) InOrder() []string {
	var result []string
	t.inOrderRecursive(t.Root, &result)
	return result
}

func (t *BST) inOrderRecursive(node *TreeNode, result *[]string) {
	if node != nil {
		t.inOrderRecursive(node.Left, result)
		*result = append(*result, node.Key)
		t.inOrderRecursive(node.Right, result)
	}
}

// PreOrder обход дерева (Корень-Левый-Правый)
func (t *BST) PreOrder() []string {
	var result []string
	t.preOrderRecursive(t.Root, &result)
	return result
}

func (t *BST) preOrderRecursive(node *TreeNode, result *[]string) {
	if node != nil {
		*result = append(*result, node.Key)
		t.preOrderRecursive(node.Left, result)
		t.preOrderRecursive(node.Right, result)
	}
}

// PostOrder обход дерева (Левый-Правый-Корень)
func (t *BST) PostOrder() []string {
	var result []string
	t.postOrderRecursive(t.Root, &result)
	return result
}

func (t *BST) postOrderRecursive(node *TreeNode, result *[]string) {
	if node != nil {
		t.postOrderRecursive(node.Left, result)
		t.postOrderRecursive(node.Right, result)
		*result = append(*result, node.Key)
	}
}

// Height возвращает высоту дерева
func (t *BST) Height() int {
	return t.heightRecursive(t.Root)
}

func (t *BST) heightRecursive(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := t.heightRecursive(node.Left)
	rightHeight := t.heightRecursive(node.Right)
	
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// PrintLevelOrder выводит дерево по уровням (для наглядности)
func (t *BST) PrintLevelOrder() {
	if t.IsEmpty() {
		fmt.Println("Дерево пусто!")
		return
	}
	
	height := t.Height()
	for i := 1; i <= height; i++ {
		fmt.Printf("Уровень %d: ", i)
		t.printLevel(t.Root, i)
		fmt.Println()
	}
}

func (t *BST) printLevel(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if level == 1 {
		fmt.Printf("%s ", node.Key)
	} else if level > 1 {
		t.printLevel(node.Left, level-1)
		t.printLevel(node.Right, level-1)
	}
}

// Print визуализирует дерево (упрощенная версия)
func (t *BST) Print() {
	if t.IsEmpty() {
		fmt.Println("Дерево пусто!")
		return
	}
	fmt.Println("Структура дерева:")
	t.printTree(t.Root, 0, "")
}

func (t *BST) printTree(node *TreeNode, level int, prefix string) {
	if node == nil {
		return
	}
	
	if level == 0 {
		fmt.Printf("%s\n", node.Key)
	} else {
		fmt.Printf("%s└── %s\n", prefix, node.Key)
	}
	
	if node.Left != nil || node.Right != nil {
		// Для левого поддерева используем "│   " если есть правый ребенок
		if node.Right != nil {
			t.printTree(node.Left, level+1, prefix+"│   ")
		} else {
			t.printTree(node.Left, level+1, prefix+"    ")
		}
		
		// Для правого поддерева всегда используем "    "
		t.printTree(node.Right, level+1, prefix+"    ")
	}
}

// Clear очищает дерево
func (t *BST) Clear() {
	t.Root = nil
	t.size = 0
}