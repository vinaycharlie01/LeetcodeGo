# class Book:
#     def __init__(self,id,title,author):
#         print(self)
#         self.id=id
#         self.title=title
#         self.author=author
#     def hello(self):
#         print("Hello world")
# b=Book(1,"Go","Goole")
# print(b.hello())


# def Hello(self):
#     return "Hello"

# b=Hello("gjgj")
# print(b)



# class Vehical:
#     def __init__(self,name,color,speed):
#         self.name = name
#         self.color = color
#         self.speed = speed
    

#     def SpeedControl(self):
#         if self.speed>100:
#             print("Speed control")
#         else:
#             print("Continue")




class Person:
    def __init__(self, name, sex, profession):
        # data members (instance variables)
        self.name = name
        self.sex = sex
        self.profession = profession

    # Behavior (instance methods)
    def show(self):
        print('Name:', self.name, 'Sex:', self.sex, 'Profession:', self.profession)

    # Behavior (instance methods)
    def work(self):
        print(self.name, 'working as a', self.profession)

# create object of a class
jessa = Person('Jessa', 'Female', 'Software Engineer')

# call methods
jessa.show()
jessa.work()
