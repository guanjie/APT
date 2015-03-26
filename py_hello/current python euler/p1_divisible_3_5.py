'''
name: 1 divisible_by_3_5
author: Eric He
date: Jan 24th, 2013
The result is 234168
'''
x = 1000

def divisible_by_3_5(x):
    if x <= 0:
        return False
    if ((x % 3 == 0) or (x % 5 == 0)):
        return True
    else:
        return False

def main():
    sum = 0
    for i in range(1,x+1):
        if divisible_by_3_5(i):
            sum += i

    print 'the summation is {}'.format(sum)

if __name__ == '__main__':
    main()
