'''
name: p36_palindrome_num_sum.py
author: Eric He
date: Feb, 21st, 2013
result: 872187

hint: Build up is_palindromic function first
'''

from number_module import is_palindromic

def main():
    maximum = 1000000
    result_sum = 0
    for i in range(1, maximum):
        if is_palindromic(i):
            if(is_palindromic(bin(i)[2:])):
                print '{} is palindromic, and its binary format is palindromic too: {}'.format(i, bin(i)[2:])
                result_sum += i
    print 'finaly result_sum is: {}'.format(result_sum)

if __name__ == '__main__':
    main()

