a=[1,2,3,4,5,6]
i=0
j=len(a)-1
while i<j:
    a[i],a[j]=a[i],a[j]
    i=i+1
    j=j-1
print(a)
