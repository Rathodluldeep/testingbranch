##Multithreading
##when to use multithreading
##I/O-bound task:- Task that spend more time waiting for I/O operations(eg., file operations,network requests).
## Concurrent execution:-When you want to imporove the throughput of your application by performing multiple operation concurrently.

import threading
import time

def print_numbers():
    for i in range(5):
        time.sleep(1)
        print(f"Number: {i}")

def print_letter():
    for letter in "abcde":
        time.sleep(1)
        print(f"Letter: {letter}")
        
##create 2 threades
t1= threading.Thread(target=print_numbers)
t2=threading.Thread(target=print_letter)

t=time.time()
##start the tread
t1.start()
t2.start()

## wait for thread to complete
t1.join()
t2.join
finished_time=time.time()-t
print(finished_time)