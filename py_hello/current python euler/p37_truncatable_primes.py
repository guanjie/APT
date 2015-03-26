'''
name: p37_truncatable_primes
author: Eric He
date: Feb, 21st, 2013
result: 748317

hint: This is a bit slow...runs for a day and got the result, the is_prime method need refinary. 
'''

from number_module import is_truncatable_prime
from number_module import is_prime

def main():
    result_sum = 0
    for i in xrange(10,1000000):
        if is_prime(i):
            if is_truncatable_prime(i):
                print '{} is one of them.'.format(i)
                result_sum += i
    print 'result_sum is {}'.format(result_sum)

if __name__ == '__main__':
    main()

