const std = @import("std");
const print = std.debug.print;

fn read_file(allocator: std.mem.Allocator, target: []const u8) ![]u8 {
    const file = try std.fs.cwd().openFile(target, .{});
    defer file.close();

    const stat = try file.stat();
    const buff = try file.readToEndAlloc(allocator, stat.size);
    return buff;
}

pub fn main() !void {
    var floor: i32 = 0;
    var pos: i32 = 0;
    var first_basement_pos: i32 = 0;
    const allocator = std.heap.page_allocator;

    // allocate file to heap.
    const buff = try read_file(allocator, "input.txt");
    defer allocator.free(buff);

    // split file by newline into iterator.
    var lines = std.mem.split(u8, buff, "\n");

    // iterate split file "lines".
    while (lines.next()) |line| {
        if (line.len > 0) {
            for (line) |dir| {
                switch (dir) {
                    '(' => floor += 1,
                    ')' => floor -= 1,
                    else => unreachable,
                }
                pos += 1;
                if (floor == -1) {
                    if (first_basement_pos == 0) {
                        first_basement_pos = pos;
                    }
                }
            }
        }
    }

    print("Part 1: {d}\n", .{floor});
    print("Part 2: {d}\n", .{first_basement_pos});
}
