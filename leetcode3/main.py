def prefix_sum(lst):
    result = [0] * len(lst)
    result[0] = lst[0]
    for i in range(1, len(lst)):
        result[i] = result[i-1] + lst[i]
    return result

numbers = [1, 2, 3, 4, 5]
prefix_sum_numbers = list(map(prefix_sum, numbers))
print(prefix_sum(numbers))