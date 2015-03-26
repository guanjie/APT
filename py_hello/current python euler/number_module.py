
'''
author: Eric He
date: Feb, 6th, 2013
description: The file contains the module for some number manipulation methods
'''
#==============================main functions==========================
# Get factorial
def factorial(n):
    if n < 1:
        return -1
    if n == 1:
        return 1
    return n*factorial(n-1)

# Get combination
def combination(m, n):
    return factorial(m)/factorial(n)/factorial(m-n)

# Check even
def is_even(num):
    if num <= 0:
        return False
    if num % 2 == 0:
        return True
    else:
        return False

# Check Palindromic
def is_palindromic(string):
    num_str = str(string)
    for i in range(len(num_str)/2):
        if num_str[i] != num_str[len(num_str)-1-i]:
            return False
    return True

# Check primes
def is_prime(num):
    if num <= 1:
        return False
    if num == 2 or num == 3:
        return True
    else:
        for i in range(2,(num+1)/2+1):
            if num % i == 0:
                return False
        return True

def is_truncatable_prime(num):
    num_pivot_to_right = num
    num_pivot_to_left = num
    while(num_pivot_to_right > 0):
        if not is_prime(num_pivot_to_right):
            return False
        num_pivot_to_right /= 10
    while(num_pivot_to_left > 0):
        if not is_prime(num_pivot_to_left):
            return False
        if str(num_pivot_to_left)[1:] != '':
            num_pivot_to_left = int(str(num_pivot_to_left)[1:])
        else:
            num_pivot_to_left = 0
    return True

# Pandigital
def is_pandigital(num):
    # Pandigital number maximum num_length is 9
    num_length = len(str(num))
    if num_length > 9:
        return False
    # Checking element num from set
    num_char_set = set(str(num))
    if num_length != len(num_char_set):
        return False
    # Checking each element in the set
    for i in range(1, num_length+1):
        if not str(i) in num_char_set:
            return False
    return True

# Permutation digits
def permutation_number_list(num):
    num_str = str(num)
    import itertools
    num_list = list(itertools.permutations(num_str))
    result_list = []
    for i_str in num_list:
        result_list.append(int(''.join(i_str)))
    return result_list

# Right Angle Triangles
def is_right_angle_triangle(side_1, side_2, long_side):
    return side_1**2 + side_2**2 == long_side**2

# Has same digits
def have_same_digits(num_1, num_2):
    set_1 = set(str(num_1))
    set_2 = set(str(num_2))
    return set_1 == set_2

#======================Prime Number related===============================
# TODO prime number related





#======================problem specific===================================
# Problem 42: Triangle_words
from math import sqrt
def is_triangle_word(word):
    if not isinstance(word,str):
        print 'the instance is not a string. Wrong input'
        return False
    word_count = get_word_count(word)
    if is_triangle_number(word_count):
        return True
    return False
def get_word_count(word):
    return_count = 0
    for i in word:
        return_count += char_mapping[i]
    return return_count
def is_triangle_number(num):
    num = 2*num
    for i in range(int(sqrt(num))+1):
        if num == i*(i+1):
            return True
    return False

# Problem 44: Pentagon Numbers
first_100000_pentagon_numbers = [i*(3*i-1)/2 for i in range(1, 100000) if i%1 == 0]

# Problem 45: Triangular, Pentagonal, Hexagonal numbers
first_100000_triangular_numbers = [i*(i+1)/2 for i in range(1, 100000)]
first_100000_hexagonal_numbers = [i*(2*i-1) for i in range(1, 100000)]

#======================static methods, values=============================
chars = 'abcdefghijklmnopqrstuvwxyz'
char_mapping = {}
for i in range(len(chars)):
    char_mapping[chars[i]] = i+1

#===========================Less useful===================================
def sum_of_squares(num):
    summation = 0
    for i in range(1, num + 1):
        summation += i**2
    return summation

def square_of_sum(num):
    summation = 0
    for i in range(1, num + 1):
        summation += i
    # got the sum first, next step power 2
    return summation**2


