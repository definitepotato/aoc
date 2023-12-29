#!/usr/bin/env python3

input = open('input.txt', 'r')

# collect all the fishies
for line in input:
  fishies = line.split(',')

# convert all fishies to integers
for i in range(0, len(fishies)):
  fishies[i] = int(fishies[i])

def reduce_school(f):
  for index, fish in enumerate(f):
    if f[index] == 'A':
      continue

    f[index] = int(fish) - 1
  return f

def scan_school(f):
  for index, fish in enumerate(f):
    if fish == 'A':
      f[index] = 6
      f.append(8)

    if fish == 0:
      f[index] = 'A'
  return f

for x in range(80):
  reduce_school(fishies)
  scan_school(fishies)

print(len(fishies))