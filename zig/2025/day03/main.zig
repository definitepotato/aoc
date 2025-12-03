const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const input_file = @embedFile("input.txt");

// [1]
// u8 characters are stored as integer byte values
// | Character | ASCII Value |
// | --------- | ----------- |
// |    '0'    |     48      |
// |    '1'    |     49      |
// |    ...    |     ...     |
// |    '9'    |     57      |
//
// if you subtract '0' it will "convert" from it's u8 byte value to an integer value
// i.e. '5' - '0' = 53 - 48 = 5

const test_input =
    \\987654321111111
    \\811111111111119
    \\234234234234278
    \\818181911112111
;

pub fn main() void {
    var part1: usize = 0;
    var banks = std.mem.tokenizeAny(u8, input_file, "\n");
    while (banks.next()) |batteries| {
        var joltage: usize = 0;
        for (0..batteries.len) |i| {
            for (i..batteries.len) |j| {
                if (i == j) continue;
                // const batt1: usize = @intCast(batteries[i] - '0');
                // const batt2: usize = @intCast(batteries[j] - '0');
                const jolt: usize = std.fmt.parseInt(usize, &[2]u8{ batteries[i], batteries[j] }, 10) catch unreachable;
                if (joltage < jolt) joltage = jolt;
            }
        }
        print("{s} => {d}\n", .{ batteries, joltage });
        part1 += joltage;
        joltage = 0;
    }

    print("Part 1: {d}\n", .{part1});
}
