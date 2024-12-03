const std = @import("std");
const print = std.debug.print;
const file = @embedFile("input.txt");

pub fn main() void {
    var instructions = std.mem.tokenize(u8, file, "\n");

    while (instructions.next()) |instr| {}
}
