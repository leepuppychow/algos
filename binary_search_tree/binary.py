class Node:
  def __init__(self, data):
    self.data = data
    self.left = None
    self.right = None

class BinarySearchTree:
  def __init__(self, root):
    self.root = root
  
  def insert(self, data, current):
    if data < current.data:
      if current.left is None:
        current.left = Node(data)
        return
      else:
        self.insert(data, current.left)
    else:
      if current.right is None:
        current.right = Node(data)
        return
      else: 
        self.insert(data, current.right)

    