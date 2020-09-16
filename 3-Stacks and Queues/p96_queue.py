class Queue:
    def __init__(self):
        self.items = []

    def add(self):
        self.items.append(item)

    def remove(self):
        self.items = self.items[1:]

    def peek(self):
        return self.items[0]
    
    def isEmpty(self):
        return self.items == []

    
    def 