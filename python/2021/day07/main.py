#!/usr/bin/env python3

input = open('input.txt', 'r')

for line in input:
  crabs = [int(crab) for crab in line.split(',')]
max_crab = max(crabs)

all_moves = []
moves = 0

for pos in range(max_crab): # iterate every possible number of moves up to highest number in crabs list.
  for x in range(len(crabs)): # iterate all crabs in crabs list.
    if crabs[x] > pos:
      moves += crabs[x] - pos

    if crabs[x] < pos:
      moves += pos - crabs[x]

  all_moves.append(moves)
  moves = 0

print(f'Part 1: {min(all_moves)}')

all_moves = []
moves = 0
steps = 0

for pos in range(max_crab):
  for x in range(len(crabs)):
    if crabs[x] > pos:
      moves = crabs[x] - pos
      for i in range(1, moves+1):
        steps += i

    if crabs[x] < pos:
      moves = pos - crabs[x]
      for i in range(1, moves+1):
        steps += i
    
  all_moves.append(steps)
  moves = 0
  steps = 0

print(f'Part 2: {min(all_moves)}')