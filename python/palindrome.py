n=int(input("Enter your number"))
num=n
result=0
while num >0:
    lastdigit=num%10
    result=(result*10)+lastdigit
    num=num//10
print(n==result)
print(result)
