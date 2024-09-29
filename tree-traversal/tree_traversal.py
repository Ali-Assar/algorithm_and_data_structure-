from collections import deque

class Node:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

class BinarySearchTree:
    def __init__(self):
        self.root = None

    def insert(self, value):
        new_node = Node(value)
        if self.root is None:
            self.root = new_node
            return True
        temp = self.root
        while True:
            if new_node.value == temp.value:
                return False
            if new_node.value < temp.value:
                if temp.left is None:
                    temp.left = new_node
                    return True
                temp = temp.left
            else:
                if temp.right is None:
                    temp.right = new_node
                    return True
                temp = temp.right

    def BFS(self):
        results = []
        if self.root is None:
            return results
        queue = deque([self.root])

        while queue:
            current_node = queue.popleft()
            results.append(current_node.value)
            if current_node.left:
                queue.append(current_node.left)
            if current_node.right:
                queue.append(current_node.right)
        return results
    
    def dfs_pre_order(self):
        results = []
        def traverse(node):
            results.append(node.value)
            if node.left:
                traverse(node.left)
            if node.right:
                traverse(node.right)
        traverse(self.root)
        return results

    def dfs_post_order(self):
        results = []
        def traverse(node):
            if node.left:
                traverse(node.left)
            if node.right:
                traverse(node.right)
            results.append(node.value)
        traverse(self.root)
        return results

    def dfs_in_order(self):
        results = []
        def traverse(node):
            if node.left:
                traverse(node.left)
            results.append(node.value)
            if node.right:
                traverse(node.right)
        traverse(self.root)
        return results

tree = BinarySearchTree()
for value in [28, 45, 98, 11, 21, 86, 21]:
    tree.insert(value)

print("BFS:", tree.BFS())
print("DFS Pre-order:", tree.dfs_pre_order())
print("DFS Post-order:", tree.dfs_post_order())
print("DFS In-order:", tree.dfs_in_order())
