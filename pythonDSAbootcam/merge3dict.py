def merge_three_dictionaries(dict1, dict2, dict3):
    '''Merge Multiple Dictionaries
Merge Three Dictionaries

Design a Python function named merge_three_dictionaries to merge exactly three dictionaries into one.'''
    # Your code goes here
    merged_dict = dict1.copy()
    
    for key,value in dict2.items():
        merged_dict[key] = value
    
    for key,value in dict3.items():
        merged_dict[key]=value
    return merged_dict
