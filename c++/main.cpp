#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <sstream>
#include <algorithm>
#include <cctype>   
#include "Array/Array.h"
#include "ForwardList/SinglyLinkedList.h"
#include "LinkedList/DoublyLinkedList.h"
#include "Queue/Queue.h"
#include "Stack/Stack.h"
#include "Tree/Tree.h"

// Глобальные указатели на структуры данных
Array* globalArray = nullptr;
ForwardList* globalForwardList = nullptr;
LinkedList* globalLinkedList = nullptr;
Queue* globalQueue = nullptr;
Stack* globalStack = nullptr;
BST* globalTree = nullptr;


bool isInteger(const string& str) {
    if (str.empty()) return false;
    
    size_t start = 0;
    if (str[0] == '-') {
        if (str.length() == 1) return false;
        start = 1;
    }
    
    for (size_t i = start; i < str.length(); i++) {
        if (!isdigit(str[i])) return false;
    }
    return true;
}

// Функция для разделения строки на токены
vector<string> split(const string& str) {
    vector<string> tokens;
    stringstream ss(str);
    string token;
    while (ss >> token) {
        tokens.push_back(token);
    }
    return tokens;
}

void printHelpMessage() {
    cout << "  === ОПЕРАЦИИ ДОБАВЛЕНИЯ ===" << endl;
    cout << "  MPUSH <value>                  - добавить в массив" << endl;
    cout << "  MINSERT <index> <value>        - вставить в массив по индексу" << endl;
    cout << "  FPUSHHEAD <value>              - добавить в начало односвязного списка" << endl;
    cout << "  FPUSHTAIL <value>              - добавить в конец односвязного списка" << endl;
    cout << "  FINSERTBEFORE <target> <value> - вставить перед элементом" << endl;
    cout << "  FINSERTAFTER <target> <value>  - вставить после элемента" << endl;
    cout << "  LPUSHHEAD <value>              - добавить в начало двусвязного списка" << endl;
    cout << "  LPUSHTAIL <value>              - добавить в конец двусвязного списка" << endl;
    cout << "  LINSERTBEFORE <target> <value> - вставить перед элементом" << endl;
    cout << "  LINSERTAFTER <target> <value>  - вставить после элемента" << endl;
    cout << "  QPUSH <value>                  - добавить в очередь" << endl;
    cout << "  SPUSH <value>                  - добавить в стек" << endl;
    cout << "  TINSERT <value>                 - вставить в дерево (целое число)" << endl;
    cout << "" << endl;
    cout << "  === ОПЕРАЦИИ УДАЛЕНИЯ ===" << endl;
    cout << "  MDEL <index>                   - удалить из массива по индексу" << endl;
    cout << "  FDEL <value>                   - удалить из односвязного списка по значению" << endl;
    cout << "  FDELHEAD                       - удалить из головы односвязного списка" << endl;
    cout << "  FDELTAIL                       - удалить из хвоста односвязно списка" << endl;
    cout << "  FDELBEFORE <value>             - удалить из односвязного списка перед значением" << endl;
    cout << "  FDELAFTER <value>              - удалить из односвязного списка после значения" << endl; 
    cout << "  LDEL <value>                   - удалить из двусвязного списка по значению" << endl;
    cout << "  LDELHEAD                       - удалить из головы двусвязного списка" << endl;
    cout << "  LDELTAIL                       - удалить из хвоста двусвязного списка" << endl;
    cout << "  LDELBEFORE <value>             - удалить из двусвязного списка перед значением" << endl;
    cout << "  LDELAFTER <value>              - удалить из двусвязного списка после значения" << endl; 
    cout << "  QPOP                           - удалить из очереди (из начала)" << endl;
    cout << "  SPOP                           - удалить из стека (с вершины)" << endl;
    cout << "  TDEL <value>                    - удалить из дерева по значению (целое число)" << endl;
    cout << "" << endl;
    cout << "  === ОПЕРАЦИИ ЧТЕНИЯ ===" << endl;
    cout << "  MGET <index>                   - получить элемент массива по индексу" << endl;
    cout << "  TGET <value>                    - проверить наличие элемента в дереве (целое число)" << endl;
    cout << "  MLENGTH                        - длина массива" << endl;
    cout << "  FLENGTH                        - длина односвязного списка" << endl;
    cout << "  LLENGTH                        - длина двусвязного списка" << endl;
    cout << "  QLENGTH                        - длина очереди" << endl;
    cout << "  SLENGTH                        - длина стека" << endl;
    cout << "" << endl;
    cout << "  === ОПЕРАЦИИ ВЫВОДА ===" << endl;
    cout << "  PRINT [M/F/L/Q/S/T]            - вывести структуру" << endl;
    cout << "  PRINT_REVERSED [F/L]           - вывести список в обратном порядке" << endl;
    cout << "  CLEAR <M/F/L/Q/S/T>            - очистить структуру" << endl;
    cout << "----------------" << endl;
}

