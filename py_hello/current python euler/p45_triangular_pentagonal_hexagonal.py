'''
name: p45_triangular_pentagonal_hexagonal
author: Eric He
date: Feb, 26th, 2013
result: 1533776805

hint: Get the numbers in the list in the first place. Then match the number up by the one line matcher. Could use hash-map to cache the data and then do num_map.has_key() to get log(n) result.
'''
from number_module import first_100000_pentagon_numbers, first_100000_triangular_numbers, first_100000_hexagonal_numbers

def main():
    a = first_100000_triangular_numbers
    b = first_100000_pentagon_numbers
    c = first_100000_hexagonal_numbers

    a_b_overlap = [var for var in a if var in b]
    a_b_c_overlap = [var for var in a_b_overlap if var in c]

    print a_b_c_overlap

#    overlapping = [var for var in c if var in [var for var in a if a in b]]
#    print overlapping

if __name__ == '__main__':
    main()

