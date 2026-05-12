a=[5,7,9,13,20,25,32]
b=[8,11,11,15]

n=len(a)
m=len(b)
c=[0]*(n+m)
i = 0
j = 0
k = 0

while(i<n or j < m):
    if(i<n and j < m):
        if (a[i]<= b[j]):
            c[k]=a[i]
            i+=1
        else:
            c[k]=b[j]
            j+= 1
        k += 1
    elif i < n:
        c[k] = a[i]
        i+=1
        k+=1
    else:
        c[k] = b[j]
        j+=1
        k+=1
print(c)