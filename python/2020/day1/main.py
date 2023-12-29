#!/usr/bin/env python3

with open('input.txt') as input:
  e = input.readlines()
  e = [int(e.strip()) for e in e]

# part 1
for i in e:
  for j in e:
    if i + j == 2020:
      print(i * j)

# part 2
for i in e:
  for j in e:
    for k in e:
      if i + j + k == 2020:
        print(i * j * k)