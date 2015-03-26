'''
name: p29_distinct_powers 
author: Eric He
date: Feb, 13th, 2013
result: 9183

hint: use set() 
'''

def main():
    result_set = set()
    for i in range(2, 101):
        for j in range(2, 101):
            result_set.add(i**j)
    print 'there are {} many distinct terms in the sequence'.format(len(result_set))

if __name__ == '__main__':
    main()

