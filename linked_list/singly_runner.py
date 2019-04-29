from singly import Node

node1 = Node('node1', None)
node2 = Node('node2', None)
node3 = Node('node3', None)
node4 = Node('node4', None)
node5 = Node('node5', None)
node6 = Node('node6', None)

node1.insert(node2)
node2.insert(node3)
node3.append(node4)
node3.append(node5)
node3.append(node6)

node3.delete('node1')

node3.print_list()