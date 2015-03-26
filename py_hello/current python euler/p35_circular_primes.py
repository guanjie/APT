'''
name: p35_circular_primes.py 
author: Eric He
date: Feb, 19th, 2013
result: 

hint: 
'''

from number_module import is_prime
from number_module import permutation_number_list

def is_circular_prime(num):
    num_list = permutation_number_list(num) 
    for i in num_list:
        if not is_prime(i):
            return False
    return True

def main():
    num = 1000000
    result_count = 0
    for i in range(2, num):
        if is_prime(i):
            if is_circular_prime(i):
                result_count += 1
                print "{} is a circular_prime.".format(i)
    print "There are total {} circular primes".format(result_count)

if __name__ == '__main__':
    main()

