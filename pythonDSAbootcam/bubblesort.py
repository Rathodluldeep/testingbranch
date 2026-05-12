li=[10,25,20,50,45,15,40]
n=len(li)

for i in range(n):
    for j in range(n-i-1):
        if li[j] > li[j+1]:
            li[j], li[j+1] = li[j+1], li[j]
        
print("after sorting" + str(li))