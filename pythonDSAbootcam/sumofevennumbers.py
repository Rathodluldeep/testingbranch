def sum_of_even_numbers(n):
    """
    Function to return the sum of the first n even natural numbers.
    
    Parameters:
    n (int): The number of even numbers to sum.
    
    Returns:
    int: The sum of the first n even natural numbers.
    """
    # Your code here
    total_sum =0
    current_even_number= 2
    for i in range(n):
        total_sum += current_even_number
        current_even_number += 2
    return total_sum