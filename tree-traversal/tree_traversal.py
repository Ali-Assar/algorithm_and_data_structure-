class Node:
    def __init__(self,value) :
        self.value =  value
        self.left = None
        self.right = None


class BinarySearchTree:
    def __init__(self):
        self.root = None

    def insert(self,value):
        new_node = Node(value)
        if self.root is None:
            self.root = new_node
            return True
        else:
            temp = self.root
            while (True):
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
        current_node = self.root
        queue = []
        results = []
        queue.append(current_node )

        while len (queue) > 0:
            current_node = queue.pop(0)
            results.append(current_node.value)
            if current_node.left is not None:
                queue.append(current_node.left)
            if current_node.right is not None:
                queue.append(current_node.right)
        return results
    
    def dfs_pre_order (self):
        results = []
        def traverse (current_node):
            results.append(current_node.value)
            if current_node.left is not None:
                traverse(current_node.left)
            if current_node.right is not None:
                traverse(current_node.right)
        traverse(self.root)
        return results
    
    def dfs_post_order(self):
        results = []    
        def traverse (current_node):
            
            if current_node.left is not None:
                traverse(current_node.left)
            if current_node.right is not None:
                traverse(current_node.right)
            results.append(current_node.value)
        traverse(self.root)
        return results
    
    def dfs_in_order(self):
        results = []    
        def traverse (current_node):
            
            if current_node.left is not None:
                traverse(current_node.left)
            results.append(current_node.value)
            if current_node.right is not None:
                traverse(current_node.right)
           
        traverse(self.root)
        return results

tree = BinarySearchTree()
tree.insert(28)
tree.insert(45)
tree.insert(98)
tree.insert(11)
tree.insert(21)
tree.insert(86)
tree.insert(21)

print("BFS result:", tree.BFS())



print("dfs pre order result:",tree.dfs_pre_order())

print("dfs post order result:",tree.dfs_post_order())

print("dfs in order result:",tree.dfs_in_order())