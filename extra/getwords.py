#!/usr/bin/python3.5

file = open('words_alpha.txt','r')
output = open('longwords_alpha.txt','w')
for word in file:
    if len(word)>=5:
        output.write(word)
file.close()
output.close()