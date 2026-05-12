'''You are asked to ensure that the first and last names of people begin with a capital letter in their passports. For example, alison heck should be capitalised correctly as Alison Heck.


Given a full name, your task is to capitalize the name appropriately.

Input Format

A single line of input containing the full name, .

Constraints

The string consists of alphanumeric characters and spaces.
Note: in a word only the first character is capitalized. Example 12abc when capitalized remains 12abc.'''
def capitalize_name(name):
    words = name.split(" ")
    capitalized_words = []
    for word in words:
        if word:
            capitalized_word = word[0].upper() + word[1:] if word[0].isalpha() else word
            capitalized_words.append(capitalized_word)
        else:
            capitalized_words.append("")  # For multiple spaces
    return " ".join(capitalized_words)

# Example usage
full_name = input()
print(capitalize_name(full_name))
