const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const ds = @import("ds.zig");
const input_file = @embedFile("input.txt");
var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
const allocator = arena.allocator();

const test_input =
    \\.......S.......
    \\...............
    \\.......^.......
    \\...............
    \\......^.^......
    \\...............
    \\.....^.^.^.....
    \\...............
    \\....^.^...^....
    \\...............
    \\...^.^...^.^...
    \\...............
    \\..^...^.....^..
    \\...............
    \\.^.^.^.^.^...^.
    \\...............
;

pub fn main() void {
    var part1: usize = 0;

    var grid = ds.Matrix(u8).initFromText(allocator, input_file) catch unreachable;

    for (1..grid.height) |y| {
        for (1..grid.width) |x| {
            if (grid.getXY(x, y) == '.' and grid.getXY(x, y - 1) == '|' or grid.getXY(x, y - 1) == 'S') {
                grid.setXY(x, y, '|');
            } else if (grid.getXY(x, y) == '^' and grid.getXY(x, y - 1) == '|' or grid.getXY(x, y - 1) == 'S') {
                grid.setXY(x - 1, y, '|');
                grid.setXY(x + 1, y, '|');
                grid.setXY(x, y, 'x');
                part1 += 1;
            }
        }
    }

    print("Part 1: {d}\n", .{part1});
    arena.deinit();
}
