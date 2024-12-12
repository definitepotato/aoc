#!/usr/bin/env python3

file = open('input.txt', 'r')
binaries = file.read().split("\n")

def binary_to_decimal(binary):
  return int(binary, 2)

def oxygen_criteria(binaries, x):
  zeros = 0
  ones = 0
  zlist = []
  olist = []

  if len(binaries) == 1:
    return binaries[0]

  for binary in binaries:
    if binary[x] == '0':
      zeros += 1
      zlist.append(binary)
    else:
      ones += 1
      olist.append(binary)

  if zeros > ones:
    x += 1
    return oxygen_criteria(zlist, x)
  elif zeros == ones:
    x += 1
    return oxygen_criteria(olist, x)
  else:
    x += 1
    return oxygen_criteria(olist, x)

def co2_criteria(binaries, x):
  zeros = 0
  ones = 0
  zlist = []
  olist = []

  if len(binaries) == 1:
    return binaries[0]

  for binary in binaries:
    if binary[x] == '0':
      zeros += 1
      zlist.append(binary)
    else:
      ones += 1
      olist.append(binary)

  if zeros > ones:
    x += 1
    return co2_criteria(olist, x)
  elif zeros == ones:
    x += 1
    return co2_criteria(zlist, x)
  else:
    x += 1
    return co2_criteria(zlist, x)

def part1():
  gamma = ''
  epsilon = ''
  width = int(len(binaries[0]))

  for x in range(width):
    zeros = 0
    ones = 0

    for binary in binaries:
      if binary[x] == '0':
        zeros += 1
      else:
        ones += 1

    if zeros > ones:
      gamma += '0'
      epsilon += '1'
    else:
      gamma += '1'
      epsilon += '0'

  gd = binary_to_decimal(gamma)
  ed = binary_to_decimal(epsilon)
  pc = gd * ed

  print(gamma)
  print(epsilon)
  print(pc)

def part2():
  oxygen = 0
  co2 = 0
  
  oxygen = binary_to_decimal(oxygen_criteria(binaries, 0))
  co2 = binary_to_decimal(co2_criteria(binaries, 0))

  print(oxygen * co2)

# part1()
part2()