num=int(input("Enter Your Number"))
counter=0
while num > 0:
    lastdigit= num % 10
    counter += 1
    num = num // 10
print(counter)