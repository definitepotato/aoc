const std = @import("std");
const print = std.debug.print;
const assert = std.debug.assert;
const input_file = @embedFile("input.txt");
var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
const allocator = arena.allocator();

const test_input =
    \\7,1
    \\11,1
    \\11,7
    \\9,7
    \\9,5
    \\2,5
    \\2,3
    \\7,3
;

pub fn main() void {

    // build array
    var list: std.ArrayList([2]i64) = .empty;
    var tok_it = std.mem.tokenizeAny(u8, input_file, "\n");
    while (tok_it.next()) |pos| {
        var pos_it = std.mem.tokenizeAny(u8, pos, ",");
        while (pos_it.next()) |x| {
            const y = pos_it.next().?;
            const x_i = std.fmt.parseInt(i64, x[0..], 10) catch unreachable;
            const y_i = std.fmt.parseInt(i64, y[0..], 10) catch unreachable;

            list.append(allocator, [2]i64{ x_i, y_i }) catch unreachable;
        }
    }

    var max: u64 = 0;
    for (list.items) |pos1| {
        for (list.items) |pos2| {
            const x1 = pos1[0];
            const y1 = pos1[1];
            const x2 = pos2[0];
            const y2 = pos2[1];

            if (x1 == x2 and y1 == y2) continue;

            if (x2 >= x1 and y2 >= y1) {
                const w = (@abs(x1 - x2)) + 1;
                const h = (@abs(y1 - y2)) + 1;
                const a = w * h;
                if (a > max) max = a;
            }
        }
    }

    print("{d}\n", .{max});
    arena.deinit();
}
