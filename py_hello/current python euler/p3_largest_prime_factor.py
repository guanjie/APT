'''
name: 3 Largest prime factor
author: Eric He
date: Jan, 24th, 2013
result: 6857

weird: line-33, thought should return first number power 2 out of bound, but it's the last one could be divided. 
'''

def is_prime(num):
    if num <= 1:
        return False
    if num == 2 or num == 3:
        return True
    else:
        for i in range(2,num):
            if num % i == 0:
                return False
        return True

# Idea: check and divide from SMALLEST number
def main():
    pivot = 2
    total_number = 600851475143
    while pivot**2 < total_number:
        if total_number%pivot == 0:
            total_number = total_number/pivot
            print 'divided by this number: {}'.format(pivot)
            # if divisible, continue and try the pivit again
            continue
        # not dividible, pivot ++
        pivot += 1

    print 'the pivot square just over total_number is: {}'.format(pivot)
    print 'the final value can not be divided is: {}'.format(total_number) 

if __name__ == '__main__':
    main()

