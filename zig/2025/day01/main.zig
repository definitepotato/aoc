const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const file_input = @embedFile("input.txt");

const test_input =
    \\L68
    \\L30
    \\R48
    \\L5
    \\R60
    \\L55
    \\L1
    \\L99
    \\R14
    \\L82
;

const Instruction = struct {
    dir: u8,
    distance: i32,
};

fn processInstruction(instr: []const u8) Instruction {
    const distance = std.fmt.parseInt(i32, instr[1..], 10) catch unreachable;
    return .{
        .dir = instr[0],
        .distance = distance,
    };
}

pub fn main() void {
    var dial: i32 = 50;
    var start = dial;

    var part_1: i32 = 0;
    var part_2: i32 = 0;

    var lines = std.mem.tokenizeAny(u8, file_input, "\n");
    while (lines.next()) |line| {
        const instr = processInstruction(line);

        const turns = @divFloor(instr.distance, 100);
        part_2 += turns;

        if (instr.dir == 'R') {
            if (dial + @mod(instr.distance, 100) >= 100) part_2 += 1;
            dial += instr.distance;
        }

        if (instr.dir == 'L') {
            if (dial > 0 and (dial - @mod(instr.distance, 100)) <= 0) part_2 += 1;
            start = @mod((100 - start), 100);
            dial -= instr.distance;
        }

        dial = @mod(dial, 100);
        if (dial == 0) part_1 += 1;
    }

    print("Part 1: {d}\n", .{part_1});
    print("Part 2: {d}\n", .{part_2});
}
