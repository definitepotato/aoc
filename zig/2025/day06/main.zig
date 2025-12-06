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
    arena.deinit();
}
