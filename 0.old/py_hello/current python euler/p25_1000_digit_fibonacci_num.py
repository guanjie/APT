'''
name: p25_1000_digit_fibonacci_num 
author: Eric He
date: Feb, 8th, 2013
result: 4782th

hint: By using list and append the last 2 digits all the time it's really beautiful and straighforward 
'''

def main():
    fib_list = [1,1]
    while(True):
        fib_list.append(fib_list[-1] + fib_list[-2])
        if len(str(fib_list[-1])) >= 1000:
            print 'the number is: {}'.format(fib_list[-1])
            print 'its the {}th number'.format(len(fib_list))
            break

if __name__ == '__main__':
    main()

