'''
name: p18_maximum_path_sum_I
author: Eric He
date: Feb, 1st, 2013
result: 

hint: 
'''







def main():
    num = 2**1000
    num_string = str(num)
    sum = 0
    for i in num_string:
        sum += int(i)

    print 'the total digit sum is: {}'.format(sum)




if __name__ == '__main__':
    main()

