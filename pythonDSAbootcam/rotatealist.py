def rotate_list(lst, k):
    '''Rotate a List
Rotate a List (Without Slicing)

You are given a list of integers and an integer k. Write a Python function to rotate the list to the right by k positions without using slicing. A rotation shifts elements from the end of the list to the front.'''
    # Your code goes here
    if not lst:
        return []
        
    k=k%len(lst)
    for _ in range(k):
        last_element=lst.pop()
        lst.insert(0,last_element)
    return lst
