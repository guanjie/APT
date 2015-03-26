'''
name: p34_digit_factorials.py 
author: Eric He
date: Feb, 19th, 2013
result: My number is different: 145, and I don't think I am wrong in any way

hint: Don't know where I made the mistake
'''
from number_module import factorial

def is_curious(num):
    digit_sum = 0
    num_str = str(num)
    for i_str in num_str:
        i = int(i_str)
        digit_sum += factorial(i)
    return num == digit_sum

def main():
    result_sum = 0
    for i in range(10,100000):
        if is_curious(i):
            print '{} is a curious number'.format(i)
            result_sum += i
    print "{} is the total sum".format(result_sum)

if __name__ == '__main__':
    main()

