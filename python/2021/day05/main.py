#!/usr/bin/env python3


input = open('input.txt', 'r')
G = {}

for line in input:
  start, end = line.strip().split(' -> ')
  x1, y1 = start.split(',')
  x2, y2 = end.split(',')
  x1 = int(x1.strip())
  y1 = int(y1.strip())
  x2 = int(x2.strip())
  y2 = int(y2.strip())
  x1, x2 = min(x1,x2), max(x1,x2)
  y1, y2 = min(y1,y2), max(y1,y2)
  if x1==x2 or y1==y2:
    for x in range(x1, x2+1):
      for y in range(y1, y2+1):
        if (x,y) not in G:
          G[(x,y)]=0
        G[(x,y)] += 1

ans = 0
for k in G:
  if G[k] > 1:
    ans += 1
print(ans)