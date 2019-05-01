import unittest
from binary import Node, BinarySearchTree

class TestBinarySearchTree(unittest.TestCase):
  def setUp(self):
    self.root = Node(50)
    self.tree = BinarySearchTree(self.root)
    self.tree.insert(40, self.root)
    self.tree.insert(60, self.root)
    self.tree.insert(45, self.root)
    self.tree.insert(65, self.root)
    self.tree.insert(20, self.root)

  def test_insert(self):
    self.assertEqual(self.tree.root.data, 50)
    self.assertEqual(self.tree.root.left.data, 40)
    self.assertEqual(self.tree.root.left.left.data, 20)
    self.assertEqual(self.tree.root.left.right.data, 45)

    self.assertEqual(self.tree.root.right.data, 60)
    self.assertEqual(self.tree.root.right.right.data, 65)
    self.assertEqual(self.tree.root.right.left, None)

  def test_search(self):
    self.assertTrue(self.tree.search(50))
    self.assertTrue(self.tree.search(40))
    self.assertTrue(self.tree.search(60))
    self.assertTrue(self.tree.search(45))
    self.assertTrue(self.tree.search(65))
    self.assertTrue(self.tree.search(20))

    self.assertFalse(self.tree.search(100))
    self.assertFalse(self.tree.search(0))
  
if __name__ == '__main__':
    unittest.main()