'''
name: 5_smallest_multiple 
author: Eric He
date: Jan, 26th, 2013
result:  232792560

hint: STAGE 1-brute force 
'''
divisible_under = 20

def divisible_by_under(total):
    for i in range (divisible_under, 0, -1):
        if total%i != 0:
            return False
    return True

def main():
    total = 1
    while(True):
        if divisible_by_under(total):
            print 'the number is : {}'.format(total)
            break
            return
        total += 1
    print 'smallest multiple is: {}'.format(total)

if __name__ == '__main__':
    main()

