class Node:
  def __init__(self, data, child):
    self.data = data
    self.child = child

  def insert(self, new_node):
    new_node.child = self
    return new_node

  def append(self, new_node):
    if self.child == None:
      self.child = new_node
    else:
      self.child.append(new_node)

  def includes(self, data):
    if self.data == data:
      return True
    elif self.child == None and self.data != data:
      return False
    else:
      return self.child.includes(data)

  def delete(self, data):
    if self.data == data:
      new_head = self.child
      self.child = None
      return new_head
    elif self.child == None:
      print(data, "not found")
      return
    elif self.child.data == data:
      to_be_deleted = self.child
      self.child = self.child.child
      to_be_deleted.child = None
      return
    else:
      return self.child.delete(data)
  
  def print_list(self):
    if self.child == None:
      print(self.data)
      return
    else:
      print(self.data)
      self.child.print_list()