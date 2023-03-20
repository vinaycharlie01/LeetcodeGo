def traffic_signal_message(f):
    def wrapper(color):
        print(f"{color.upper()} : ",end="")
        f(color)
    return wrapper

@traffic_signal_message
def stop_message(color):
    print("STOP")

@traffic_signal_message
def slow_down_message(color):
    print("SLOW DOWN")

@traffic_signal_message
def go_message(color):
    print("GO")

stop_message("red")
slow_down_message("yellow")
go_message("green")
