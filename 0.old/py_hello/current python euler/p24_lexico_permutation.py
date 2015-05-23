'''
name: p24_lexicographic_permutation 
author: Eric He
date: Feb, 6st, 2013
result: 2783915460

hint: if use itertools.permutations. 
'''



def main():
    import itertools
    lists = list(itertools.permutations('0123456789'))
    print ''.join(lists[1000000-1])

if __name__ == '__main__':
    main()

