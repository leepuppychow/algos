from singly import Node, SinglyLinkedList

list1 = SinglyLinkedList()
list1.insert(1)
list1.insert(2)
list1.insert(3)
list1.insert(4)
list1.insert(5)

list1.append(6)

list1.delete(5)
list1.delete(1)

list1.print_list()
list1.reverse()
list1.print_list()