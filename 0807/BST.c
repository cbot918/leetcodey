#include<stdio.h>
#include<stdlib.h>
#include<stdbool.h>

struct TreeNode {
  int val;
  struct TreeNode *left;
  struct TreeNode *right;
};

/* Medium 1008. Construct Binary Search Tree from Preorder Traversal

Given an array of integers preorder, which represents the preorder traversal of a BST (i.e., binary search tree), construct the tree and return its root.

It is guaranteed that there is always possible to find a binary search tree with the given requirements for the given test cases.

A binary search tree is a binary tree where for every node, any descendant of Node.left has a value strictly less than Node.val, and any descendant of Node.right has a value strictly greater than Node.val.

A preorder traversal of a binary tree displays the value of the node first, then traverses Node.left, then traverses Node.right.

Input: preorder = [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]

Input: preorder = [1,3]
Output: [1,null,3]

*/
struct TreeNode* bstFromPreorder(int data[], int len){
  struct TreeNode *root = malloc(sizeof(struct TreeNode));

  if(len == 0)
    return NULL;
  if(len == 1){
    root->val = data[0];
    root->left = root->right = NULL;
    return root;
  }
  if(len > 1){
    root->val = data[0];
    root->left = root->right = NULL;
    for(int i=1; i<len; i++){
      int flag = 0;
      struct TreeNode* n = root;
      while(flag == 0){
        if(data[i] < n->val){
          if(n->left != NULL) n =  n->left;
          else {
            struct TreeNode* g = malloc(sizeof(struct TreeNode));
            g->val = data[i];
            g->left = g->right = NULL;
            n->left = g;
            flag ++;
          }
        } else {
          if(n->right != NULL) n = n-> right;
          else {
            struct TreeNode* g = malloc(sizeof(struct TreeNode));
            g->val = data[i];
            g->left = g->right = NULL;
            n->right = g;
            flag ++;
          }
        }
      }

    }
  }
  return root;
}

/* Easy 226. Invert Binary Tree

 Given the root of a binary tree, invert the tree, and return its root.

  Input: root = [4,2,7,1,3,6,9]
  Output: [4,7,2,9,6,3,1]

  Input: root = [2,1,3]
  Output: [2,3,1]

  Input: root = []
  Output: []
*/
struct TreeNode* invertTree(struct TreeNode* root){
  
  if (root == NULL){
    return root;
  }
  
  invertTree(root->left);
  invertTree(root->right);
  
  struct TreeNode* cur = NULL;
  cur = root->left;
  root->left = root->right;
  root->right = cur;

  return root;
  
}

/* Easy 104. Maximum Depth of Binary Tree

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Input: root = [3,9,20,null,null,15,7]
Output: 3

Input: root = [1,null,2]
Output: 2

*/
int maxDepth(struct TreeNode* root){
  if(!root) return 0;
  int lh = maxDepth(root->left);
  int rh = maxDepth(root->right);
  if(lh>rh) return lh+1;
  else      return rh+1;
}

/* Easy 100. Same Tree 

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Input: p = [1,2,3], q = [1,2,3]
Output: true

Input: p = [1,2], q = [1,null,2]
Output: false

Input: p = [1,2], q = [1,null,2]
Output: false
*/
bool isSameTree(struct TreeNode* p, struct TreeNode* q){
    if(p==NULL && q==NULL) return true;
    if(p==NULL || q==NULL) return false;
    if(p->val != q->val) return false;
    return isSameTree(p->left,q->left) && isSameTree(p->right, q->right);
}

/* Easy 572. Subtree of Another Tree
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.

Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true

Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
*/
bool isSubtree(struct TreeNode* root, struct TreeNode* subRoot) {
    if(!root ) return false;
    if(isSameTree(root, subRoot)) return true;
    return isSubtree(root->left, subRoot) || isSubtree(root->right, subRoot);  
}



void inorder(struct TreeNode* current){
  if(current != NULL){
    inorder(current->left);
    printf("%d ",current->val);
    inorder(current->right);
  }
}

int main(){
  int data[] = {4,2,7,1,3,6,9};
  int len = (int)sizeof(data)/sizeof(data[0]);

  // Medium 1008. Construct Binary Search Tree from Preorder Traversal
  struct TreeNode *root = bstFromPreorder(data, len);

  // Easy 226. Invert Binary Tree
  // invertTree(root);

  // Easy 104. Maximum Depth of Binary Tree
  // printf("%d \n", maxDepth(root));

  // Easy 100. Same Tree
  // int t1[] = {1,2,3}; struct TreeNode *p = bstFromPreorder(t1, 3);
  // int t2[] = {1,2,3}; struct TreeNode *q = bstFromPreorder(t2, 3);
  // printf("isSameTree: %s\n",isSameTree(p,q)? "true":"false");

  // Easy 572. Subtree of Another Tree (broken answer in local, dont know why)
  int tree[] = {3,4,5,1,2}; struct TreeNode *r = bstFromPreorder(tree, 5);
  int subTree[] = {4,1,2}; struct TreeNode *s = bstFromPreorder(subTree, 3);
  printf("isSubtree: %s\n", isSubtree(r,s)? "true":"false");

  // other
  // inorder(root);

}