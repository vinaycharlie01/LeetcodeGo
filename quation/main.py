# def greet(a):
#     return "Hello World!"

# greeting = greet

# def greet():
#     return "Goodbye World!"

# print(greeting(10))


# a=["Mike", "Pinky", "Brain", "Dot"]
# dis={name.upper:age for name,age in zip(a)}
# print(dis)

# a = ["Mike", "Pinky", "Brain", "Dot"]
# a = ["Mike", "Pinky", "Brain", "Dot"]
# a = ["Mike", "Pinky", "Brain", "Dot"]
# dis = {age.upper(): (len(age) if len(age) > 3 else "length is 3" if len(age) == 3 
#                     else "short") for age in a}

# a=[a[i] if i%2==0 else "wrong" if i%2==1 else "write" for i in range(len(a))]
# print(a)


# d=["Hi","Hello","world"]
dis={"vinay":30,"karthi":40,"hello":40}
max_key ={k:v for k, v in dis.items() if v==max(dis.values())}

# l={val:max(dis.items()) for i,val in enumerate(dis)}
print(max_key)

