'''
name: 6_sum_square_difference 
author: Eric He
date: Jan, 26th, 2013
result: 25164150 

hint: straight-forward 
'''

def sum_of_squares(num):
    sum = 0
    for i in range(1, num + 1):
        sum += i**2
    return sum

def square_of_sum(num):
    sum = 0
    for i in range(1, num + 1):
        sum += i
    # got the sum first, next step power 2
    return sum**2

def main():
    num = 100
    print 'diff is: {}'.format(abs(sum_of_squares(num)-square_of_sum(num)))
    pass

if __name__ == '__main__':
    main()

