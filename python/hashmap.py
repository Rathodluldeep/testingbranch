n=[5,3,2,2,1,5,5,7,5,10]
m=[10,111,1,9,5,67,2]

# constain
# 1) 1<=n[i]<=10 (len !> 11)
# 2) n can have 10^8 element
# 2) m can have 10^8 element

hash_list=[0]*11
for num in n:
    hash_list[num] +=1

result={}
for num in m:
    if 1<=num <= 10:
        result[num]=hash_list[num]
    else:
        result[num]=0
print(result)
