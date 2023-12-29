#!/usr/bin/env python3

instructions = [
    'move 1 from 2 to 1',
    'move 3 from 1 to 3',
    'move 2 from 2 to 1',
    'move 1 from 1 to 2'
]

cargo = {
    1: ['Z', 'N'],
    2: ['M', 'C', 'D'],
    3: ['P']
}

for instruction in instructions:
    i = instruction.split(' ')
    # i[1] move this many crates
    # i[3] from here
    # i[5] to here
    move_count = int(i[1])
    from_stack = int(i[3])
    to_stack = int(i[5])

    for i in range(move_count):
        crate = cargo[from_stack].pop()
        cargo[to_stack].append(crate)

print(cargo)

cargo2 = {
    1: ['Z', 'N'],
    2: ['M', 'C', 'D'],
    3: ['P']
}

for instruction in instructions:
    i = instruction.split(' ')
    # i[1] move this many crates
    # i[3] from here
    # i[5] to here
    move_count = int(i[1])
    from_stack = int(i[3])
    to_stack = int(i[5])

    if move_count == 1:
        crate = cargo2[from_stack].pop()
        cargo2[to_stack].append(crate)

    if move_count > 1:
        holding = []
        for i in range(move_count):
            crate = cargo2[from_stack].pop()
            holding.append(crate)
        holding.reverse()
        for crate in holding:
            cargo2[to_stack].append(crate)

print(cargo2)
