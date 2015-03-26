'''
name: p20_factorial_digit_sum 
author: Eric He
date: Feb, 4st, 2013
result: 648

hint: make the long a string and add back up the digits 
'''
def factorial(n):
    if n == 1:
        return 1
    elif n < 1:
        print 'error!'
        return
    else:
        return n*factorial(n-1)

def main():
    num = 100
    fact = factorial(num)
    fact_string = str(fact)
    summation = 0

    for i in fact_string:
        summation += int(i)
    print 'the sigit sum is {}'.format(summation)

if __name__ == '__main__':
    main()

