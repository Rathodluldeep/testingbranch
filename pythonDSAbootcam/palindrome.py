def is_palindrome(s):
    s=s.lower().replace(" ","")
    return s==s[::-1]

print(is_palindrome(" A man a plan a canal Panama "))
print(is_palindrome("Hello"))