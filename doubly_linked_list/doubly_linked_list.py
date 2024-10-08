class Node:
    def __init__(self,value):
        self.value = value
        self.next = None
        self.prev = None

class DoublyLinkedList:
    def __init__ (self ,value):
        new_node = Node(value)
        self.head = new_node
        self.tail = new_node
        self.length = 1

    def print_list(self):
        temp=self.head
        while temp is not None:
            print(temp.value)
            temp = temp.next
    
    def append(self, value):
        new_node = Node(value)
        if self.head is None:
            self.head = new_node
            self.tail = new_node
        else:
            self.tail.next = new_node
            new_node.prev = self.tail
            self.tail = new_node
        self.length +=1
        return True
    
    def pop(self):
        if self.length == 0:
            return None
        temp = self.tail
        self.tail = self.tail.prev
        self.tail.next = None
        temp.prev = None
        self.length -=1
    
    def prepend(self,value):
        new_node = Node(value)
        if self.head is None:
            self.head = new_node
            self.tail = new_node
        else:
            new_node.next = self.head
            self.head.prev = new_node
            self.head = new_node
        self.length +=1
    
    def pop_first(self):
        if self.length == 0:
            return None
        temp = self.head
        if self.length == 1:
            self.head = None
            self.tail = None
        else:
            self.head = self.head.next
            self.head.prev = None
            temp.next = None
            self.length -= 1
        return temp
    
    def get(self,index):
        if index < 0 or index >= self.length:
            return None
        temp = self.head       
        if index < self.length/2:
            for _ in range(index):
                temp= temp.next
        else:
            temp = self.tail
            for _ in range(self.length -1, index,-1):
                temp= temp.prev
        return temp
    
    def set_value(self,index,value):
         temp = self.get(index)
         if temp is not None:
             temp.value = value
             return True
         else:
            return False
         
    def insert(self,index,value):
        if index <0 or index >self.length:
            return False
        if index == 0:
            return self.prepend(value)
        if index == self.length:
            return self.append(value)
        new_node = Node(value)
        before = self.get(index -1)
        after = before.next

        new_node.prev = before
        new_node.next = after
        before.next = new_node
        after.prev = new_node    
        
        self.length += 1
        return True
    
    def remove(self,index):
        if index <0 or index >self.length:
            return False
        if index == 0:
            return self.pop()
        if index == self.length:
            return self.pop_first()
        
        temp = self.get(index)
        temp.next.prev = temp.prev
        temp.prev.next = temp.next
        temp.next = None
        temp.prev = None

        self.length -= 1
        return temp
        

double_link_list = DoublyLinkedList(7)
double_link_list.append(5)
double_link_list.append(6)
double_link_list.append(4)
double_link_list.append(3)
double_link_list.pop()
double_link_list.prepend(1)
double_link_list.pop_first()
double_link_list .print_list()
print(double_link_list.get(1))
print(double_link_list.set_value(1,22))
print(double_link_list.insert(4,33))
print(double_link_list.remove(2))