'''
name: p52_permuted_multiples
author: Eric He
date: mar, 14th, 2013
result: 142857

hint: Beautiful. Put num_char into the list and then get set and join it back up
'''
from number_module import have_same_digits

def main():
    for i in xrange(1,10000000):
        # Get the list
        list_i = [2*i, 3*i, 4*i, 5*i, 6*i]
        # Get the set of each str(num)
        list_set = [set(str(num)) for num in list_i]
        # Join each mem in list_set to the only list_order after putting into set
        list_str = [''.join(mem) for mem in list_set]
        # Put the list into a set: wish to get only 1 result
        set_list_str = set(list_str)
        # Now we have 1 result, it means all 5 vals have the same number
        # proposition
        if len(set_list_str) == 1:
            print i, set_list_str
            for index in range(2,7):
                print i*index
            break
    return

if __name__ == '__main__':
    main()

