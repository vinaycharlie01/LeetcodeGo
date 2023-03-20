# import math

# class Shape:
#     def area(self):
#         pass

# class Rectangle(Shape):
#     def __init__(self, width, height):
#         self.width = width
#         self.height = height

#     def area(self):
#         return self.width * self.height

# class Circle(Shape):
#     def __init__(self, radius):
#         self.radius = radius

#     def area(self):
#         return math.pi * self.radius ** 2

# class AreaCalculator:
#     def __init__(self):
#         self.shapes = []

#     def add_shape(self, shape):
#         self.shapes.append(shape)

#     def total_area(self):
#         total = 0.0
#         for shape in self.shapes:
#             total += shape.area()
#         return total

# # example usage
# rectangle = Rectangle(10, 5)
# circle = Circle(7)
# calculator = AreaCalculator()
# calculator.add_shape(rectangle)
# calculator.add_shape(circle)
# print("Total area:", calculator.total_area())


class CartItem:
    def get_id(self):
        pass
    
    def get_title(self):
        pass
    
    def get_price(self):
        pass
    
    def set_title(self, title):
        pass
    
    def set_price(self, price):
        pass
    
class Book(CartItem):
    def __init__(self, id, title, author, price):
        self.id = id
        self.title = title
        self.author = author
        self.price = price
    
    def get_id(self):
        return self.id
    
    def get_title(self):
        return self.title
    
    def get_price(self):
        return self.price
    
    def set_title(self, title):
        self.title = title
    
    def set_price(self, price):
        self.price = price
    
class Electronics(CartItem):
    def __init__(self, id, title, brand, price):
        self.id = id
        self.title = title
        self.brand = brand
        self.price = price
    
    def get_id(self):
        return self.id
    
    def get_title(self):
        return self.title
    
    def get_price(self):
        return self.price
    
    def set_title(self, title):
        self.title = title
    
    def set_price(self, price):
        self.price = price
    
class Cart:
    def __init__(self):
        self.items = []
        
    def create(self, item):
        self.items.append(item)
        
    def read(self, id):
        for item in self.items:
            if item.get_id() == id:
                return item
        return None
    
    def update(self, id, title, price):
        for item in self.items:
            if item.get_id() == id:
                item.set_title(title)
                item.set_price(price)
                break
    
    def delete(self, id):
        for i, item in enumerate(self.items):
            if item.get_id() == id:
                del self.items[i]
                break
                
# Create a new cart
cart = Cart()

# Add some items to the cart
book = Book(1, "The Great Gatsby", "F. Scott Fitzgerald", 9.99)
cart.create(book)

electronics = Electronics(2, "Echo Dot", "Amazon", 49.99)
cart.create(electronics)

# Update the price of the book
cart.update(1, "The Great Gatsby (Paperback)", 12.99)

# Remove the Echo Dot from the cart
cart.delete(2)

# Print the remaining items in the cart
for item in cart.items:
    print(f"{item.get_title()}: ${item.get_price():.2f}")
