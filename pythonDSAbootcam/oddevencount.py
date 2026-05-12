def count_even_odd(lst):
    '''Count Number of Odd and Even Elements in a List
Count Even and Odd Numbers in a List

You are given a list of integers. Write a Python program that counts and returns the number of even and odd numbers in the list.

Parameters:

lst (List of integers): The list of integers where you will count the even and odd numbers.

Returns:

A tuple (even_count, odd_count) where even_count is the number of even numbers and odd_count is the number of odd numbers.'''
    even_count=0
    odd_count=0
    # Your code goes here
    for num in lst:
        if num%2==0:
            even_count += 1
        else:
            odd_count += 1
    
    return even_count, odd_count
