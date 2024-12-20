const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const file = @embedFile("input.txt");

const Towels = "r, wr, b, g, bwu, rb, gb, br";

const test_input =
    \\brwrr
    \\bggr
    \\gbbr
    \\rrbgbr
    \\ubwu
    \\bwurrg
    \\brgr
    \\bbrgwb
;

fn completePattern(pattern: []const u8) bool {
    if (pattern.len == 0) {
        return true;
    }

    var done: bool = false;
    var towels = std.mem.tokenizeAny(u8, Towels, ", ");
    while (towels.next()) |towel| {
        if (done) {
            return done;
        }

        if (std.mem.startsWith(u8, pattern, towel)) {
            done = completePattern(pattern[towel.len..]);
        }
    }

    return false;
}

pub fn main() void {
    var part1: usize = 0;

    var tok = std.mem.tokenizeAny(u8, test_input, "\n");
    while (tok.next()) |pattern| {
        if (completePattern(pattern)) {
            part1 += 1;
        }
    }

    print("Part 1: {d}\n", .{part1});
}
