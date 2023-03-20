class Person:
    def __init__(self,a):
        self.a=a
    # def __init__(self,a,b):
    #     self.a = a
    #     self.b = b
    def a1(self):
        print(self)

p1=Person(10)
print(p1)
print(p1.a1())