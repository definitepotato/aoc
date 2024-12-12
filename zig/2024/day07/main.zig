const std = @import("std");
const math = std.math;
const print = std.debug.print;
const file = @embedFile("test.txt");

fn processRow(row: []const u8, map: *std.AutoHashMap(usize, []usize)) !void {
    const allocator = std.heap.page_allocator;

    var list = std.ArrayList(usize).init(allocator);
    defer allocator.free(list);

    var r = std.mem.tokenize(u8, row, ": ");
    const key = r.next();
    const key_int = try std.fmt.parseInt(usize, key.?, 10);

    while (r.next()) |value| {
        const value_int = try std.fmt.parseInt(usize, value, 10);
        try list.append(value_int);
    }

    const values = try list.toOwnedSlice();
    try map.put(key_int, values);
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var map = std.AutoHashMap(usize, []usize).init(allocator);
    defer map.deinit();

    var line = std.mem.tokenize(u8, file, "\n");
    while (line.next()) |l| {
        try processRow(l, &map);
    }

    var map_it = map.iterator();
    while (map_it.next()) |entry| {
        print("{any}\n", .{entry});
    }
}
