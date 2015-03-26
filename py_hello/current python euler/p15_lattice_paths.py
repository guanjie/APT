'''
name: p15_lattice_paths 
author: Eric He
date: Feb, 1st, 2013
result: 137846528820

hint: Doing math the raw way is easier sometime, Running API is sometime more time-consuming 
'''

def factorial(n):
    if n == 1:
        return 1
    return n*factorial(n-1)

def combination(m, n):
    return factorial(m)/factorial(n)/factorial(m-n)

def main():
    square_side = 20
    print 'the total ways of routes is: {}'.format(combination(square_side*2, square_side))

if __name__ == '__main__':
    main()

