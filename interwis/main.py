# def multiply(x, y):
#     result = 0
#     for i in range(y):
#         result += x
#     return result


import time

interval =2
timer = time.time() + interval

while True:
    if time.time() >= timer:
        print("******")
        print("*****")
        print("****")
        print("***")
        print("**")
        print("*")
        timer = time.time() + interval