void executeCommand(const vector<string>& tokens) {
    if (tokens.empty()) return;
    
    string command = tokens[0];
    
    // Массив (M)
    if (command == "MPUSH" && tokens.size() > 1) {
        if (!globalArray) globalArray = CreateArray(8);
        Push(globalArray, tokens[1]);
    }
    else if (command == "MINSERT" && tokens.size() > 2) {
        if (!globalArray) globalArray = CreateArray(8);
        Insert(globalArray, stoi(tokens[1]), tokens[2]);
    }
    else if (command == "MREPLACE" && tokens.size() > 2) {
        if (globalArray) ReplaceByIndex(globalArray, stoi(tokens[1]), tokens[2]);
    }
    else if (command == "MDEL" && tokens.size() > 1) {
        if (globalArray) DeleteByIndex(globalArray, stoi(tokens[1]));
    }
    else if (command == "MGET" && tokens.size() > 1) {
        if (globalArray) {
            string value = GetByIndex(globalArray, stoi(tokens[1]));
            cout << "MGET [" << tokens[1] << "] = " << value << endl;
        }
    }
    else if (command == "MLENGTH") {
        if (globalArray) {
            cout << "Длина массива: " << GetLength(globalArray) << endl;
        }
    }
    
    // Односвязный список (F)
    else if (command == "FPUSHHEAD" && tokens.size() > 1) {
        if (!globalForwardList) globalForwardList = CreateForwardList();
        AddNode(Position::Head, CreateSingleNode(tokens[1]), globalForwardList, nullptr);
    }
    else if (command == "FPUSHTAIL" && tokens.size() > 1) {
        if (!globalForwardList) globalForwardList = CreateForwardList();
        AddNode(Position::Tail, CreateSingleNode(tokens[1]), globalForwardList, nullptr);
    }
    else if (command == "FINSERTAFTER" && tokens.size() > 2) {
        if (globalForwardList) {
            SingleNode* found = FindNodeByValue(globalForwardList, tokens[1]);
            if (found) AddNode(Position::After, CreateSingleNode(tokens[2]), globalForwardList, found);
        }
    }
    else if (command == "FINSERTBEFORE" && tokens.size() > 2) {
        if (globalForwardList) {
            SingleNode* found = FindNodeByValue(globalForwardList, tokens[1]);
            if (found) AddNode(Position::Before, CreateSingleNode(tokens[2]), globalForwardList, found);
        }
    }
    else if (command == "FDEL" && tokens.size() > 1) {
        if (globalForwardList) RemoveNodeByValue(globalForwardList, tokens[1]);
    }
    else if (command == "FDELHEAD") {
        if (globalForwardList) RemoveNode(Position::Head, globalForwardList, nullptr);
    }
    else if (command == "FDELTAIL") {
        if (globalForwardList) RemoveNode(Position::Tail, globalForwardList, nullptr);
    }
    else if (command == "FDELBEFORE" && tokens.size() > 1) {
        if (globalForwardList) RemoveNode(Position::Before, globalForwardList, FindNodeByValue(globalForwardList, tokens[1]));
    }
    else if (command == "FDELAFTER" && tokens.size() > 1) {
        if (globalForwardList) RemoveNode(Position::After, globalForwardList, FindNodeByValue(globalForwardList, tokens[1]));
    }
    else if (command == "FLENGTH") {
        if (globalForwardList) {
            int count = 0;
            SingleNode* curr = globalForwardList->head;
            while (curr != nullptr) {
                count++;
                curr = curr->next;
            }
            cout << "Длина односвязного списка: " << count << endl;
        }
    }
    
    // Двусвязный список (L)
    else if (command == "LPUSHHEAD" && tokens.size() > 1) {
        if (!globalLinkedList) globalLinkedList = CreateLinkedList();
        AddNode(Position::Head, CreateDoubleNode(tokens[1]), globalLinkedList, nullptr);
    }
    else if (command == "LPUSHTAIL" && tokens.size() > 1) {
        if (!globalLinkedList) globalLinkedList = CreateLinkedList();
        AddNode(Position::Tail, CreateDoubleNode(tokens[1]), globalLinkedList, nullptr);
    }
    else if (command == "LINSERTAFTER" && tokens.size() > 2) {
        if (globalLinkedList) {
            DoubleNode* found = FindNodeByValue(globalLinkedList, tokens[1]);
            if (found) AddNode(Position::After, CreateDoubleNode(tokens[2]), globalLinkedList, found);
        }
    }
    else if (command == "LINSERTBEFORE" && tokens.size() > 2) {
        if (globalLinkedList) {
            DoubleNode* found = FindNodeByValue(globalLinkedList, tokens[1]);
            if (found) AddNode(Position::Before, CreateDoubleNode(tokens[2]), globalLinkedList, found);
        }
    }
    else if (command == "LDEL" && tokens.size() > 1) {
        if (globalLinkedList) RemoveNodeByValue(globalLinkedList, tokens[1]);
    }
    else if (command == "LDELHEAD") {
        if (globalLinkedList) RemoveNode(Position::Head, globalLinkedList, nullptr);
    }
    else if (command == "LDELTAIL") {
        if (globalLinkedList) RemoveNode(Position::Tail, globalLinkedList, nullptr);
    }
    else if (command == "LDELBEFORE" && tokens.size() > 1) {
        if (globalLinkedList) RemoveNode(Position::Before, globalLinkedList, FindNodeByValue(globalLinkedList, tokens[1]));
    }
    else if (command == "LDELAFTER" && tokens.size() > 1) {
        if (globalLinkedList) RemoveNode(Position::After, globalLinkedList, FindNodeByValue(globalLinkedList, tokens[1]));
    }
    else if (command == "LLENGTH") {
        if (globalLinkedList) {
            int count = 0;
            DoubleNode* curr = globalLinkedList->head;
            while (curr != nullptr) {
                count++;
                curr = curr->next;
            }
            cout << "Длина двусвязного списка: " << count << endl;
        }
    }
    
    // Очередь (Q)
    else if (command == "QPUSH" && tokens.size() > 1) {
        if (!globalQueue) globalQueue = CreateQueue();
        Push(globalQueue, CreateQueueElement(tokens[1]));
    }
    else if (command == "QPOP") {
        if (globalQueue && !isEmpty(globalQueue)) {
            string popedValue = Pop(globalQueue);
            cout << "Вырезанное значение из очереди: " << popedValue << endl;
        }
    }
    else if (command == "QLENGTH") {
        if (globalQueue) {
            int count = 0;
            QueueElement* curr = globalQueue->first;
            while (curr != nullptr) {
                count++;
                curr = curr->next;
            }
            cout << "Длина очереди: " << count << endl;
        }
    }
    
    // Стек (S)
    else if (command == "SPUSH" && tokens.size() > 1) {
        if (!globalStack) globalStack = CreateStack();
        Push(globalStack, CreateStackElement(tokens[1]));
    }
    else if (command == "SPOP") {
        if (globalStack && !isEmpty(globalStack)) {
            string popedValue = Pop(globalStack);
            cout << "Вырезанное значение из стека: " << popedValue << endl;
        }
    }
    else if (command == "SLENGTH") {
        if (globalStack) {
            int count = 0;
            StackElement* curr = globalStack->last;
            while (curr != nullptr) {
                count++;
                curr = curr->prev;
            }
            cout << "Длина стека: " << count << endl;
        }
    }
    
    // Дерево (T) 
    else if (command == "TINSERT" && tokens.size() > 1) {
        if (!isInteger(tokens[1])) {
            cout << "ОШИБКА: введите число " << endl;
            return;
        }
        if (!globalTree) globalTree = CreateBST();
        InsertNode(globalTree, stoi(tokens[1]));
    }
    else if (command == "TDEL" && tokens.size() > 1) {
        if (!isInteger(tokens[1])) {
            cout << "ОШИБКА: введите число " << endl;
            return;
        }
        if (globalTree) RemoveNodeByValue(globalTree, stoi(tokens[1]));
    }
    else if (command == "TGET" && tokens.size() > 1) {
        if (!isInteger(tokens[1])) {
            cout << "ОШИБКА: введите число " << endl;
            return;
        }
        if (globalTree) {
            BSTNode* found = FindNode(globalTree->root, stoi(tokens[1]));
            if (found) cout << "TGET " << tokens[1] << " = found" << endl;
            else cout << "TGET " << tokens[1] << " = not found" << endl;
        }
    }
    
    // Операции вывода
    else if (command == "PRINT") {
        if (tokens.size() > 1) {
            string structure = tokens[1];
            if (structure == "M" && globalArray) {
                cout << "Array: ";
                Print(globalArray);
            }
            else if (structure == "F" && globalForwardList) {
                cout << "Forward List: ";
                Print(globalForwardList);
            }
            else if (structure == "L" && globalLinkedList) {
                cout << "Linked List: ";
                Print(globalLinkedList);
            }
            else if (structure == "Q" && globalQueue) {
                cout << "Queue: ";
                Print(globalQueue);
            }
            else if (structure == "S" && globalStack) {
                cout << "Stack: ";
                Print(globalStack);
            }
            else if (structure == "T" && globalTree) {
                cout << "Tree : ";
                PrintInorder(globalTree);
            }
        }
        else {
            if (globalArray) {
                cout << "Array: ";
                Print(globalArray);
            }
            if (globalForwardList) {
                cout << "Forward List: ";
                Print(globalForwardList);
            }
            if (globalLinkedList) {
                cout << "Linked List: ";
                Print(globalLinkedList);
            }
            if (globalQueue) {
                cout << "Queue: ";
                Print(globalQueue);
            }
            if (globalStack) {
                cout << "Stack: ";
                Print(globalStack);
            }
            if (globalTree) {
                cout << "Tree : ";
                PrintInorder(globalTree);
            }
        }
    }
    else if (command == "PRINT_REVERSED" && tokens.size() > 1) {
        string structure = tokens[1];
        if (structure == "F" && globalForwardList) {
            cout << "Forward List (reversed): ";
            PrintReversed(globalForwardList);
        }
        else if (structure == "L" && globalLinkedList) {
            cout << "Linked List (reversed): ";
            PrintReversed(globalLinkedList);
        }
    }
    else if (command == "CLEAR") {
        if (tokens.size() > 1) {
            string structure = tokens[1];
            if (structure == "M" && globalArray) {
                Destroy(globalArray);
                globalArray = nullptr;
            }
            else if (structure == "F" && globalForwardList) {
                Destroy(globalForwardList);
                globalForwardList = nullptr;
            }
            else if (structure == "L" && globalLinkedList) {
                Destroy(globalLinkedList);
                globalLinkedList = nullptr;
            }
            else if (structure == "Q" && globalQueue) {
                ClearQueue(globalQueue);
                globalQueue = nullptr;
            }
            else if (structure == "S" && globalStack) {
                ClearStack(globalStack);
                globalStack = nullptr;
            }
            else if (structure == "T" && globalTree) {
                Destroy(globalTree);
                globalTree = nullptr;
            }
        }
    }
    else {
        cout << "Неизвестная команда или параметр: " << command << endl;
    }
}

