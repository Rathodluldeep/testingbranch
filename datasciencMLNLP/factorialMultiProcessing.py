'''
Real-world example:- Multiprocessing CPU-Bound task
Scenario:- Factorial Calculation
Factorial clculations,especially for large numbers,
involve significant computational work.Multiprocessing can be
used to distribute the workload across multiple CPU cores,improving performance
'''

import multiprocessing
import math,sys,time

##Increasing maximum number of digits for integer coversion

sys.set_int_max_str_digits(100000)

##function of compute factorial of a given number

def computer_factorial(number):
    print(f"Computing factorial of {number}")
    result=math.factorial(number)
    print(f"Factorial of {number} is {result}")
    return result

if __name__=="__main__":
    numbers=[5000,6000,700,8000]
    
    start_time=time.time()
    
    ##create a pool of worker process
    with multiprocessing.Pool() as pool:
        results=pool.map(computer_factorial,numbers)
    
    
    end_time=time.time()-start_time
