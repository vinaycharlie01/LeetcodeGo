# from abc import ABC, abstractmethod

# class Shape(ABC):
#     @abstractmethod
#     def area(self):
#         pass
#     @abstractmethod
#     def perimeter(self):
#         pass

# class Rectangle(Shape):
#     def __init__(self, width, height):
#         self.width = width
#         self.height = height

#     def area(self):
#         return self.width * self.height

#     def perimeter(self,a):
#         return a * (self.width + self.height)

# class Circle(Shape):
#     def __init__(self, radius):
#         self.radius = radius

#     def area(self):
#         return 3.14 * self.radius ** 2

#     def perimeter(self):
#         return 2 * 3.14 * self.radius

# rect = Rectangle(5, 10)
# print("Area of rectangle:", rect.area())
# print("Perimeter of rectangle:", rect.perimeter(10))

# a=[]

# circle = Circle(5)
# print("Area of circle:", circle.area())
# print("Perimeter of circle:", circle.perimeter())

# class Car:
#     def __init__(self, make, model):
#         self.__make = make
#         self.__model = model
    
#     def get_make(self):
#         return self.__make
    
#     def get_model(self):
#         return self.__model
    
#     def start(self):
#         print("Starting car...")

# # Client code
# my_car = Car("Toyota", "Camry")
# print("Make:", my_car.get_make())
# print("Model:", my_car.get_model())
# my_car.start()


from abc import ABC, abstractmethod


class Payment(ABC):
    @abstractmethod
    def pay(self):
        pass
class Cash(Payment):
    def  __init__(self,money):
        self.money=money
    def pay(self):
        return "Sucessfull your cash paymment"

class Onlinepay(Payment):
    def  __init__(self,money):
        self.money=money
    def pay(self):
        return "Sucessfull your Online paymment"
class USAdoller(Payment):
    def  __init__(self,money):
        self.money=money
    def pay(self):
        return "Sucessfull your USAdoller paymment"
a=Cash(10)
a.money=20
print(a.pay())