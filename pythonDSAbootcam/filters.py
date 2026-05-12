def even(num):
    if num%2==0:
        return True
    

lst=[1,2,3,4,5,6,7,8,9,10,11,12,13]
print(list(filter(even,lst)))

numbers=[1,2,3,4,5,6,7,8,9]
greater_than_five=list(filter(lambda x:x>5,numbers))
print(greater_than_five)

even_and_greater_than_five=list(filter(lambda x:x>5 and x%2==0,numbers))
print(even_and_greater_than_five)


people=[
    {'name':'Krish','age':32},
    {'name':'Jack','age':33}
]

def age_greater_35(person):
    return person['age']>25

print(list(filter(age_greater_35,people)))

