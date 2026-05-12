n=int(input("etner your number"))
num=n
total=0
nod=len(str(n))
while num>0:
    id=num%10
    total=total + (id**nod)
    num=num//10
print(total==n)  