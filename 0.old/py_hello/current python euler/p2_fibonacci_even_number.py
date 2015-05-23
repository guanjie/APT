'''
name: 2 Even Fibonacci Numbers
author: Eric He
date: Jan, 24th, 2013
result: 4613732
'''
bound = 4000000

def is_even(num):
    if num <= 0:
        return False
    if num % 2 == 0:
        return True
    else:
        return False

def get_fibonacci_list():
    fib_list = []
    fib_list.append(1)
    fib_list.append(2)
    last_2_sum = fib_list[-1] + fib_list[-2]
    while(last_2_sum < bound):
        fib_list.append(last_2_sum)
        last_2_sum = fib_list[-1] + fib_list[-2]
    return fib_list

def main():
    sum = 0
    fib_list = get_fibonacci_list()
    for i in fib_list:
        if(is_even(i)):
            sum += i
    print 'the summation of all fib numbers less than 4000000 is: {}'.format(sum)

if __name__ == '__main__':
    main()

