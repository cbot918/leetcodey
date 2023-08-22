#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
struct TreeNode{
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

typedef struct TreeNode Node;

Node *root, *current = NULL;

void insert(int val){
  Node *prev = NULL;
  Node *p = (Node*)malloc(sizeof(Node));
  p->val = val;
  p->left = p->right = NULL;

  current = root;
  if(current == NULL){
    root = p;
    return;
  }
  
  while(current != NULL){
    prev = current;
    if(p->val < current->val){
      current = current->left;
    } else {
      current = current->right;
    }
  }
  
  if (p->val < prev->val){
    prev->left = p;
  } else {
    prev->right = p;
  }

}
void createTree(int data[], int len){
  for(int i=0; i<len; i++){
    insert(data[i]);
  }
  
}
void preorder(Node *cur){
  if(cur){
    printf("%d ", cur->val);  
    preorder(cur->left);
    preorder(cur->right);
  }
}
Node* search(Node *cur, int val){
  
  while(cur != NULL){
    if(cur->val == val) return cur;
    if(cur->val > val) cur = cur->left;
    else cur = cur->right;
  }
  return NULL;
}

// 1. delete leaf
// 2. delete single-leaf node
// 3. delete two-leaf node
void delete(Node *root, int val){
  
  
}

int main(){
  // int data[] = {4,2,1,7};
  // int data[] = {5,4,3,2,1};
  // int data[] = {6,4,2,1,5,8,7,9,10};
  int data[] = {29,13,6,4,10,12,11,19,35,38};
  // int data[] = {5,4,3,2,1};
  int len = (int)sizeof(data)/sizeof(data[0]);
  
  // insert
  createTree(data, len);
  
  // traverse
  // preorder(root);

  // search
  // Node *s = search(root, 10);
  // if(s) printf("%d", s->val);
  // else printf("not found");

  // delete(root, 2);
  preorder(root);

  // delete


}