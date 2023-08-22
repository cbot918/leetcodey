class TreeNode:
  def __init__(self, x):
    self.val = x
    self.left = None
    self.right = None


root = None


def insert(val:int):
  if root == None:
    root = TreeNode(val)
    root.left = None
    root.right = None
    return
  cur = root
  
    
# def createTree():


if __name__ == '__main__':
  insert(3)