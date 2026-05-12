a=[7,9,3,4,1,5,10,15,12,9,18,23,14]

def mergeSort(left, right):
    if (left < right):
        mid=(left + right)//2
        mergeSort(left, mid)
        mergeSort(mid + 1, right)
        merge(left, mid, right)
        
def merge(start, mid, end):
    i=start
    j=mid + 1
    k=0
    c=[0]*(end-start +1)
     
    
    while(i<= mid or j <= end):
        if(i<=mid and j<= end):
            if(a[i] <=a[j]):
                c[k]=a[i]
                i+= 1
            else:
                c[k] =a[j]
                j+= 1
        elif(i<= mid):
            c[k] =a[i]
            i+=1
        else:
            c[k]=a[j]
            j +=1
        k+=1
    for i in range(len(c)):
        a[start+i]=c[i]
mergeSort(0, len(a) - 1)

print(a)