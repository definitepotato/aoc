const std = @import("std");
const print = std.debug.print;
const file = @embedFile("input.txt");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var list_left = std.ArrayList(i32).init(allocator);
    defer list_left.deinit();

    var list_right = std.ArrayList(i32).init(allocator);
    defer list_right.deinit();

    // input here.
    var tok = std.mem.tokenize(u8, file, " \n");
    while (tok.next()) |t| {
        const int_left = try std.fmt.parseInt(i32, t, 10);
        const int_right = try std.fmt.parseInt(i32, tok.next().?, 10);

        try list_left.append(int_left);
        try list_right.append(int_right);
    }

    std.mem.sort(i32, list_left.items, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, list_right.items, {}, comptime std.sort.asc(i32));

    // part1: calculate distance
    var total_distance: u32 = 0;
    for (list_left.items, 0..) |_, i| {
        total_distance += @abs(list_left.items[i] - list_right.items[i]);
    }

    print("Part 1: {d}\n", .{total_distance});

    // part2: calculate similarity
    var total_similarity: i32 = 0;
    var count_similarity: i32 = 0;
    for (list_left.items) |num_left| {
        for (list_right.items) |num_right| {
            if (num_left == num_right) {
                count_similarity += 1;
            }
        }
        total_similarity += num_left * count_similarity;
        count_similarity = 0;
    }

    print("Part 2: {d}\n", .{total_similarity});
}
