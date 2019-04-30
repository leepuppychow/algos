class Node:
  def __init__(self, data):
    self.data = data
    self.child = None

class SinglyLinkedList:
  def __init__(self):
    self.head = None

  def insert(self, data):
    new_node = Node(data)
    new_node.child = self.head
    self.head = new_node

  def append(self, data):
    current = self.head
    while current.child is not None:
      current = current.child
    current.child = Node(data)

  def includes(self, data):
    current = self.head
    while current.data != data:
      if current.child is None:
        return False
      current = current.child
    return True

  def delete(self, data):
    current = self.head
    if current.data == data:
      head = current
      self.head = current.child
      head.child = None
      return
    while current.child is not None:
      if current.child.data == data:
        to_be_deleted = current.child
        current.child = current.child.child
        to_be_deleted.child = None
        return
      else: 
        current = current.child
    print(data, "not found")
  
  def reverse(self):
    prev = None
    current = self.head

    while current != None:
      next = current.child
      current.child = prev
      prev = current
      current = next
    self.head = prev
    
  def print_list(self):
    current = self.head
    while current is not None:
      print(current.data)
      current = current.child