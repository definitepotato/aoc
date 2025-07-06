const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const input = @embedFile("input.txt");

const test_input = "5 10 25";

pub fn maybeTriangle(list: std.ArrayList(usize)) bool {
    if (list.items[0] + list.items[1] <= list.items[2]) return false;
    if (list.items[1] + list.items[2] <= list.items[0]) return false;
    if (list.items[0] + list.items[2] <= list.items[1]) return false;
    return true;
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var count: usize = 0;

    var lines = std.mem.tokenizeAny(u8, input, "\n");
    while (lines.next()) |line| {
        var list = std.ArrayList(usize).init(allocator);
        defer list.deinit();

        var sides = std.mem.tokenizeAny(u8, line, "  ");
        while (sides.next()) |side| {
            const side_i = try std.fmt.parseInt(usize, side[0..], 10);
            try list.append(side_i);
        }

        if (maybeTriangle(list)) {
            count += 1;
        }
    }

    print("{d}\n", .{count});
}
