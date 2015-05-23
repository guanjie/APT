'''
name: p14_longest_collatz_sequence 
author: Eric He
date: Feb, 1st, 2013
result: 837799 

hint: Separating the bigger methods into smaller ones.  
'''

def is_even(num):
    if num < 0: 
        print 'number is less than 0'
        return 
    return num%2 == 0

def get_new_number(num):
    if is_even(num):
        return num/2
    else:
        return 3*num + 1

def get_i_terms(n):
    terms = 1
    while(n > 1):
        n = get_new_number(n)
        terms += 1
    return terms

def main():
    total_num = 1000000
    starting_num = 1
    total_terms = 1
    for i in range(1, total_num + 1):
        i_terms = get_i_terms(i)
        if i_terms > total_terms:
            total_terms = i_terms
            starting_num = i
    print 'we have the longest sequence, there are {} terms, starting from {}'.format(total_terms, starting_num)

if __name__ == '__main__':
    main()

