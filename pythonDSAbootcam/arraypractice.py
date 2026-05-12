from array import *

#create array
my_array=array('i',[1,2,3,4,5])
for i in my_array:
    print(i)
    
#acces individual element through indices
for i in range(len(my_array)):
    print(my_array[i])
    print(my_array[3])

#append any value to array using append()
my_array.append(6)
print(my_array)

#insert value in an array using insert() method
my_array.insert(2,8)
my_array.insert(5,9)
my_array.insert(9,10)
print(my_array)

# extend python array using extend() method
my_array2=array('i',[11,12,11,12,12,12,11,13,14,15])
my_array.extend(my_array2)
print(my_array)

# add item from list into array using fromlist() method
templist=[20,21,22]
my_array.fromlist(templist)
print(my_array)

# remove any array element using remove() method
my_array.remove(10)
print(my_array)

# remove last array element using pop() method 
my_array.pop()
my_array.pop()
my_array.pop()
my_array.pop()
print(my_array)

# fetch any element through its index using index() method
print(my_array.index(11))

#reverse a python array by using reverse() method
my_array.reverse()
print(my_array)

#get array buffer inforamtion through buffer_info() method

print(my_array.buffer_info())

#check for number ocurance of element using count() method
print(my_array2.count(12))

#convert array to string using tostring() method
tempstr=my_array2.tobytes()
print(tempstr)

#convert array to a python list with same element using tolist() method
print(my_array.tolist())

#slice element from an arrya
print(my_array2[1:4]) 