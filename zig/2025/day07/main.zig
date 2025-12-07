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

// count number of timelines produced by a tachyon
pub fn solvePart2(memo: *std.AutoHashMap([2]usize, usize), pos: [2]usize, grid: *ds.Matrix(u8)) usize {
    // stop when we reach the bottom
    if (pos[1] == grid.height - 1) return 1;

    // efficiency!
    const val_ptr = memo.get(pos);
    if (val_ptr) |val| return val;

    var ans: usize = 0; // number of timelines

    // particle continues straight down if '.'
    if (grid.getXY(pos[0], pos[1]) == '.') {
        ans = solvePart2(memo, [2]usize{ pos[0], pos[1] + 1 }, grid); // recurse next position.
    } else { // split!!
        if (pos[0] > 0) ans = solvePart2(memo, [2]usize{ pos[0] - 1, pos[1] + 1 }, grid); // recurse left split.
        if (pos[0] < grid.width - 1) ans += solvePart2(memo, [2]usize{ pos[0] + 1, pos[1] + 1 }, grid); // recurse right split.
    }

    memo.put(pos, ans) catch unreachable;
    return ans;
}

pub fn solvePart1() void {
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
}

pub fn main() void {
    solvePart1();

    // find starting X
    var start: usize = 0;
    var input_it = std.mem.tokenizeAny(u8, input_file, "\n");
    while (input_it.next()) |line| {
        for (line, 0..) |c, idx| {
            if (c == 'S') {
                start = idx;
            }
        }
    }

    var map = std.AutoHashMap([2]usize, usize).init(allocator);
    var grid = ds.Matrix(u8).initFromText(allocator, input_file) catch unreachable;
    const part2 = solvePart2(&map, [2]usize{ start, 1 }, &grid);

    print("Part 2: {d}\n", .{part2});
    arena.deinit();
}
