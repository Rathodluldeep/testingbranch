def remove_duplicates(lst):
    '''Remove Duplicate in a List
Remove Duplicates from a List

You are given a list of integers. Write a Python program that removes any duplicate elements from the list and returns a new list with only unique elements. The order of elements in the list should be maintained.

Parameters:

lst (List of integers): The list of integers from which duplicates should be removed.

Returns:

A list of integers where all duplicates have been removed, preserving the original order.'''
    lst1=[]
    # Your code goes here
    for num in lst:
        if num not in lst1:
            lst1.append(num)
            
    return lst1
        