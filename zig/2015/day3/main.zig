const std = @import("std");
const print = std.debug.print;
const input = @embedFile("input.txt");

const pos = struct {
    x: i32,
    y: i32,
};

fn has_been_visited(list: std.ArrayList(pos), position: pos) bool {
    for (list.items) |p| {
        if (p.x == position.x and p.y == position.y) {
            return true;
        }
    }
    return false;
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var unique_houses: i32 = 1;
    var visited = std.ArrayList(pos).init(allocator);
    defer visited.deinit();

    var current_pos: pos = .{ .x = 0, .y = 0 };
    try visited.append(current_pos);

    for (input) |direction| {
        switch (direction) {
            94 => current_pos.y += 1, // up
            62 => current_pos.x += 1, // right
            118 => current_pos.y -= 1, // down
            60 => current_pos.x -= 1, // left
            10 => continue,
            else => unreachable,
        }

        switch (has_been_visited(visited, current_pos)) {
            false => unique_houses += 1,
            true => continue,
        }
        try visited.append(current_pos);
    }

    print("Part 1: {d}\n", .{unique_houses});
}
