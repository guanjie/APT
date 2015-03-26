'''
name: p40_pandigital_prime
author: Eric He
date: Feb, 22nd, 2013
result: 7652413

hint: xrange from now on. Takes forever to run the script. is_prime is still the problem. Fix and upgrade it when it's really unbearable
'''
from number_module import is_pandigital
from number_module import is_prime

def main():
    for i in xrange(1, 1000000000):
        if is_pandigital(i):
            print 'now its here:{}'.format(i)
            if is_prime(i):
                print '{} is one'.format(i)

if __name__ == '__main__':
    main()

