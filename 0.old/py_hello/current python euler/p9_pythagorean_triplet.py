'''
name: 9_pythagorean triplet 
author: Eric He
date: Jan, 26th, 2013
result: 31875000 

hint: Mentally go through the process that a < b < c and c**2 equals a**2 and b**2 
'''
total = 1000
def main():
    for a in range(1,total):
        for b in range(a+1, total):
            c = total - a - b
            if c > b:
                if a**2 + b**2 == c**2:
                    print 'we got the 3 number: {}, {}, {}'.format(a,b,c)
                    print 'then the product of the 3 is: {}'.format(a*b*c)
                    break

if __name__ == '__main__':
    main()

