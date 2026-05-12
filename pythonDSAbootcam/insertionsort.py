li=[15,25,20,50,45,10,40]
n=len(li)
for i in range(1,n):
    j=i-1
    key=li[i]
    while(j >= 0 and li[j] > key):
        li[j+1]=li[j]
        j -= 1
    li[j+1] = key
    print(li[:i+1])

print("after full sorting" + str(li))    