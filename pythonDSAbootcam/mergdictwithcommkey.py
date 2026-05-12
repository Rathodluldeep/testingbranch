def merge_dicts_with_overlapping_keys(dicts):
    '''Merge Dictionaries with Common Keys
Problem Description

Merge Dictionaries with Overlapping Keys

Design a Python function named merge_dicts_with_overlapping_keys that merges multiple dictionaries into a single dictionary. If a key appears in more than one dictionary, sum up their values.'''
    # Your code goes here
    result = {}
    for d in dicts:
        for key,value in d.items():
            if key in result:
                result[key] +=value
            else:
                result[key] = value
    return result
