'''
name: p39_integer_right_triangles
author: Eric He
date: Feb, 22nd, 2013
result: 840

hint: Separate to smaller methods, always. ATTN: forgot to return the result_count, it nothing returned to the method
'''
from number_module import is_right_angle_triangle

def count_right_angle_triangle(num):
    result_count = 0
    if num < 4:
        return 0
    for side_1 in range(1,num/3):
        for side_2 in range(side_1, (num-side_1)/2):
            long_side = num-side_1-side_2
            if long_side > 0 and is_right_angle_triangle(side_1, side_2, long_side):
                result_count += 1
    return result_count

def main():
    i_pivot = 0
    count_pivot = 0
    for i in range(4, 1000):
        if count_right_angle_triangle(i) > count_pivot:
            i_pivot = i
            count_pivot = count_right_angle_triangle(i)
    print 'number has most ways is: {}, the ways is {}'.format(i_pivot,count_pivot) 

if __name__ == '__main__':
    main()

