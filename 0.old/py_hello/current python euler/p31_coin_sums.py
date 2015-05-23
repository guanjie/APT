'''
name: p31_coin_sums 
author: Eric He
date: Feb, 13th, 2013
result: 73682

hint: a nested for loop with "range" method would do the work
do it again when have time
'''

def main():
    result_sum = 0
    money = 200
    # reducing money by different graduate amount
    for a in range(money, -1, -200):
        for b in range(a, -1, -100):
            for c in range(b, -1, -50):
                for d in range(c, -1, -20):
                    for e in range(d, -1, -10):
                        for f in range(e, -1, -5):
                            for g in range(f, -1, -2):
                                for h in range(g,-1,-1):
                                    if h == 0:
                                        result_sum += 1
    print 'there are in tatal {} number combinations'.format(result_sum)


if __name__ == '__main__':
    main()

