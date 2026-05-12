def reverse_list(lst):
    '''Program to Reverse a List
Reverse a List (Non-Slicing Approach)

You are given a list of integers. Write a Python program that reverses the list without using slicing (lst[::-1]). The program should return the reversed list.

Parameters:

lst (List of integers): The list of integers to be reversed.

Returns:

A list of integers where the order of elements is reversed from the input list.'''
    start=0
    end=len(lst)-1
    # Your code goes here
    while start<end:
        lst[start],lst[end]=lst[end],lst[start]
        start+=1
        end-=1
    return lst
    
