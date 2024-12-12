#!/usr/bin/env python3
# -*- coding=utf-8 -*-

file = open("input.txt", "r")

left_list = []
right_list = []

for line in file:
    row = line.strip().split("   ")
    left_list.append(int(row[0]))
    right_list.append(int(row[1]))

total_distance = 0
for idx, left_int in enumerate(sorted(left_list)):
    right_int = sorted(right_list)[idx]
    total_distance += abs(left_int - right_int)

print(f"Part 1: {total_distance}")

total_similarity = 0
count_similarity = 0

for idx, left_int in enumerate(sorted(left_list)):
    for idx, right_int in enumerate(sorted(right_list)):
        if left_int == right_int:
            count_similarity += 1

    total_similarity += left_int * count_similarity
    count_similarity = 0

print(f"Part 2: {total_similarity}")