void executeCommandsFromFile(const string& filename) {
    ifstream file(filename);
    if (!file.is_open()) {
        cout << "Не удалось открыть файл: " << filename << endl;
        return;
    }
    
    string line;
    int lineNumber = 0;
    while (getline(file, line)) {
        lineNumber++;
        if (!line.empty()) {
            
            vector<string> tokens = split(line);
            executeCommand(tokens);
        }
    }
    file.close();
}

void appendCommandToFile(const string& filename, const string& command) {
    ofstream file(filename, ios::app);
    if (file.is_open()) {
        file << command << endl;
        file.close();
    }
}

void removeLastCommandFromFile(const string& filename) {
    ifstream file(filename);
    if (!file.is_open()) return;
    
    vector<string> lines;
    string line;
    while (getline(file, line)) {
        if (!line.empty()) {
            lines.push_back(line);
        }
    }
    file.close();
    
    if (!lines.empty()) {
        ofstream outFile(filename);
        for (size_t i = 0; i < lines.size() - 1; ++i) {
            outFile << lines[i] << endl;
        }
        outFile.close();
    }
}

int main(int argc, char* argv[]) {
    string filename;
    string query;
    bool undo = false;
    
    for (int i = 1; i < argc; ++i) {
        string arg = argv[i];
        if (arg == "--help" || arg == "-h") {
            printHelpMessage();
        }
        if (arg == "--file" && i + 1 < argc) {
            filename = argv[++i];
        }
        else if (arg == "--query" && i + 1 < argc) {
            query = argv[++i];
        }
        else if (arg == "--undo") {
            undo = true;
        }
    }
    
    if (filename.empty()) {
        cout << "Использование: " << argv[0] << " --file <filename> [--query <command>] [--undo]" << endl;
        return 1;
    }
    
    if (undo) {
        removeLastCommandFromFile(filename);
        cout << "Последняя команда из файла " << filename << " отменена" << endl;
    }
    else if (!query.empty()) {
        appendCommandToFile(filename, query);
        cout << "В файл добавлена команда: " << query << endl;
    }
    
    executeCommandsFromFile(filename);
    
    if (globalArray) Destroy(globalArray);
    if (globalForwardList) Destroy(globalForwardList);
    if (globalLinkedList) Destroy(globalLinkedList);
    if (globalQueue) ClearQueue(globalQueue);
    if (globalStack) ClearStack(globalStack);
    if (globalTree) Destroy(globalTree);
    
    return 0;
}