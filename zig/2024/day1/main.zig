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
    var tok = std.mem.tokenize(u8, file, " \n");
    while (tok.next()) |t| {
        const intLeft = try std.fmt.parseInt(i32, t, 10);
        const intRight = try std.fmt.parseInt(i32, tok.next().?, 10);

        try left.append(intLeft);
        try right.append(intRight);
    }

    std.mem.sort(i32, left.items, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, right.items, {}, comptime std.sort.asc(i32));
    // print("[DEBUG] {any}\n", .{left.items});
    // print("[DEBUG] {any}\n", .{right.items});

    // part1: calculate distance
    var total_distance: u32 = 0;
    for (left.items, 0..) |_, i| {
        total_distance += @abs(left.items[i] - right.items[i]);
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
