const std = @import("std");
const ds = @import("ds.zig");
const print = std.debug.print;
const assert = std.debug.assert;
const input_file = @embedFile("input.txt");
var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
const allocator = arena.allocator();

const test_input =
    \\..@@.@@@@.
    \\@@@.@.@.@@
    \\@@@@@.@.@@
    \\@.@@@@..@.
    \\@@.@@@@.@@
    \\.@@@@@@@.@
    \\.@.@.@.@@@
    \\@.@@@.@@@@
    \\.@@@@@@@@.
    \\@.@.@@@.@.
;

pub fn solvePart2() void {
    var part2: usize = 0;
    var score: usize = 0;

    var mess = ds.Slice([2]usize).init(allocator);
    var grid = ds.Matrix(u8).initFromText(allocator, input_file) catch unreachable;
    while (true) {
        for (0..grid.height) |y| {
            for (0..grid.width) |x| {
                const down: bool = if (y < grid.height - 1) true else false;
                const up: bool = if (y > 0) true else false;
                const left: bool = if (x > 0) true else false;
                const left_down: bool = if (left and y < grid.height - 1) true else false;
                const left_up: bool = if (left and y > 0) true else false;
                const right: bool = if (x < grid.width - 1) true else false;
                const right_down: bool = if (right and y < grid.height - 1) true else false;
                const right_up: bool = if (right and y > 0) true else false;

                var acc: usize = 0;
                if (grid.getXY(x, y) == '@') {
                    if (down and grid.getXY(x, y + 1) == '@') acc += 1; // down
                    if (left_down and grid.getXY(x - 1, y + 1) == '@') acc += 1; // left, down
                    if (left and grid.getXY(x - 1, y) == '@') acc += 1; // left
                    if (left_up and grid.getXY(x - 1, y - 1) == '@') acc += 1; // left, up
                    if (up and grid.getXY(x, y - 1) == '@') acc += 1; // up
                    if (right_up and grid.getXY(x + 1, y - 1) == '@') acc += 1; // right, up
                    if (right and grid.getXY(x + 1, y) == '@') acc += 1; // right
                    if (right_down and grid.getXY(x + 1, y + 1) == '@') acc += 1; // right,down

                    if (acc < 4) {
                        score += 1;
                        part2 += 1;
                        mess.append([2]usize{ x, y }) catch unreachable;
                    }
                }
            }
        }
        const clean_up = mess.toOwned() catch unreachable;
        if (clean_up) |roll| {
            for (roll) |r| {
                grid.setXY(r[0], r[1], '.');
            }
        }
        if (score == 0) break;
        score = 0;
    }

    print("Part 2: {d}\n", .{part2});
}

pub fn solvePart1() void {
    var part1: usize = 0;

    const grid = ds.Matrix(u8).initFromText(allocator, input_file) catch unreachable;
    for (0..grid.height) |y| {
        for (0..grid.width) |x| {
            const down: bool = if (y < grid.height - 1) true else false;
            const up: bool = if (y > 0) true else false;
            const left: bool = if (x > 0) true else false;
            const left_down: bool = if (left and y < grid.height - 1) true else false;
            const left_up: bool = if (left and y > 0) true else false;
            const right: bool = if (x < grid.width - 1) true else false;
            const right_down: bool = if (right and y < grid.height - 1) true else false;
            const right_up: bool = if (right and y > 0) true else false;

            var acc: usize = 0;
            if (grid.getXY(x, y) == '@') {
                if (down and grid.getXY(x, y + 1) == '@') acc += 1; // down
                if (left_down and grid.getXY(x - 1, y + 1) == '@') acc += 1; // left, down
                if (left and grid.getXY(x - 1, y) == '@') acc += 1; // left
                if (left_up and grid.getXY(x - 1, y - 1) == '@') acc += 1; // left, up
                if (up and grid.getXY(x, y - 1) == '@') acc += 1; // up
                if (right_up and grid.getXY(x + 1, y - 1) == '@') acc += 1; // right, up
                if (right and grid.getXY(x + 1, y) == '@') acc += 1; // right
                if (right_down and grid.getXY(x + 1, y + 1) == '@') acc += 1; // right,down

                if (acc < 4) {
                    part1 += 1;
                }
            }
        }
    }

    print("Part 1: {d}\n", .{part1});
}

pub fn main() void {
    solvePart1();
    solvePart2();
    arena.deinit();
}
