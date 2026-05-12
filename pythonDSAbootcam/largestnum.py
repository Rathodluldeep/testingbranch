def find_largest(numbers):
    '''Largest Element in a List
Find the Largest Element in a List

Write a Python function that finds and returns the largest element in a given list of integers.

Parameters:

numbers (List of integers): The input list containing integers.

Returns:

An integer representing the largest element in the input list.

Example:

Input: numbers = [3, 8, 2, 10, 5]
Output: 10
    '''
    large=numbers[0]
    # Your code goes here
    for items in numbers[1:]:
        if items>large:
            large=items
            
    return large