def count_word_frequency(sentence):
    '''Words Frequency in a Sentence
Count Word Frequency

Design a Python function named count_word_frequency to count the frequency of words in a sentence and store the counts in a dictionary.'''
    # Your code goes here
    word_count={}
    words=sentence.split()
    for word in words:
        if word in word_count:
            word_count[word] +=1
        else:
            word_count[word] =1
    return word_count
