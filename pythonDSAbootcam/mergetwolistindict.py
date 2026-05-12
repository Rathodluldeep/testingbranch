def merge_lists_to_dictionary(keys, values):
    '''Merge 2 List into Dictionary
Merge Lists to Dictionary

Design a Python function named merge_lists_to_dictionary to merge two lists into a dictionary where elements from the first list act as keys and elements from the second list act as values.'''
    # Your code goes here
    if(len(keys) != len(values)):
        return False
    result = {}
        
    for i in range(len(keys)):
        result[keys[i]]=values[i]
    return result