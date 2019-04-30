import unittest
from singly import Node, SinglyLinkedList

class TestSinglyLinkedList(unittest.TestCase):
  def setUp(self):
    self.list = SinglyLinkedList()
    self.list.insert(1)
    self.list.insert(2)
    self.list.insert(3)
    self.list.insert(4)
    self.list.insert(5)
    self.list.append(6)

  def test_insert(self):
    self.list.insert(100)
    self.assertEqual(self.list.head.data, 100)
    self.assertNotEqual(self.list.head.data, 5)

  def test_includes(self):
    self.assertEqual(self.list.includes(3), True)
    self.assertEqual(self.list.includes(5), True)
    self.assertEqual(self.list.includes(6), True)
    
    self.assertEqual(self.list.includes(100), False)
    self.assertEqual(self.list.includes('hello'), False)

  def test_delete(self):
    self.assertEqual(self.list.includes(3), True)
    self.list.delete(3)
    self.assertEqual(self.list.includes(3), False)

    # Deleting the head
    self.list.delete(5)
    self.assertNotEqual(self.list.head.data, 5)
    self.assertEqual(self.list.head.data, 4)

    x = self.list.delete(10000)
    self.assertEqual(x, "not found")
  
if __name__ == '__main__':
    unittest.main()