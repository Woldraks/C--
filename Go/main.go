package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"dbms/array"
	"dbms/doublylinkedlist"
	"dbms/queue"
	"dbms/singlylinkedlist"
	"dbms/stack"
	"dbms/tree"
	"dbms/types"
)

// Глобальные структуры данных
var (
	globalArray       *array.Array
	globalForwardList *singlylinkedlist.ForwardList
	globalLinkedList  *doublylinkedlist.LinkedList
	globalQueue       *queue.Queue
	globalStack       *stack.Stack
	globalTree        *tree.BST
)

func printHelpMessage() {
	fmt.Println("  === ОПЕРАЦИИ ДОБАВЛЕНИЯ ===")
	fmt.Println("  MPUSH <value>                  - добавить в массив")
	fmt.Println("  MINSERT <index> <value>        - вставить в массив по индексу")
	fmt.Println("  FPUSHHEAD <value>              - добавить в начало односвязного списка")
	fmt.Println("  FPUSHTAIL <value>              - добавить в конец односвязного списка")
	fmt.Println("  FINSERTBEFORE <target> <value> - вставить перед элементом")
	fmt.Println("  FINSERTAFTER <target> <value>  - вставить после элемента")
	fmt.Println("  LPUSHHEAD <value>              - добавить в начало двусвязного списка")
	fmt.Println("  LPUSHTAIL <value>              - добавить в конец двусвязного списка")
	fmt.Println("  LINSERTBEFORE <target> <value> - вставить перед элементом")
	fmt.Println("  LINSERTAFTER <target> <value>  - вставить после элемента")
	fmt.Println("  QPUSH <value>                  - добавить в очередь")
	fmt.Println("  SPUSH <value>                  - добавить в стек")
	fmt.Println("  TINSERT <value>                - вставить значение в дерево")
	fmt.Println("")
	fmt.Println("  === ОПЕРАЦИИ УДАЛЕНИЯ ===")
	fmt.Println("  MDEL <index>                   - удалить из массива по индексу")
	fmt.Println("  FDEL <value>                   - удалить из односвязного списка по значению")
	fmt.Println("  FDELHEAD                       - удалить из головы односвязного списка")
	fmt.Println("  FDELTAIL                       - удалить из хвоста односвязно списка")
	fmt.Println("  FDELBEFORE <value>             - удалить из односвязного списка перед значением")
	fmt.Println("  FDELAFTER <value>              - удалить из односвязного списка после значения")
	fmt.Println("  LDEL <value>                   - удалить из двусвязного списка по значению")
	fmt.Println("  LDELHEAD                       - удалить из головы двусвязного списка")
	fmt.Println("  LDELTAIL                       - удалить из хвоста двусвязного списка")
	fmt.Println("  LDELBEFORE <value>             - удалить из двусвязного списка перед значением")
	fmt.Println("  LDELAFTER <value>              - удалить из двусвязного списка после значения")
	fmt.Println("  QPOP                           - удалить из очереди (из начала)")
	fmt.Println("  SPOP                           - удалить из стека (с вершины)")
	fmt.Println("  TDEL <value>                   - удалить значение из дерева")
	fmt.Println("")
	fmt.Println("  === ОПЕРАЦИИ ЧТЕНИЯ ===")
	fmt.Println("  MGET <index>                   - получить элемент массива по индексу")
	fmt.Println("  TFIND <value>                  - найти значение в дереве")
	fmt.Println("  TMIN                           - найти минимальное значение")
	fmt.Println("  TMAX                           - найти максимальное значение")
	fmt.Println("  THEIGHT                        - получить высоту дерева")
	fmt.Println("  TSIZE                          - получить количество узлов")
	fmt.Println("")
	fmt.Println("  === ОПЕРАЦИИ ВЫВОДА ===")
	fmt.Println("  PRINT [structure]              - вывести все структуры или конкретную (M,F,L,Q,S,T)")
	fmt.Println("  TPRINTINORDER                   - вывести дерево в порядке возрастания")
	fmt.Println("  TPRINTPREORDER                   - вывести дерево в прямом порядке")
	fmt.Println("  TPRINTPOSTORDER                  - вывести дерево в обратном порядке")
	fmt.Println("  TPRINTTREE                       - вывести структуру дерева")
	fmt.Println("  CLEAR <structure>              - очистить конкретную структуру (M,F,L,Q,S,T)")
	fmt.Println("----------------")
}

