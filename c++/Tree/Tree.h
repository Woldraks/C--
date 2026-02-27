#pragma once
using namespace std; 

// Узел бинарного дерева поиска (BST)
struct BSTNode {
    int key;                 // ИЗМЕНЕНО: теперь int вместо string
    BSTNode* left;
    BSTNode* right;
};

// Структура BST, хранит указатель на корень дерева
struct BST {
    BSTNode* root;
};


BSTNode* CreateBSTNode(int key);           

BST* CreateBST();

bool isEmpty(BST* tree);

void InsertNode(BST* tree, int key);       

void inorderPrint(BSTNode* node);

void PrintInorder(BST* tree);

BSTNode* findMinNode(BSTNode* node);

BSTNode* FindNode(BSTNode* root, int key); 

BSTNode* findAndRemoveNode(BSTNode* root, int key);

void RemoveNodeByValue(BST* tree, int key); 

void clearBSTRec(BSTNode* node);

void Destroy(BST* tree);