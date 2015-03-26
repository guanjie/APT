'''
name: p22_names_scores 
author: Eric He
date: Feb, 4st, 2013
result: 850081394 a bit different than the website but it feels like mine is more correct

hint: remember to do line.strip() to strip off the line breakers 
'''
# Getting the char_mapping 
chars = 'abcdefghijklmnopqrstuvwxyz'
char_mapping = {}
for i in range(len(chars)):
    char_mapping[chars[i]] = i+1

def get_word_count(word):
    word_count = 0
    for i in word:
        word_count += char_mapping[i]
    return word_count

def main():
    f = open("names.txt",'r')
    total_count = 0
    pivot = 1
    for line in f.readlines():
        line = line.strip()
        total_count += pivot*get_word_count(line)
        pivot += 1
        print 'the current line is: {}'.format(line)
    print 'total_count is: {}'.format(total_count)

if __name__ == '__main__':
    main()

