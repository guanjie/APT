'''
name: 10_summation_primes
author: Eric He
date: Jan, 26th, 2013
result: 

hint: 
'''
from p3_largest_prime_factor import is_prime

total = 2000000

def main():
    # question asks to check BELOW 2 million
    sum = 0
    for i in range(1, total):
        if is_prime(i):
            sum += i
    print 'final sum of the primes is: {}'.format(sum)

if __name__ == '__main__':
    main()

