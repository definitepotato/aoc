const std = @import("std");
const print = std.debug.print;
const mem = std.mem;
const input = @embedFile("input.txt");

fn solve(ns: std.BoundedArray(usize, 15), num_ops: comptime_int) usize {
    const combo_max = std.math.pow(usize, num_ops, (ns.len - 2));
    const goal = ns.get(0);
    var combo: u32 = 0;

    combos: while (combo < combo_max) : (combo += 1) {
        var result = ns.get(1);
        var ops = combo;

        for (ns.constSlice()[2..]) |n| {
            if (result > goal) {
                continue :combos;
            }

            const op = ops % num_ops;
            ops /= num_ops;
            switch (op) {
                0 => result += n,
                1 => result *= n,
                2 => {
                    const adj: usize = if (n >= 100) 1000 else if (n >= 10) 100 else 10;
                    result = result * adj + n;
                },
                else => unreachable,
            }
        }
        if (result == goal) return goal;
    }
    return 0;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    var lines = std.mem.tokenizeScalar(u8, input, '\n');
    var result1: usize = 0;
    var result2: usize = 0;

    while (lines.next()) |line| {
        var ns = std.BoundedArray(usize, 15).init(0) catch unreachable;
        var nums = std.mem.tokenizeAny(u8, line, ": ");

        while (nums.next()) |num| {
            try ns.append(try std.fmt.parseInt(usize, num, 10));
        }

        result1 += solve(ns, 2);
        result2 += solve(ns, 3);
    }

    print("Part 1: {d}\n", .{result1});
    print("Part 2: {d}\n", .{result2});
}
