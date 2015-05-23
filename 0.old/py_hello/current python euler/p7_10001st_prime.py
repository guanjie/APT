'''
name: 7_10001st_prime 
author: Eric He
date: Jan, 26th, 2013
result: 104743 

hint: using module is_prime 
'''
from p3_largest_prime_factor import is_prime 

def main():
    count = 0
    pivot = 0
    while(True):
        # pivot adding 1 each loop, count adding 1 when pivot is a prime number
        if is_prime(pivot):
            count += 1
            if count == 10001:
                print 'the 10001st prime number is: {}'.format(pivot)
                break
        pivot  += 1

if __name__ == '__main__':
    main()

