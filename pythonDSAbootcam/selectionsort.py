li=[15,25,20,50,45,10,40]
n=len(li)
for i in range(n):
    smallIndex=i
    for j in range(i,n):
        if(li[j] < li[smallIndex]):
            smallIndex = j
    print("putting " +str(li[smallIndex]) + " at index "+str(i))
    li[i],li[smallIndex] =li[smallIndex], li[i]
print("after sorting" + str(li))    