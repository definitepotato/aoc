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
    const allocator = std.heap.page_allocator;

    // allocate file to heap.
    const buff = try read_file(allocator, "input.txt");
    defer allocator.free(buff);

    // allocate arraylist to heap.
    var elves = std.ArrayList(i32).init(allocator);
    defer elves.deinit();

    // split file by newline into iterator.
    var lines = std.mem.split(u8, buff, "\n");

    var most: i32 = 0;
    var current_elf: i32 = 0;

    // iterate split file "lines".
    while (lines.next()) |line| {
        if (std.mem.eql(u8, line, "")) {
            try elves.append(current_elf);

            if (current_elf > most) {
                most = current_elf;
            }
            current_elf = 0;
            continue;
        }

        const num = try std.fmt.parseInt(i32, line, 10); // string to int.
        current_elf += num;
    }
    print("Part 1: {d}\n", .{most});

    std.mem.sort(i32, elves.items, {}, comptime std.sort.desc(i32)); // sort list by order desc.
    print("Part 2: {d}\n", .{elves.items[0] + elves.items[1] + elves.items[2]});
}
