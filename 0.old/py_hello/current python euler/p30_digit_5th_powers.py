'''
name: p30_digit_fifth_powers 
author: Eric He
date: Feb, 13th, 2013
result: 443839

hint: ATTN: when you modify a number from input, make a duplication 
'''

def is_5th_power_sum(num):
    temp = num
    result_sum = 0
    while(temp != 0):
        i = temp%10
        result_sum += i**5
        temp = temp/10
    return num == result_sum

def main():
    result_sum = 0
    for i in range(10, 2000000):
        if is_5th_power_sum(i):
            print '{} is such a number'.format(i)
            result_sum += i
    print 'the result sum of the 5th_power sum is: {}'.format(result_sum)

if __name__ == '__main__':
    main()

