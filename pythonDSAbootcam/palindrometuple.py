def is_palindromic_tuple(tup):
    '''Palindromic Tuple
Check if Tuple is Palindromic

Design a Python function named is_palindromic_tuple to check if a tuple is palindromic, meaning it reads the same forwards and backwards.'''
    # Your code goes here
    start =0
    end=len(tup)-1
    while start <= end:
        if tup[start] != tup[end]:
            return False
        start +=1
        end -= 1
    return True