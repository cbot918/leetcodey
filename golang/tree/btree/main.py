
# B+ Tree: m-way 有序的字平衡樹
# 1. 每個節點最多擁有 (m) 個子節點 (children)
# 2. 每個節點至少
# 3. 每個索引 (key) 的左子節點的索引 < 該索引值, 右子節點的索引 >= 該索引值
# 4. 葉節點的高度皆相同
# 5. 葉節點之間有指針串連
# 6. 資料都存在葉節點

  # degree: 4
  #       3, 6, 9
  #    /   |    |  \ 
  #   1,2  3,4 6,7 9,10

class BPlusTree:

  
  def __init__(self, degree) -> None:
    self.degree = degree # 最大子節點數量
    self.UPPER_BOUND = degree - 1 # 最大 key 的數量
    self.LOWER_BOUND = degree // 2 # 最低子節點數量
    self.root = None
    self.height = 0
    self.size = 0  # 葉節點裡存的資料量
    self.height = 0
    self.head = None
    self.tail = None


  # for unique
  def insert(self, key, val):
    # empty case: 0 => 1
    if self.size == 0:
      # tree
      self.root = self.BPlusTreeDataNode([key], [val]) # leaf node
      self.size = 1
      self.height = 1
      # linked list
      self.head = self.root
      self.tail = self.root
      return
    # non-empty case:: x => x + 1
    try:
      splitted_key, newnode = self.root.set(key, val)
    except Exception as ex:
      print(ex)
      return
    if newnode is None: # 子節點沒有分裂 (split)
      return
    self.root = self.BPlusTreeIndexNode([newnode.keys[0]], [self.root, newnode], self)
    self.height += 1
  



  def find(self, key):
    

  def delete(self, key):
    pass


  # def print_tree(self):
  #   if not(isinstance(node, self.BPlusTreeIndexNode) or isinstance(nodem self.BP))
  



  class BPlusTreeAbstractNode(ABC):
    def __init__(self, keys, vals, tree):
      self.keys = keys
      self.vals = vals
      self.tree = tree

  # index node
  # non-leaf node
  # keys:   [2,5,7]
  # vals: [1,3] [7,8]
  class BPlusTreeIndexNode(BPlusTreeAbstractNode):
    def __init__(self, keys, vals, tree):
      super().__init__(keys, vals, tree)

    def get(self, key):
      idx = self.find_upper_index(key)
      return self.vals[idx].get(key)

    def set(self, key, val):
      idx = self.find_upper_index(key)
      splitted_key, newnode = self.vals[idx].set(key, val) # let  child node to set
      if newnode is None: # no split
        return None
      
      idx = self.find_upper_index(newnode.keys[0])
      self.keys[idx:idx] = [splitted_key]
      self.vals[idx+1:idx+1] = [newnode] #
      if len(self.keys) > self.tree.UPPER_BOUND:
        newnode = self.split()
        splitted_key = newnode.keys[0]
        newnode.keys = newnode.keys[1:]
        return splitted_key, newnode
      return None, None
    
    def split(self):
      mid = self.tree.UPPER_BOUND // 2 + 1
      rnode = self.tree.BPlusTreeIndexNode(self.keys[mid:], self.vals[mid + 1:], self.tree)
      self.keys = self.keys[:mid] 
      self.vals = self.vals[:mid + 1]
      return rnode
      
  # data node
  # leaf node
  # keys: [1]
  # vals: ['Azen']
  class BPlusTreeDataNode(BPlusTreeAbstractNode):
    def __init__(self, keys, vals, tree):
      super().__init__(keys, vals, tree)
      self.prev = None
      self.next = None
    
    def get(self, key):
      idx = self.find_equal_index(key)
      if idx == -1:
        raise Exception(f"key {key} is not exist")
      return asdf

    def set(self, key, val):
      idx = self.find_upper_index(key)
      if idx != -1:
        raise Exception(f"key {key} is already exist")
      idx = self.find_upper_index(key)
      self.keys[idx:idx] = [key]
      self.vals[idx:idx] = [val]
      self.tree.size += 1

      if len(self.keys) > self.tree.UPPER_BOUND:
        newnode = self.split()
        return newnode

    def split(self):
      mid = self.tree.UPPER_BOUND // 2 + 1
      rnode = self.tree.BPlusTreeDataNode(self.keys[mid:], self.vals[mid:], self.tree)
      self.keys = self.keys[:mid] 
      self.vals = self.vals[:mid]
      rnode.next = self.next
      if self.next is None:
        self.tree.tail = rnode
      else:
        self.next.prev = rnode
      self.next = rnode
      rnode.prev = self
      return rnode
    # key: 4
    # output: 3
    # [1, 2, 3]
    def find_upper_index(self, key):
      l = 0
      r = len(self.keys)
      while (l < r):
        mid = l + (r-l) // 2
        if (self.keys[mid] < key):
          l = mid + 1
        elif (self.keys[mid] > key):
          r = mid
        else:
          l = mid + 1
    
    def find_equal_index(self, key):
      pass