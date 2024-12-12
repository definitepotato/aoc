const std = @import("std");
const print = std.debug.print;
const exit = std.process.exit;
const file = @embedFile("input.txt");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var instructions = std.ArrayList([]const u8).init(allocator);
    defer instructions.deinit();

    var direction = std.ArrayList(u8).init(allocator);
    defer direction.deinit();

    var instructions_amount = std.ArrayList(usize).init(allocator);
    defer instructions_amount.deinit();

    var visited = std.ArrayList(usize).init(allocator);
    defer visited.deinit();

    var ins_tok = std.mem.tokenizeAny(u8, file, "\n ");
    while (ins_tok.next()) |ins| {
        try instructions.append(ins);
        // print("{s}: ", .{ins});

        if (ins_tok.next()) |amt| {
            const dir = amt[0];
            const amount = try std.fmt.parseInt(usize, amt[1..], 10);

            try direction.append(dir);
            try instructions_amount.append(amount);
        }
    }

    var idx: usize = 0;
    var accumulator: usize = 0;
    while (true) {
        try visited.append(idx);

        switch (instructions.items[idx][0]) {
            'n' => idx += 1,
            'a' => if (direction.items[idx] == '+') {
                accumulator += instructions_amount.items[idx];
                idx += 1;
            } else {
                accumulator -= instructions_amount.items[idx];
                idx += 1;
            },
            'j' => if (direction.items[idx] == '+') {
                idx += instructions_amount.items[idx];
            } else {
                idx -= instructions_amount.items[idx];
            },
            else => unreachable,
        }

        for (visited.items) |v| {
            if (v == idx) {
                print("Part 1: {d}\n", .{accumulator});
                exit('1');
            }
        }
    }
}