// executeCommand выполняет команду
func executeCommand(tokens []string) {
	if len(tokens) == 0 {
		return
	}

	command := tokens[0]

	// Массив (M)
	switch command {
	case "MPUSH":
		if len(tokens) > 1 {
			if globalArray == nil {
				globalArray = array.New(8)
			}
			globalArray.Push(tokens[1])
		}

	case "MINSERT":
		if len(tokens) > 2 {
			if globalArray == nil {
				globalArray = array.New(8)
			}
			index, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Println("Ошибка: неверный индекс")
				return
			}
			err = globalArray.Insert(index, tokens[2])
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		}

	case "MREPLACE":
		if len(tokens) > 2 && globalArray != nil {
			index, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Println("Ошибка: неверный индекс")
				return
			}
			err = globalArray.ReplaceByIndex(index, tokens[2])
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		}

	case "MDEL":
		if len(tokens) > 1 && globalArray != nil {
			index, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Println("Ошибка: неверный индекс")
				return
			}
			err = globalArray.DeleteByIndex(index)
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		}

	case "MGET":
		if len(tokens) > 1 && globalArray != nil {
			index, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Println("Ошибка: неверный индекс")
				return
			}
			value, err := globalArray.GetByIndex(index)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("MGET [%s] = %s\n", tokens[1], value)
			}
		}

	// Односвязный список (F)
	case "FPUSHHEAD":
		if len(tokens) > 1 {
			if globalForwardList == nil {
				globalForwardList = singlylinkedlist.New()
			}
			globalForwardList.AddNode(types.Head, singlylinkedlist.CreateNode(tokens[1]), nil)
		}

	case "FPUSHTAIL":
		if len(tokens) > 1 {
			if globalForwardList == nil {
				globalForwardList = singlylinkedlist.New()
			}
			globalForwardList.AddNode(types.Tail, singlylinkedlist.CreateNode(tokens[1]), nil)
		}

	case "FINSERTAFTER":
		if len(tokens) > 2 && globalForwardList != nil {
			found := globalForwardList.FindNodeByValue(tokens[1])
			if found != nil {
				globalForwardList.AddNode(types.After, singlylinkedlist.CreateNode(tokens[2]), found)
			} else {
				fmt.Printf("Узел со значением %s не найден\n", tokens[1])
			}
		}

	case "FINSERTBEFORE":
		if len(tokens) > 2 && globalForwardList != nil {
			found := globalForwardList.FindNodeByValue(tokens[1])
			if found != nil {
				globalForwardList.AddNode(types.Before, singlylinkedlist.CreateNode(tokens[2]), found)
			} else {
				fmt.Printf("Узел со значением %s не найден\n", tokens[1])
			}
		}

	case "FDEL":
		if len(tokens) > 1 && globalForwardList != nil {
			err := globalForwardList.RemoveNodeByValue(tokens[1])
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		}

	case "FDELHEAD":
		if globalForwardList != nil {
			err := globalForwardList.RemoveNode(types.Head, nil)
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		}

	case "FDELTAIL":
		if globalForwardList != nil {
			err := globalForwardList.RemoveNode(types.Tail, nil)
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		}

	case "FDELBEFORE":
		if len(tokens) > 1 && globalForwardList != nil {
			found := globalForwardList.FindNodeByValue(tokens[1])
			if found != nil {
				err := globalForwardList.RemoveNode(types.Before, found)
				if err != nil {
					fmt.Println("Ошибка:", err)
				}
			} else {
				fmt.Printf("Узел со значением %s не найден\n", tokens[1])
			}
		}

	case "FDELAFTER":
		if len(tokens) > 1 && globalForwardList != nil {
			found := globalForwardList.FindNodeByValue(tokens[1])
			if found != nil {
				err := globalForwardList.RemoveNode(types.After, found)
				if err != nil {
					fmt.Println("Ошибка:", err)
				}
			} else {
				fmt.Printf("Узел со значением %s не найден\n", tokens[1])
			}
		}

	// Двусвязный список (L)
	case "LPUSHHEAD":
		if len(tokens) > 1 {
			if globalLinkedList == nil {
				globalLinkedList = doublylinkedlist.New()
			}
			globalLinkedList.AddNode(types.Head, doublylinkedlist.CreateNode(tokens[1]), nil)
		}

	case "LPUSHTAIL":
		if len(tokens) > 1 {
			if globalLinkedList == nil {
				globalLinkedList = doublylinkedlist.New()
			}
			globalLinkedList.AddNode(types.Tail, doublylinkedlist.CreateNode(tokens[1]), nil)
		}

	case "LINSERTAFTER":
		if len(tokens) > 2 && globalLinkedList != nil {
			found := globalLinkedList.FindNodeByValue(tokens[1])
			if found != nil {
				globalLinkedList.AddNode(types.After, doublylinkedlist.CreateNode(tokens[2]), found)
			} else {
				fmt.Printf("Узел со значением %s не найден\n", tokens[1])
			}
		}

	case "LINSERTBEFORE":
		if len(tokens) > 2 && globalLinkedList != nil {
			found := globalLinkedList.FindNodeByValue(tokens[1])
			if found != nil {
				globalLinkedList.AddNode(types.Before, doublylinkedlist.CreateNode(tokens[2]), found)
			} else {
				fmt.Printf("Узел со значением %s не найден\n", tokens[1])
			}
		}

	case "LDEL":
		if len(tokens) > 1 && globalLinkedList != nil {
			err := globalLinkedList.RemoveNodeByValue(tokens[1])
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		}

	case "LDELHEAD":
		if globalLinkedList != nil {
			err := globalLinkedList.RemoveNode(types.Head, nil)
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		}

	case "LDELTAIL":
		if globalLinkedList != nil {
			err := globalLinkedList.RemoveNode(types.Tail, nil)
			if err != nil {
				fmt.Println("Ошибка:", err)
			}
		}

	case "LDELBEFORE":
		if len(tokens) > 1 && globalLinkedList != nil {
			found := globalLinkedList.FindNodeByValue(tokens[1])
			if found != nil {
				err := globalLinkedList.RemoveNode(types.Before, found)
				if err != nil {
					fmt.Println("Ошибка:", err)
				}
			} else {
				fmt.Printf("Узел со значением %s не найден\n", tokens[1])
			}
		}

	case "LDELAFTER":
		if len(tokens) > 1 && globalLinkedList != nil {
			found := globalLinkedList.FindNodeByValue(tokens[1])
			if found != nil {
				err := globalLinkedList.RemoveNode(types.After, found)
				if err != nil {
					fmt.Println("Ошибка:", err)
				}
			} else {
				fmt.Printf("Узел со значением %s не найден\n", tokens[1])
			}
		}

	// Очередь (Q)
	case "QPUSH":
		if len(tokens) > 1 {
			if globalQueue == nil {
				globalQueue = queue.New()
			}
			globalQueue.Push(queue.CreateElement(tokens[1]))
		}

	case "QPOP":
		if globalQueue != nil && !globalQueue.IsEmpty() {
			value, err := globalQueue.Pop()
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Вырезанное значение из очереди: %s\n", value)
			}
		}

	// Стек (S)
	case "SPUSH":
		if len(tokens) > 1 {
			if globalStack == nil {
				globalStack = stack.New()
			}
			globalStack.Push(stack.CreateElement(tokens[1]))
		}

	case "SPOP":
		if globalStack != nil && !globalStack.IsEmpty() {
			value, err := globalStack.Pop()
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Вырезанное значение из стека: %s\n", value)
			}
		}

	// Дерево (T)
	case "TINSERT":
		if len(tokens) > 1 {
			if globalTree == nil {
				globalTree = tree.New()
			}
			globalTree.Insert(tokens[1])
			fmt.Printf("Вставлено значение: %s\n", tokens[1])
		}

	case "TDEL":
		if len(tokens) > 1 && globalTree != nil {
			err := globalTree.Delete(tokens[1])
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Удалено значение: %s\n", tokens[1])
			}
		}

	case "TFIND":
		if len(tokens) > 1 && globalTree != nil {
			found := globalTree.Find(tokens[1])
			if found != nil {
				fmt.Printf("Значение '%s' найдено в дереве\n", tokens[1])
			} else {
				fmt.Printf("Значение '%s' не найдено\n", tokens[1])
			}
		}

	case "TMIN":
		if globalTree != nil {
			min, err := globalTree.Min()
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Минимальное значение: %s\n", min)
			}
		}

	case "TMAX":
		if globalTree != nil {
			max, err := globalTree.Max()
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Максимальное значение: %s\n", max)
			}
		}

	case "THEIGHT":
		if globalTree != nil {
			height := globalTree.Height()
			fmt.Printf("Высота дерева: %d\n", height)
		}

	case "TSIZE":
		if globalTree != nil {
			size := globalTree.Size()
			fmt.Printf("Количество узлов: %d\n", size)
		}

	case "TPRINTINORDER":
		if globalTree != nil {
			result := globalTree.InOrder()
			fmt.Printf("In-order обход (возрастание): %v\n", result)
		}

	case "TPRINTPREORDER":
		if globalTree != nil {
			result := globalTree.PreOrder()
			fmt.Printf("Pre-order обход (корень-левый-правый): %v\n", result)
		}

	case "TPRINTPOSTORDER":
		if globalTree != nil {
			result := globalTree.PostOrder()
			fmt.Printf("Post-order обход (левый-правый-корень): %v\n", result)
		}

	case "TPRINTTREE":
		if globalTree != nil {
			globalTree.Print()
		}

	// Операции вывода
	case "PRINT":
		if len(tokens) > 1 {
			structure := tokens[1]
			switch structure {
			case "M":
				if globalArray != nil {
					fmt.Print("Array: ")
					globalArray.Print()
				}
			case "F":
				if globalForwardList != nil {
					fmt.Print("Forward List: ")
					globalForwardList.Print()
				}
			case "L":
				if globalLinkedList != nil {
					fmt.Print("Linked List: ")
					globalLinkedList.Print()
				}
			case "Q":
				if globalQueue != nil {
					fmt.Print("Queue: ")
					globalQueue.Print()
				}
			case "S":
				if globalStack != nil {
					fmt.Print("Stack: ")
					globalStack.Print()
				}
			case "T":
				if globalTree != nil {
					fmt.Print("Tree: ")
					globalTree.Print()
				}
			}
		} else {
			// Вывод всех структур
			if globalArray != nil {
				fmt.Print("Array: ")
				globalArray.Print()
			}
			if globalForwardList != nil {
				fmt.Print("Forward List: ")
				globalForwardList.Print()
			}
			if globalLinkedList != nil {
				fmt.Print("Linked List: ")
				globalLinkedList.Print()
			}
			if globalQueue != nil {
				fmt.Print("Queue: ")
				globalQueue.Print()
			}
			if globalStack != nil {
				fmt.Print("Stack: ")
				globalStack.Print()
			}
			if globalTree != nil {
				fmt.Print("Tree: ")
				globalTree.Print()
			}
		}

	case "CLEAR":
		if len(tokens) > 1 {
			structure := tokens[1]
			switch structure {
			case "M":
				if globalArray != nil {
					globalArray.Clear()
					globalArray = nil
					fmt.Println("Массив очищен")
				}
			case "F":
				if globalForwardList != nil {
					globalForwardList.Clear()
					globalForwardList = nil
					fmt.Println("Односвязный список очищен")
				}
			case "L":
				if globalLinkedList != nil {
					globalLinkedList.Clear()
					globalLinkedList = nil
					fmt.Println("Двусвязный список очищен")
				}
			case "Q":
				if globalQueue != nil {
					globalQueue.Clear()
					globalQueue = nil
					fmt.Println("Очередь очищена")
				}
			case "S":
				if globalStack != nil {
					globalStack.Clear()
					globalStack = nil
					fmt.Println("Стек очищен")
				}
			case "T":
				if globalTree != nil {
					globalTree.Clear()
					globalTree = nil
					fmt.Println("Дерево очищено")
				}
			}
		}

	case "HELP", "--help", "-h":
		printHelpMessage()

	default:
		fmt.Printf("Неизвестная команда или параметр: %s\n", command)
	}
}

