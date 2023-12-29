#!/usr/bin/env python3

from collections import Counter

def match_legend(a, l) -> list:
  f = []
  for i,v in enumerate(a):
    if i < len(a)-1:
      c = a[i] + a[i+1]

      for k,v in l.items():
        if c == k:
          f.append(a[i])
          f.append(v)
  f.append(a[-1])

  return f

with open('input.txt') as file:
  lines = file.readlines()
ans = list(lines[0].rstrip())

pairs = []
for x,y in enumerate(lines):
  if x > 1:
    pairs.append(y.rstrip())

legend = {}
for pair in pairs:
  key, value = pair.split(' -> ')
  legend[key] = value

for n in range(10):
  ans = match_legend(ans, legend)

big = max(Counter(ans).values())
small = min(Counter(ans).values())

print(big - small)