'''
name: p48_self_powers
author: Eric He
date: mar, 13th, 2013
result: 9110846700

hint: when do sum, put numbers in the list and use default tools sum()
'''

def main():
    num_list = [i**i for i in xrange(1,1001)]
    result_sum = sum(num_list)
    # get the last 10 digits of the result_sum
    print result_sum%(10**10) 

if __name__ == '__main__':
    main()

