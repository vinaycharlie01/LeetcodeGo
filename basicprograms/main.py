# def SivenAlog(min1,max1):
#     siven=[True for i in range(max1+1)]
#     i=2
#     while i*i<=max1:
#         if siven[i]:
#             for j in range(i*i,max1+1,i):
#                 siven[j]=False
#         i+=1
    
    # prime=[i for i in range(min1,max1+1) if result[i]]
#     print(prime)
# SivenAlog(10,20)
class Person:
    def __init__(self,name):
        self.name = name
    def add(self,a,b):
        return a+b
    def add(self,a,b,c):
        return a+b

class Person2(Person):
    def __init__(self,name,age):
        self.age=age
        super().__init__(name)
    def add(self,a,b):
        return a+b


b=Person("vinau")
print(b.add(10,20,30))
c=Person2("rhr",20)
print(c.add(20,20))
        