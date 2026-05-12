def generate_hollow_square(n):
    """
    Function to return a hollow square pattern of '*' of side n as a list of strings.
    
    Parameters:
    n (int): The size of the square.
    
    Returns:
    list: A list of strings where each string represents a row of the hollow square.
    """
    # Your code here
    square=[]
    if n ==1:
        return ['*']
        
        square.append('*' * n)
    for i in range(n-2):
        square.append('*' + ' ' *(n-2) + '*')
        square.append('*' * n)
        
    return square
print(generate_hollow_square(5))



def generate_square(n):
    """
    Function to return a square pattern of '*' of side n as a list of strings.
    
    Parameters:
    n (int): The size of the square.
    
    Returns:
    list: A list of strings where each string represents a row of the square.
    """
    square=[]
    for i in range(n):
        square.append('*' * n)
        
    return square

print(generate_square(7))



def generate_rectangle(n, m):
    """
    Function to return a rectangle pattern of '*' with length n and breadth m as a list of strings.
    
    Parameters:
    n (int): The number of rows in the rectangle.
    m (int): The number of columns in the rectangle.
    
    Returns:
    list: A list of strings where each string represents a row of the rectangle pattern.
    """
    # Your code here
    rectangle=[]
    for _ in range(n):
        row='*'*m
        rectangle.append(row)
    return rectangle
    
print(generate_rectangle(3,5))




def generate_triangle(n):
    """
    Function to return a right-angled triangle of '*' of side n as a list of strings.
    
    Parameters:
    n (int): The height and base of the triangle.
    
    Returns:
    list: A list of strings where each string represents a row of the triangle.
    """
    # Your code here
    triangle=[]
    for i in range(1, n+1):
        triangle.append('*' * i)
    return triangle
    
print(generate_triangle(5))


def generate_inverted_triangle(n):
    """
    Function to return an inverted right-angled triangle of '*' of side n as a list of strings.
    
    Parameters:
    n (int): The height and base of the triangle.
    
    Returns:
    list: A list of strings where each string represents a row of the triangle.
    """
    # Your code here
    triangle=[]
    for i in range(n,0,-1):
        triangle.append('*'*i)
        
    return triangle
print(generate_inverted_triangle(6))


def generate_pyramid(n):
    """
    Function to return a pyramid pattern of '*' of side n as a list of strings.
    
    Parameters:
    n (int): The number of rows in the pyramid.
    
    Returns:
    list: A list of strings where each string represents a row of the pyramid.
    """
    # Your code here
    pyramid=[]
    for i in range(1,n+1):
        stars='*' * (2*i-1)
        spaces=' '*(n-i)
        
        pyramid.append(spaces + stars + spaces)
        
    return pyramid
print(generate_pyramid(5))



def generate_inverted_pyramid(n):
    """
    Function to return an inverted pyramid pattern of '*' of side n as a list of strings.
    
    Parameters:
    n (int): The number of rows in the inverted pyramid.
    
    Returns:
    list: A list of strings where each string represents a row of the inverted pyramid.
    """
    # Your code here
    pyramid=[]
    for i in range(1,n+1):
        stars='*'*(2*(n-i+1)-1)
        spaces=' '*(i-1)
        pyramid.append(spaces+stars+spaces)
    return pyramid
        
print(generate_inverted_pyramid(5))


def generate_number_triangle(n):
    """
    Function to return a right-angled triangle of repeated numbers of side n as a list of strings.
    
    Parameters:
    n (int): The height of the triangle.
    
    Returns:
    list: A list of strings where each string represents a row of the triangle.
    """
    # Your code here
    triangle=[]
    for i in range(1,n+1):
        row=str(i) * i
        triangle.append(row)
    return triangle
        
print(generate_number_triangle(5))

def generate_floyds_triangle(n):
    """
    Function to return the first n rows of Floyd's Triangle as a list of strings.
    
    Parameters:
    n (int): The number of rows in the triangle.
    
    Returns:
    list: A list of strings where each string represents a row of Floyd's Triangle.
    """
    # Your code here
    Triangle=[]
    current_num=1
    for i in range(1,n+1):
        row=' '.join(str(current_num + j) for j in range(i))
        Triangle.append(row)
        current_num+=i
    return Triangle
print(generate_floyds_triangle(6))

def generate_diamond(n):
    """
    Function to return a diamond pattern of '*' of side n as a list of strings.
    
    Parameters:
    n (int): The number of rows for the upper part of the diamond.
    
    Returns:
    list: A list of strings where each string represents a row of the diamond.
    """
    # Your code here
    diamond=[]
    for i in range(1,n+1):
        stars='*'*(2*i-1)
        spaces=' '*(n-i)
        diamond.append(spaces+stars+spaces)
    
    for i in range(n-1,0,-1):
        stars='*'*(2*i-1)
        spaces=' '*(n-i)
        diamond.append(spaces+stars+spaces)
        
    return diamond
print(generate_diamond(6))

def generate_number_pyramid(n):
    """
    Function to return a pyramid pattern of numbers of height n as a list of strings.
    
    Parameters:
    n (int): The height of the pyramid.
    
    Returns:
    list: A list of strings where each string represents a row of the pyramid pattern.
    """
    # Your code here
    pyramid=[]
    for i in range(1,n+1):
        spaces=' '*(n-i)
        numbers=' '.join(str(x) for x in range(1,i+1))
        pyramid.append(spaces+numbers+spaces)
    return pyramid
print(generate_number_pyramid(6))

def generate_hollow_inverted_right_angled_triangle(n):
    """
    Function to return a hollow inverted right-angled triangle of '*' of side n as a list of strings.
    
    Parameters:
    n (int): The height of the triangle.
    
    Returns:
    list: A list of strings where each string represents a row of the triangle.
    """
    # Your code here
    triangle=[]
    for i in range(n,0,-1):
        if i==n:
            triangle.append('*' * n)
        elif i==1:
            triangle.append('*')
        else:
            spaces=' '*(i-2)
            triangle.append('*' + spaces + '*')
    return triangle
print(generate_hollow_inverted_right_angled_triangle(5))

def generate_sandglass(n):
    """
    Function to return a sandglass pattern of '*' of side n as a list of strings.
    
    Parameters:
    n (int): The height of the sandglass.
    
    Returns:
    list: A list of strings where each string represents a row of the sandglass pattern.
    """
    # Your code here
    sandglass=[]
    for i in range(n):
        stars='*' *(2*(n-i)-1)
        spaces=' ' *i
        sandglass.append(spaces+stars+spaces)
        
    for i in range(n-1):
        stars='*' * (2*(i+1)+1)
        spaces=' ' * (n-i-2)
        
        sandglass.append(spaces+stars+spaces)
    return sandglass
print(generate_sandglass(7))


def generate_right_angled_triangle(n):
    """
    Function to return a right-angled triangle of '*' of side n as a list of strings.
    
    Parameters:
    n (int): The height of the triangle.
    
    Returns:
    list: A list of strings where each string represents a row of the triangle.
    """
    # Your code here
    triangle = []
    
    
    for i in range(1, n + 1):
        spaces = ' ' * (n - i)
        stars = '*' * i
        triangle.append(spaces + stars)
    
    return triangle
print(generate_right_angled_triangle(6))