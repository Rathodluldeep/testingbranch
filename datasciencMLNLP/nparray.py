import numpy as np

# create array using numpy
arr1=np.array([1,2,3,4,5])
print(arr1)
print(type(arr1))
print(arr1.shape)

arr2=np.array([[1,2,3,4,5],[2,3,4,5,6]])
print(arr2)
print(arr2.shape)

print(np.arange(0,10,2).reshape(5,1))
print(np.ones((3,4)))

#identity matrix
print(np.eye(3))

arr=np.array([[1,2,3],[4,5,6]])
print(arr)
print(arr.ndim)
print(arr.size)
print(arr.dtype)
print(arr.itemsize)

#vectorize operation
arr1=np.array([1,2,3,4,5])
arr2=np.array([10,20,30,40,50])

##element wise operation
print(arr1+arr2)
print(arr1-arr2)
print(arr1*arr2)

# universal functions of array
arr=np.array([2,3,4,5,6])
#sqyare root
print(np.sqrt(arr))
 #exponential
print(np.exp(arr))
#sine
print(np.sin(arr))
#natural log
print(np.log(arr))

## array slicing and indexing
arr=np.array([[1,2,3,4],[5,6,7,8],[9,10,11,12]])
print(arr[0][0])
print(arr[1:,2:])
print(arr[0:2,2:])
print(arr[1:,1:3])

# modify array elements

arr[0,0]=100
print(arr)
arr[1:]=100
print(arr)

# statistical concepts--Normalization
# to have a mean of 0 and standard deviation of 1 
data=np.array([1,2,3,4,5])

# calculate the mean and standard deviation
mean=np.mean(data)
std_dev = np.std(data)

# Normalize the data
normalized_data=(data - mean)/std_dev
print("Normalized data:",normalized_data)

# mean
mena=np.mean(data)
print('Mean:',mean)

# Median
median = np.median(data)
print("Median:", median)

# Standard deviation
std_dev=np.std(data)
print("Standard Deviation:", std_dev)

# Variance
variance = np.var(data)
print("Variance:",variance)

## Logical operations
data=np.array([1,2,3,4,5,6,7,8,9,10])
print(data[(data>5) & (data<8)])