// executeCommandsFromFile выполняет команды из файла
func executeCommandsFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Не удалось открыть файл: %s\n", filename)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			tokens := strings.Fields(line)
			executeCommand(tokens)
		}
	}
}

// appendCommandToFile добавляет команду в файл
func appendCommandToFile(filename string, command string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(command + "\n"); err != nil {
		fmt.Printf("Ошибка при записи в файл: %v\n", err)
	} else {
		fmt.Printf("В файл добавлена команда: %s\n", command)
	}
}

// removeLastCommandFromFile удаляет последнюю команду из файла
func removeLastCommandFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Не удалось открыть файл: %s\n", filename)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			lines = append(lines, line)
		}
	}

	if len(lines) > 0 {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Ошибка при создании файла: %v\n", err)
			return
		}
		defer file.Close()

		for i := 0; i < len(lines)-1; i++ {
			if _, err := file.WriteString(lines[i] + "\n"); err != nil {
				fmt.Printf("Ошибка при записи в файл: %v\n", err)
				return
			}
		}
		fmt.Printf("Последняя команда из файла %s отменена\n", filename)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Использование: %s --file <filename> [--query <command>] [--undo]\n", os.Args[0])
		fmt.Println("Используйте --help для просмотра всех команд")
		return
	}

	var filename string
	var query string
	var undo bool

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch arg {
		case "--file":
			if i+1 < len(os.Args) {
				filename = os.Args[i+1]
				i++
			}
		case "--query":
			if i+1 < len(os.Args) {
				query = os.Args[i+1]
				i++
			}
		case "--undo":
			undo = true
		case "--help", "-h":
			printHelpMessage()
			return
		}
	}

	if filename == "" {
		fmt.Printf("Использование: %s --file <filename> [--query <command>] [--undo]\n", os.Args[0])
		return
	}

	if undo {
		removeLastCommandFromFile(filename)
	} else if query != "" {
		appendCommandToFile(filename, query)
	}

	// Выполнение команд из файла
	fmt.Printf("Выполнение команд из файла: %s\n", filename)
	executeCommandsFromFile(filename)
}