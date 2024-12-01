const std = @import("std");
const print = std.debug.print;
const file = @embedFile("input.txt");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var left = std.ArrayList(i32).init(allocator);
    defer left.deinit();

    var right = std.ArrayList(i32).init(allocator);
    defer right.deinit();

    // input here.
    var tok = std.mem.tokenize(u8, file, "\n");
    while (tok.next()) |t| {
        const ins = std.mem.tokenize(u8, t, " ");

        const a = try std.fmt.parseInt(i32, &[_]u8{ins.buffer[0]}, 10);
        const b = try std.fmt.parseInt(i32, &[_]u8{ins.buffer[1]}, 10);
        const c = try std.fmt.parseInt(i32, &[_]u8{ins.buffer[2]}, 10);
        const d = try std.fmt.parseInt(i32, &[_]u8{ins.buffer[3]}, 10);
        const e = try std.fmt.parseInt(i32, &[_]u8{ins.buffer[4]}, 10);
        const intLeft = (a * 10000) + (b * 1000) + (c * 100) + (d * 10) + e;

        const f = try std.fmt.parseInt(i32, &[_]u8{ins.buffer[8]}, 10);
        const g = try std.fmt.parseInt(i32, &[_]u8{ins.buffer[9]}, 10);
        const h = try std.fmt.parseInt(i32, &[_]u8{ins.buffer[10]}, 10);
        const i = try std.fmt.parseInt(i32, &[_]u8{ins.buffer[11]}, 10);
        const j = try std.fmt.parseInt(i32, &[_]u8{ins.buffer[12]}, 10);
        const intRight = (f * 10000) + (g * 1000) + (h * 100) + (i * 10) + j;

        try left.append(intLeft);
        try right.append(intRight);
    }

    std.mem.sort(i32, left.items, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, right.items, {}, comptime std.sort.asc(i32));
    // print("[DEBUG] {any}\n", .{left.items});
    // print("[DEBUG] {any}\n", .{right.items});

    // part1: calculate distance
    var total_distance: i32 = 0;
    for (left.items, 0..) |_, i| {
        if (right.items[i] > left.items[i]) {
            total_distance += right.items[i] - left.items[i];
        }

        if (right.items[i] < left.items[i]) {
            total_distance += left.items[i] - right.items[i];
        }
    }

    print("Part 1: {d}\n", .{total_distance});

    // part2: calculate similarity
    var total_similarity: i32 = 0;
    var count_similarity: i32 = 0;
    for (left.items) |numLeft| {
        for (right.items) |numRight| {
            if (numLeft == numRight) {
                count_similarity += 1;
            }
        }
        total_similarity += numLeft * count_similarity;
        count_similarity = 0;
    }

    print("Part 2: {d}\n", .{total_similarity});
}
