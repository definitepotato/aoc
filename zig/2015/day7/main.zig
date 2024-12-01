const std = @import("std");
const print = std.debug.print;
const input = @embedFile("input.txt");

const sample_input =
    \\ 123 -> x
    \\ 456 -> y
    \\ x AND y -> d
    \\ x OR y -> e
    \\ x LSHIFT 2 -> f
    \\ y RSHIFT 2 -> g
    \\ NOT x -> h
    \\ NOT y -> i
;

const OPS = enum {
    OP_ERR,
    OP_AND,
    OP_OR,
    OP_NOT,
    OP_LSHIFT,
    OP_RSHIFT,
    OP_STORE,
};

pub fn main() void {
    var tok = std.mem.tokenize(u8, sample_input, "\n");

    while (tok.next()) |t| {
        var ins = std.mem.tokenize(u8, t, " ");

        while (ins.next()) |i| {
            print("{s}\n", .{i});
        }
        print("------\n", .{});
    }
}
