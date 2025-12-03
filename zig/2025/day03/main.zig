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
// const batt1: usize = @intCast(batteries[i] - '0');
// const batt2: usize = @intCast(batteries[j] - '0');
// var buf: [12]u8 = @splat(0); // fill a buffer with 0 values

const test_input =
    \\987654321111111
    \\811111111111119
    \\234234234234278
    \\818181911112111
;

pub fn solvePart1() void {
    var part1: usize = 0;
    var banks = std.mem.tokenizeAny(u8, input_file, "\n");
    while (banks.next()) |batteries| {
        var joltage: usize = 0;
        for (0..batteries.len) |i| {
            for (i..batteries.len) |j| {
                if (i == j) continue;
                const jolt: usize = std.fmt.parseInt(usize, &[2]u8{ batteries[i], batteries[j] }, 10) catch unreachable;
                if (joltage < jolt) joltage = jolt;
            }
        }
        part1 += joltage;
        joltage = 0;
    }

    print("Part 1: {d}\n", .{part1});
}

pub fn max(slice: []const u8) ?u8 {
    if (slice.len == 0) return null;

    var max_val: u8 = 0; // start smol
    for (slice) |byte| {
        if (byte > max_val) {
            max_val = byte;
        }
    }
    return max_val;
}

pub fn solvePart2() void {
    var part2: u64 = 0;

    var banks = std.mem.tokenizeAny(u8, input_file, "\n");
    while (banks.next()) |batteries| {
        var buf: [12]u8 = @splat(0);
        var last_idx: usize = 0;

        for (0..12) |i| {
            const right: usize = 11 - i;
            const left: usize = if (i == 0) last_idx else last_idx + 1;
            buf[i] = max(batteries[left .. batteries.len - right]).?;
            last_idx = left + std.mem.indexOfScalar(u8, batteries[left..], buf[i]).?;
        }

        const joltage: u64 = std.fmt.parseInt(u64, &buf, 10) catch unreachable;
        part2 += joltage;
    }

    print("Part 2: {d}\n", .{part2});
}

pub fn main() void {
    solvePart1();
    solvePart2();
}
