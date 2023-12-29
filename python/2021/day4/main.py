#!/usr/bin/env python3

def flatten_list(_2d_list):
    flat_list = []
    # Iterate through the outer list
    for element in _2d_list:
        if type(element) is list:
            # If the element is of type list, iterate through the sublist
            for item in element:
                flat_list.append(item)
        else:
            flat_list.append(element)
    return flat_list

def check_win(b):
  
  # Checking rows.
  win = True
  for x in range(0,5):    # First row.
    if b[x] != 'X':
      win = False

  if win == True:
    return b

  for x in range(6,10):   # Second row.
    if b[x] != 'X':
      win = False

  if win == True:
    return b

  for x in range(11,15):  # Third row.
    if b[x] != 'X':
      win = False

  if win == True:
    return b

  for x in range(16,20):  # Fourth row.
    if b[x] != 'X':
      win = False

  if win == True:
    return b

  for x in range(21,25):  # Fifth row.
    if b[x] != 'X':
      win = False

  if win == True:
    return b

  # Checking columns.
  win = True
  for x in range(0,25,5): # First column.
    if b[x] != 'X':
      win = False

  if win == True:
    return b

  for x in range(1,25,5): # Second column.
    if b[x] != 'X':
      win =  False

  if win == True:
    return b

  for x in range(2,25,5): # Third column.
    if b[x] != 'X':
      win = False

  if win == True:
    return b

  for x in range(3,25,5): # Fourth column.
    if b[x] != 'X':
      win = False

  if win == True:
    return b

  for x in range(4,25,5): # Fifth column.
    if b[x] != 'X':
      win = False

  if win == True:
    return b

  return win

def sum_board(b, n):
  sum = 0

  for num in b:
    if num == 'X':
      continue
    
    sum += int(num)

  return sum * int(n)

def convert(string):
  li = list(string.split(','))
  return li

# number = '7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1'
number = '4,75,74,31,76,79,27,19,69,46,98,59,83,23,90,52,87,6,11,92,80,51,43,5,94,17,15,67,25,30,48,47,62,71,85,58,60,1,72,99,3,35,42,10,96,49,37,36,8,44,70,40,45,39,0,63,2,78,68,53,50,77,20,55,38,86,54,93,26,88,12,91,95,34,9,14,33,66,41,13,28,57,29,73,56,22,89,21,64,61,32,65,97,84,18,82,81,7,16,24'
number = convert(number)
input = open('input.txt', 'r')


# Importing boards from input.txt.
i = 0
board = []
boards = []
for line in input:
  if line.strip() == '':
    if len(board) > 0: # Don't bring in empty boards
      boards.append(flatten_list(board)) # flatten board into 1d then append.
    board = [] # empty board before next iteration.
    continue
  
  row = [r.replace('\n','') for r in line.split(' ')] # remove newline characters and return line as list.
  row = [r for r in row if r] # remove empty elements.
  board.append(row) # append to board.


# Marks numbers on a board.
i = 0
for n in number:
  for b in boards:
    while i < len(b):
      if b[i] == n:
        print
        b[i] = 'X'
      i += 1

    i = 0 # Reset counter when we exit while to ensure loop repeats from zero on the next board).

    # Check if any boards are winners.
    winner = check_win(b)

    if winner != False:
      ans = sum_board(winner, n)
      print(f'{boards.index(b)+1}: {winner}, {n}, {ans}')
  
  if winner != False:
    break