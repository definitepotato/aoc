const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const ds = @import("ds.zig");
const input_file = @embedFile("input.txt");
var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
const allocator = arena.allocator();

const test_input =
    \\123 328  51 64 
    \\ 45 64  387 23 
    \\  6 98  215 314
    \\*   +   *   +  
;

pub fn lastLine(in: []const u8) ?[]const u8 {
    var slice = ds.Slice(u8).init(allocator);

    var it_in = std.mem.tokenizeAny(u8, in, " \n");
    while (it_in.next()) |line| {
        switch (line[0]) {
            '+', '*' => slice.append(line[0]) catch unreachable,
            else => continue,
        }
    }

    return slice.toOwned() catch unreachable;
}

pub fn stoi(s: []const u8) usize {
    const my_s = s;
    const s_cleaned = std.mem.trim(u8, my_s, " ");
    return std.fmt.parseInt(u64, s_cleaned[0..], 10) catch unreachable;
}

pub fn solvePart2() void {
    var part2: u64 = 0;
    var ans: [2056 * 2]u64 = @splat(0);
    part2 += 0;

    var grid = ds.Matrix(u8).initFromText(allocator, input_file) catch unreachable;
    const new_grid = grid.rotate() catch unreachable;

    const num_size: usize = 4;
    var y: usize = 0;
    var idx: usize = 0;
    var bucket: [4]u64 = @splat(0);
    while (y < new_grid.height) : (y += 1) {
        if (new_grid.getXY(num_size, y) == '+') {
            const val = [num_size]u8{ new_grid.getXY(0, y), new_grid.getXY(1, y), new_grid.getXY(2, y), new_grid.getXY(3, y) };
            const val_i = stoi(&val);
            bucket[idx] = val_i;
            var new: u64 = 0;
            for (bucket) |a| {
                new += a;
            }
            ans[y] = new;

            y += 1;
            idx = 0;
            bucket = @splat(0);
            continue;
        }
        if (new_grid.getXY(num_size, y) == '*') {
            const val = [num_size]u8{ new_grid.getXY(0, y), new_grid.getXY(1, y), new_grid.getXY(2, y), new_grid.getXY(3, y) };
            const val_i = stoi(&val);
            bucket[idx] = val_i;
            if (ans[y] == 0) ans[y] = 1;
            var new: u64 = 1;
            for (bucket) |a| {
                if (a == 0) continue;
                new *= a;
            }
            ans[y] = new;

            y += 1;
            idx = 0;
            bucket = @splat(0);
            continue;
        }
        const val = [num_size]u8{ new_grid.getXY(0, y), new_grid.getXY(1, y), new_grid.getXY(2, y), new_grid.getXY(3, y) };
        const val_i = stoi(&val);
        bucket[idx] = val_i;
        idx += 1;
    }

    for (ans) |a| part2 += a;
    print("Part 2: {d}\n", .{part2});
}

pub fn main() void {
    var part1: u64 = 0;
    const operators = lastLine(input_file).?;
    var ans: [1024]u64 = @splat(0);

    var it_input = std.mem.tokenizeAny(u8, input_file, "\n");
    while (it_input.next()) |item| {
        var it_item = std.mem.tokenizeAny(u8, item, " ");

        var idx: usize = 0;
        while (it_item.next()) |n_or_sym| : (idx += 1) {
            if (n_or_sym[0] == '+' or n_or_sym[0] == '*') continue;

            const n = std.fmt.parseInt(u64, n_or_sym, 10) catch unreachable;
            if (operators[idx] == '+') ans[idx] += n;

            if (operators[idx] == '*') {
                if (ans[idx] > 0) {
                    ans[idx] *= n;
                    continue;
                }

                ans[idx] = n;
            }
        }
    }

    for (ans) |a| part1 += a;

    print("Part 1: {d}\n", .{part1});
    solvePart2();
    arena.deinit();
}

test "rotate" {
    var grid = try ds.Matrix(u8).initFromText(allocator, test_input);
    const new_grid = try grid.rotate();

    const val = [3]u8{ new_grid.getXY(0, 0), new_grid.getXY(1, 0), new_grid.getXY(2, 0) };
    const val_cleaned = std.mem.trim(u8, &val, " ");

    const val_i = std.fmt.parseInt(usize, val_cleaned[0..], 10) catch unreachable;
    assert(val_i == 4);
}
