


def third_sort(arr):
    if len(arr)<3:
        arr.sort()
        return max(arr)
    else:
        ar=list(set(arr))
        return ar[-3]
